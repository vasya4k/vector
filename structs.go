package main

import (
	"encoding/xml"
)

type RsvpSessionInformation struct {
	Host            string
	AttrXmlns       string             `xml:" xmlns,attr"  json:""`
	RsvpSessionData []*RsvpSessionData `xml:"rsvp-session-data"`
	XMLName         xml.Name           `xml:"rsvp-session-information"`
}

type ActivePath struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"active-path" json:"active-path"`
}

type AdjustThreshold struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"adjust-threshold" json:"adjust-threshold"`
}

type AdjustTimer struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"adjust-timer" json:"adjust-timer"`
}

type Adspec struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"adspec" json:"adspec"`
}

type Bandwidth struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"bandwidth" json:"bandwidth"`
}

type DetourBranch struct {
	Adspec             []*Adspec            `xml:"adspec" json:"adspec"`
	DetourBranchLabels *DetourBranchLabels  `xml:"detour-branch-labels" json:"detour-branch-labels"`
	ExplicitRoute      *ExplicitRoute       `xml:"explicit-route" json:"explicit-route"`
	LspState           []string             `xml:"lsp-state" json:"lsp-state"`
	PacketInformation  []*PacketInformation `xml:"packet-information" json:"packet-information"`
	PathMtu            []*PathMtu           `xml:"path-mtu" json:"path-mtu"`
	RecordRoute        *RecordRoute         `xml:"record-route" json:"record-route"`
	SenderTspec        []*SenderTspec       `xml:"sender-tspec" json:"sender-tspec"`
	SkipAddress        []*SkipAddress       `xml:"skip-address" json:"skip-address"`
	SourceAddress      []string             `xml:"source-address" json:"source-address"`
	XMLName            xml.Name             `xml:"detour-branch" json:"detour-branch"`
}

type DetourBranchLabels struct {
	LabelIn  *LabelIn  `xml:"label-in" json:"label-in"`
	LabelOut *LabelOut `xml:"label-out" json:"label-out"`
	XMLName  xml.Name  `xml:"detour-branch-labels" json:"detour-branch-labels"`
}

type DisplayCount struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"display-count" json:"display-count"`
}

type DownCount struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"down-count" json:"down-count"`
}

type EgressLabelOperation struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"egress-label-operation" json:"egress-label-operation"`
}

type EncodingType struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"encoding-type" json:"encoding-type"`
}

type ExplicitRoute struct {
	AttrHeading       string   `xml:" heading,attr"  json:""`
	Address           []string `xml:"address" json:"address"`
	ExplicitRouteType []string `xml:"explicit-route-type" json:"explicit-route-type"`
}

type Gpid struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"gpid" json:"gpid"`
}

type Incomplete struct {
	XMLName xml.Name `xml:"incomplete" json:"incomplete"`
}

type IsNodeprotection struct {
	XMLName xml.Name `xml:"is-nodeprotection" json:"is-nodeprotection"`
}

type IsSoftPreemption struct {
	XMLName xml.Name `xml:"is-soft-preemption" json:"is-soft-preemption"`
}

type LabelIn struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"label-in" json:"label-in"`
}

type LabelOut struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"label-out" json:"label-out"`
}

type LoadBalance struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"load-balance" json:"load-balance"`
}

type LspAttributeFlags struct {
	XMLName xml.Name `xml:"lsp-attribute-flags" json:"lsp-attribute-flags"`
}

type LspDescription struct {
	XMLName xml.Name `xml:"lsp-description" json:"lsp-description"`
}

type LspId struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"lsp-id" json:"lsp-id"`
}

type LspType struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"lsp-type" json:"lsp-type"`
}

type MaximumAverageBandwidth struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"maximum-average-bandwidth" json:"maximum-average-bandwidth"`
}

type MaximumBandwidth struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"maximum-bandwidth" json:"maximum-bandwidth"`
}

type MinimumBandwidth struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"minimum-bandwidth" json:"minimum-bandwidth"`
}

