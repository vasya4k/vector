package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Juniper/go-netconf/netconf"
	"github.com/Sirupsen/logrus"
	elastic "gopkg.in/olivere/elastic.v3"
)

const (
	index        = "telsps"
	docType      = "telsp"
	indexMapping = `{
                        "mappings" : {
                            "telsp" : {
                                "properties" : {
                                    "time" : { "type" : "date" }
                                }
                            }
                        }
                    }`
)

var trim = strings.TrimSpace

//LSP  type for frontend rendering
type LSP struct {
	Name        string    `json:"name"`
	State       string    `json:"state"`
	Hostname    string    `json:"hostname"`
	SessionType string    `json:"sessiontype"`
	PathStatus  string    `json:"pathstatus"`
	PathType    string    `json:"pathtype"`
	History     string    `json:"history"`
	From        string    `json:"from"`
	To          string    `json:"to"`
	RecordRoute string    `json:"recordroute"`
	Bw          int64     `json:"bw"`
	Time        time.Time `json:"time"`
}

func esAdd(c *elastic.Client, lsp LSP) {
	_, err := c.Index().
		Index(index).
		Type(docType).
		BodyJson(lsp).
		Refresh(true).
		Do()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"topic": "elasticsearch",
			"event": "failure to unmarshal settings document",
		}).Error(err)
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func main() {
	ran := false
	logrus.WithFields(logrus.Fields{
		"topic": "start",
	}).Info("the  service is starting")
	for {
		if ran {
			time.Sleep(1 * time.Minute)
		}

		hosts := []string{"13.55.180.89", "13.55.70.196", "54.79.67.55"}
		lsps := pollLSPsJunipers(&hosts)
		eClient, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))

		exists, err := eClient.IndexExists("telsps").Do()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"topic": "elasticsearch",
				"event": "failure to check index",
			}).Error(err)
			os.Exit(1)
		}
		if !exists {
			// Create a new index.
			createIndex, err := eClient.CreateIndex(index).Body(indexMapping).Do()
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"topic": "elasticsearch",
					"event": "failure to create index",
				}).Error(err)
				os.Exit(1)
			}
			if !createIndex.Acknowledged {
				logrus.WithFields(logrus.Fields{
					"topic": "elasticsearch",
					"event": "failure to Acknowledged create index",
				}).Error(err)
				os.Exit(1)
			}
		}

		for _, v := range *lsps {
			for _, rdata := range v.RsvpSessionData {
				if trim(rdata.SessionType.Text) == "Ingress" {
					for _, session := range rdata.RsvpSession {
						var lsp LSP
						lsp.Name = trim(session.MplsLsp.Name)
						lsp.SessionType = "MplsLspInformation"
						lsp.State = trim(session.MplsLsp.LspState[0])
						lsp.PathStatus = trim(session.MplsLsp.MplsLspPath[0].PathState)
						lsp.PathType = trim(session.MplsLsp.MplsLspPath[0].Title)
						lsp.From = trim(session.MplsLsp.SourceAddress[0])
						lsp.To = trim(session.MplsLsp.DestinationAddress)
						lsp.Time = time.Now()
						lsp.Hostname = v.Host

						d, err := strconv.ParseInt(strings.Split(session.MplsLsp.MplsLspPath[0].Bandwidth, "kbps")[0], 0, 64)
						if err != nil {
							logrus.WithFields(logrus.Fields{
								"topic": "ParseInt",
								"event": "ParseInt",
							}).Error(err)
						}
						lsp.Bw = d*1000 + int64(random(10, 50))

						bstr, err := json.MarshalIndent(lsp, "", "  ")
						if err != nil {
							logrus.WithFields(logrus.Fields{
								"topic": "marshal",
								"event": "MarshalIndent",
							}).Error(err)
						}
						fmt.Println(fmt.Sprintf("%s", bstr))
						esAdd(eClient, lsp)
					}
				}
			}
		}
		ran = true
	}
}

func getLSPInfo(s *netconf.Session) (*MplsLspInformation, error) {
	var lsps MplsLspInformation
	reply, err := s.Exec(netconf.RawMethod("<get-mpls-lsp-information><detail/></get-mpls-lsp-information>"))
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal([]byte(reply.Data), &lsps)
	if err != nil {
		return nil, err
	}
	return &lsps, nil
}

func pollLSPsJunipers(hosts *[]string) *[]MplsLspInformation {
	var (
		lsps []MplsLspInformation
	)
	wg := &sync.WaitGroup{}
	wg.Add(len(*hosts))
	for _, host := range *hosts {
		go pollLSPsJuniper(wg, host, &lsps)
	}
	wg.Wait()
	return &lsps
}

func pollLSPsJuniper(wg *sync.WaitGroup, host string, data *[]MplsLspInformation) {
	defer wg.Done()
	session, err := netconf.DialSSH(host, netconf.SSHConfigPassword("ntc", "ntc123"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"topic": "DialSSH",
			"event": "DialSSH failure",
		}).Error(err)
		return
	}
	defer session.Close()
	lspInfo, err := getLSPInfo(session)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"topic": "getLSPInfo",
			"event": "getLSPInfo failure",
		}).Error(err)
		return
	}
	if host == "13.55.180.89" {
		lspInfo.Host = "vmx7"
	}
	if host == "13.55.70.196" {
		lspInfo.Host = "vmx8"
	}
	if host == "54.79.67.55" {
		lspInfo.Host = "vmx9"
	}
	*data = append(*data, *lspInfo)
}
