###############################
# Author: Chamil Jayasundara
# Date: 9/Sep/2017
# Desc: At FB Neteng Hackathone
################################


import requests
from ncclient import manager
import sched
import time
import json
from collections import defaultdict


######Global Config
threshold = 145000
freq = 10


class Elasticsearch_DB:
    def __init__(self, ip):
        self.url = "http://" + ip + ":9200/telsps/_search"
        # self.s = sched.scheduler(time.time, time.sleep)


    def get_data(self):
        print("hello")

        data = json.dumps(
            {
                "query": {"match_all": {}},
                "_source": ["name", "bw", "hostname", "time", "to"],
                "size": 1,
                "sort": [
                    {"time": "desc"}
                ]
            })

        response = requests.post(self.url, data=data)
        results_dict = json.loads(response.text)

        results = []

        if results_dict['hits']['hits']:
            for i in results_dict['hits']['hits']:
                results.append(results_dict['hits']['hits'][0]['_source'])

        return results
        

class Switch:

    def __init__(self, host, user, password):
        self.lsp = 1
        self.conn = manager.connect(
            host=host,
            username=user,
            password=password,
            timeout=10,
            device_params = {'name':'junos'},
            hostkey_verify=False)

    def send_cmd_and_validate(self, cmd):
        self.conn.lock()
        send_config = self.conn.load_configuration(action='set', config=cmd)

        check_config = self.conn.validate()
        print(check_config.tostring)

        compare_config = self.conn.compare_configuration()
        print(compare_config.tostring)

        self.conn.commit()
        self.conn.unlock()
        # self.conn.close_session()


def periodic_schedule(scheduler, fun):
    log_data = fun()
    print(log_data)

    if int(log_data[0]['bw']) > threshold:
        sw = switches[log_data[0]['hostname']]

        if sw.lsp > 10:
            print("Max Number of LSPs created")

        else:
            lsp = log_data[0]['name']
            to = log_data[0]['to']
            new_lsp = lsp + '-' + str(sw.lsp)
            print("Bandwidth" + lsp + "is more than  " + str(threshold) + "\n")

            print("Creating New LSP" + new_lsp + " in " + log_data[0]['hostname'] + '\n')
            # print(create_cmd(new_lsp, to))
            sw.send_cmd_and_validate(create_cmd(new_lsp, to))

            sw.lsp += 1

    scheduler.enter(freq, 1, periodic_schedule, (scheduler,fun))

def create_cmd(lsp, to):
    common_part = "set protocols mpls label-switched-path " + lsp + " "

    cmd = []

    for postfix in ["to " + to, "exclude-srlg", "least-fill", "auto-bandwidth adjust-threshold 5",
                    "auto-bandwidth minimum-bandwidth 100k",  "auto-bandwidth maximum-bandwidth 100g", "optimize-on-change link-congestion"]:
        cmd.append(common_part + postfix)

    return cmd



if __name__ == '__main__':

    switches = {
        "vmx7" : Switch('13.55.180.89', 'ntc', 'ntc123'),
        "vmx8" : Switch('13.55.70.196', 'ntc', 'ntc123'),
        "vmx9" : Switch('54.79.67.55', 'ntc', 'ntc123'),
    }


    ##DB Query
    s = sched.scheduler(time.time, time.sleep)
    db = Elasticsearch_DB("13.210.95.195")

    periodic_schedule(s, db.get_data)
    s.run()