type MplsLsp struct {
	ActivePath           *ActivePath           `xml:"active-path" json:"active-path"`
	DestinationAddress   string                `xml:"destination-address" json:"destination-address"`
	EgressLabelOperation *EgressLabelOperation `xml:"egress-label-operation" json:"egress-label-operation"`
	IsNodeprotection     *IsNodeprotection     `xml:"is-nodeprotection" json:"is-nodeprotection"`
	LoadBalance          *LoadBalance          `xml:"load-balance" json:"load-balance"`
	LspDescription       *LspDescription       `xml:"lsp-description" json:"lsp-description"`
	LspState             []string              `xml:"lsp-state" json:"lsp-state"`
	LspType              *LspType              `xml:"lsp-type" json:"lsp-type"`
	MplsLspAttributes    *MplsLspAttributes    `xml:"mpls-lsp-attributes" json:"mpls-lsp-attributes"`
	MplsLspAutobandwidth *MplsLspAutobandwidth `xml:"mpls-lsp-autobandwidth" json:"mpls-lsp-autobandwidth"`
	MplsLspPath          []*MplsLspPath        `xml:"mpls-lsp-path" json:"mpls-lsp-path"`
	Name                 string                `xml:"name" json:"name"`
	RouteCount           *RouteCount           `xml:"route-count" json:"route-count"`
	SourceAddress        []string              `xml:"source-address" json:"source-address"`
	XMLName              xml.Name              `xml:"mpls-lsp" json:"mpls-lsp"`
}

type MplsLspAttributes struct {
	EncodingType         *EncodingType         `xml:"encoding-type" json:"encoding-type"`
	Gpid                 *Gpid                 `xml:"gpid" json:"gpid"`
	MplsLspUpstreamLabel *MplsLspUpstreamLabel `xml:"mpls-lsp-upstream-label" json:"mpls-lsp-upstream-label"`
	SwitchingType        *SwitchingType        `xml:"switching-type" json:"switching-type"`
	XMLName              xml.Name              `xml:"mpls-lsp-attributes" json:"mpls-lsp-attributes"`
}

type MplsLspAutobandwidth struct {
	AdjustThreshold          *AdjustThreshold          `xml:"adjust-threshold" json:"adjust-threshold"`
	AdjustTimer              *AdjustTimer              `xml:"adjust-timer" json:"adjust-timer"`
	MaximumAverageBandwidth  *MaximumAverageBandwidth  `xml:"maximum-average-bandwidth" json:"maximum-average-bandwidth"`
	MaximumBandwidth         *MaximumBandwidth         `xml:"maximum-bandwidth" json:"maximum-bandwidth"`
	MinimumBandwidth         *MinimumBandwidth         `xml:"minimum-bandwidth" json:"minimum-bandwidth"`
	OverflowLimit            *OverflowLimit            `xml:"overflow-limit" json:"overflow-limit"`
	OverflowSampleCount      *OverflowSampleCount      `xml:"overflow-sample-count" json:"overflow-sample-count"`
	TimeToAdjust             *TimeToAdjust             `xml:"time-to-adjust" json:"time-to-adjust"`
	UnderflowLimit           *UnderflowLimit           `xml:"underflow-limit" json:"underflow-limit"`
	UnderflowMaxAvgBandwidth *UnderflowMaxAvgBandwidth `xml:"underflow-max-avg-bandwidth" json:"underflow-max-avg-bandwidth"`
	UnderflowSampleCount     *UnderflowSampleCount     `xml:"underflow-sample-count" json:"underflow-sample-count"`
	XMLName                  xml.Name                  `xml:"mpls-lsp-autobandwidth" json:"mpls-lsp-autobandwidth"`
}

type MplsLspInformation struct {
	Host            string
	AttrXmlns       string             `xml:" xmlns,attr"  json:""`
	RsvpSessionData []*RsvpSessionData `xml:"rsvp-session-data" json:"rsvp-session-data"`
	XMLName         xml.Name           `xml:"mpls-lsp-information" json:"mpls-lsp-information"`
}

type PathHistory struct {
	Log            string `xml:"log" json:"log"`
	Route          string `xml:"route" json:"route"`
	SequenceNumber string `xml:"sequence-number" json:"sequence-number"`
	Time           string `xml:"time" json:"time"`
}

type MplsLspPath struct {
	ActualBandwidth    string         `xml:"actual-bandwidth" json:"actual-bandwidth"`
	Bandwidth          string         `xml:"bandwidth" json:"bandwidth"`
	CspfStatus         []string       `xml:"cspf-status" json:"cspf-status"`
	ExplicitRoute      *ExplicitRoute `xml:"explicit-route" json:"explicit-route"`
	HoldPriority       string         `xml:"hold-priority" json:"hold-priority"`
	Name               string         `xml:"name" json:"name"`
	OptimizeTimer      string         `xml:"optimize-timer" json:"optimize-timer"`
	PathHistory        *PathHistory   `xml:"path-history" json:"path-history"`
	PathActive         *PathActive    `xml:"path-active" json:"path-active"`
	PathState          string         `xml:"path-state" json:"path-state"`
	ReceivedRro        string         `xml:"received-rro" json:"received-rro"`
	SetupPriority      string         `xml:"setup-priority" json:"setup-priority"`
	SmartOptimizeTimer string         `xml:"smart-optimize-timer" json:"smart-optimize-timer"`
	Title              string         `xml:"title" json:"title"`
}

type MplsLspUpstreamLabel struct {
	XMLName xml.Name `xml:"mpls-lsp-upstream-label" json:"mpls-lsp-upstream-label"`
}

// type Name struct {
//  Text    string   `xml:",chardata" json:""`
//  XMLName xml.Name `xml:"name" json:"name"`
// }

type OverflowLimit struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"overflow-limit" json:"overflow-limit"`
}

type OverflowSampleCount struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"overflow-sample-count" json:"overflow-sample-count"`
}

type PacketInformation struct {
	AttrHeading   string   `xml:" heading,attr"  json:""`
	Count         string   `xml:"count" json:"count"`
	EntropyLabel  string   `xml:"entropy-label" json:"entropy-label"`
	InterfaceName string   `xml:"interface-name" json:"interface-name"`
	NextHop       string   `xml:"next-hop" json:"next-hop"`
	PreviousHop   string   `xml:"previous-hop" json:"previous-hop"`
	XMLName       xml.Name `xml:"packet-information" json:"packet-information"`
}

type PathActive struct {
	XMLName xml.Name `xml:"path-active" json:"path-active"`
}

type PathMtu struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"path-mtu" json:"path-mtu"`
}

type ProtoId struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"proto-id" json:"proto-id"`
}

type PsbCreationTime struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"psb-creation-time" json:"psb-creation-time"`
}

type PsbLifetime struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"psb-lifetime" json:"psb-lifetime"`
}

type RecordRoute struct {
	AttrHeading string      `xml:" heading,attr"  json:""`
	Address     []string    `xml:"address" json:"address"`
	Incomplete  *Incomplete `xml:"incomplete" json:"incomplete"`
	Self        *Self       `xml:"self" json:"self"`
	XMLName     xml.Name    `xml:"record-route" json:"record-route"`
}

type RecoveryLabelIn struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"recovery-label-in" json:"recovery-label-in"`
}

type RecoveryLabelOut struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"recovery-label-out" json:"recovery-label-out"`
}

type ResvStyle struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"resv-style" json:"resv-style"`
}

type RouteCount struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"route-count" json:"route-count"`
}

type RsbCount struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"rsb-count" json:"rsb-count"`
}

type RsvpSession struct {
	AttrJunosStyle     string               `xml:"junos style,attr"  json:""`
	Adspec             []*Adspec            `xml:"adspec" json:"adspec"`
	DestinationAddress string               `xml:"destination-address" json:"destination-address"`
	DetourBranch       *DetourBranch        `xml:"detour-branch" json:"detour-branch"`
	ExplicitRoute      *ExplicitRoute       `xml:"explicit-route" json:"explicit-route"`
	IsNodeprotection   *IsNodeprotection    `xml:"is-nodeprotection" json:"is-nodeprotection"`
	IsSoftPreemption   *IsSoftPreemption    `xml:"is-soft-preemption" json:"is-soft-preemption"`
	LabelIn            *LabelIn             `xml:"label-in" json:"label-in"`
	LabelOut           *LabelOut            `xml:"label-out" json:"label-out"`
	LspAttributeFlags  *LspAttributeFlags   `xml:"lsp-attribute-flags" json:"lsp-attribute-flags"`
	LspId              *LspId               `xml:"lsp-id" json:"lsp-id"`
	LspPathType        string               `xml:"lsp-path-type" json:"lsp-path-type"`
	LspState           []string             `xml:"lsp-state" json:"lsp-state"`
	MplsLsp            *MplsLsp             `xml:"mpls-lsp" json:"mpls-lsp"`
	Name               string               `xml:"name" json:"name"`
	PacketInformation  []*PacketInformation `xml:"packet-information" json:"packet-information"`
	ProtoId            *ProtoId             `xml:"proto-id" json:"proto-id"`
	PsbCreationTime    *PsbCreationTime     `xml:"psb-creation-time" json:"psb-creation-time"`
	PsbLifetime        *PsbLifetime         `xml:"psb-lifetime" json:"psb-lifetime"`
	RecordRoute        *RecordRoute         `xml:"record-route" json:"record-route"`
	RecoveryLabelIn    *RecoveryLabelIn     `xml:"recovery-label-in" json:"recovery-label-in"`
	RecoveryLabelOut   *RecoveryLabelOut    `xml:"recovery-label-out" json:"recovery-label-out"`
	ResvStyle          *ResvStyle           `xml:"resv-style" json:"resv-style"`
	RouteCount         *RouteCount          `xml:"route-count" json:"route-count"`
	RsbCount           *RsbCount            `xml:"rsb-count" json:"rsb-count"`
	RsvpPathStatus     string               `xml:"rsvp-path-status" json:"rsvp-path-status"`
	SenderTspec        []string             `xml:"sender-tspec" json:"sender-tspec"`
	SourceAddress      []string             `xml:"source-address" json:"source-address"`
	SuggestedLabelIn   *SuggestedLabelIn    `xml:"suggested-label-in" json:"suggested-label-in"`
	SuggestedLabelOut  *SuggestedLabelOut   `xml:"suggested-label-out" json:"suggested-label-out"`
	TunnelId           *TunnelId            `xml:"tunnel-id" json:"tunnel-id"`
	XMLName            xml.Name             `xml:"rsvp-session" json:"rsvp-session"`
}

type RsvpSessionData struct {
	Count        string         `xml:"count" json:"count"`
	DisplayCount *DisplayCount  `xml:"display-count" json:"display-count"`
	DownCount    *DownCount     `xml:"down-count" json:"down-count"`
	RsvpSession  []*RsvpSession `xml:"rsvp-session" json:"rsvp-session"`
	SessionType  *SessionType   `xml:"session-type" json:"session-type"`
	UpCount      *UpCount       `xml:"up-count" json:"up-count"`
	XMLName      xml.Name       `xml:"rsvp-session-data" json:"rsvp-session-data"`
}

type Self struct {
	XMLName xml.Name `xml:"self" json:"self"`
}

type SenderTspec struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"sender-tspec" json:"sender-tspec"`
}

type SessionType struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"session-type" json:"session-type"`
}

type SkipAddress struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"skip-address" json:"skip-address"`
}

type SuggestedLabelIn struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"suggested-label-in" json:"suggested-label-in"`
}

type SuggestedLabelOut struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"suggested-label-out" json:"suggested-label-out"`
}

type SwitchingType struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"switching-type" json:"switching-type"`
}

type TimeToAdjust struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"time-to-adjust" json:"time-to-adjust"`
}

type TunnelId struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"tunnel-id" json:"tunnel-id"`
}

type UnderflowLimit struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"underflow-limit" json:"underflow-limit"`
}

type UnderflowMaxAvgBandwidth struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"underflow-max-avg-bandwidth" json:"underflow-max-avg-bandwidth"`
}

type UnderflowSampleCount struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"underflow-sample-count" json:"underflow-sample-count"`
}

type UpCount struct {
	Text    string   `xml:",chardata" json:""`
	XMLName xml.Name `xml:"up-count" json:"up-count"`
}
