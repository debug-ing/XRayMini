package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/lilendian0x00/xray-knife/v2/xray"
)

type Inbound struct {
	Port     string      `json:"port"`
	Protocol string      `json:"protocol"`
	Settings interface{} `json:"settings"`
}
type Config struct {
	Log      interface{}            `json:"log"`
	Inbounds []Inbound              `json:"inbounds"`
	Outbound []OutboundDetourConfig `json:"outbounds"`
}

////////////

type User struct {
	ID         string `json:"id"`
	AlterID    int    `json:"alterId"`
	Security   string `json:"security"`
	Flow       string `json:"flow"`
	Encryption string `json:"encryption"`
}

type VNext struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	Users   []User `json:"users"`
}

type Settings struct {
	VNext []VNext `json:"vnext"`
}

type Header struct {
	Type string `json:"type"`
}

type TCPSettings struct {
	Header              Header `json:"header"`
	AcceptProxyProtocol bool   `json:"acceptProxyProtocol"`
}

type StreamSettings struct {
	Network             string       `json:"network"`
	Security            string       `json:"security"`
	TLSSettings         interface{}  `json:"tlsSettings"`
	RealitySettings     interface{}  `json:"realitySettings"`
	TCPSettings         *TCPSettings `json:"tcpSettings"`
	KCPSettings         interface{}  `json:"kcpSettings"`
	WSSettings          interface{}  `json:"wsSettings"`
	HTTPSettings        interface{}  `json:"httpSettings"`
	DSSettings          interface{}  `json:"dsSettings"`
	QUICSettings        interface{}  `json:"quicSettings"`
	Sockopt             interface{}  `json:"sockopt"`
	GRPCSettings        interface{}  `json:"grpcSettings"`
	GunSettings         interface{}  `json:"gunSettings"`
	HTTPUpgradeSettings interface{}  `json:"httpupgradeSettings"`
	SplitHTTPSettings   interface{}  `json:"splithttpSettings"`
}

type OutboundDetourConfig struct {
	Protocol       string         `json:"protocol"`
	SendThrough    interface{}    `json:"sendThrough"`
	Tag            string         `json:"tag"`
	Settings       Settings       `json:"settings"`
	StreamSettings StreamSettings `json:"streamSettings"`
	ProxySettings  interface{}    `json:"proxySettings"`
	Mux            interface{}    `json:"mux"`
}

////////

func main() {
	reader := bufio.NewReader(os.Stdin)
	rawURL, _ := reader.ReadString('\n')
	rawURL = strings.TrimSpace(rawURL)
	protocol, err := xray.ParseXrayConfig(rawURL)
	if err != nil {
		fmt.Println(err)
	}
	conf, err := protocol.BuildOutboundDetourConfig(true)
	if err != nil {
		fmt.Println(err)
	}
	jsonBytes, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		fmt.Println("❌ Error converting to JSON:", err)
		return
	}
	config := Config{}
	config.Log = map[string]interface{}{}
	config.Inbounds = []Inbound{
		{
			Port:     "1080",
			Protocol: "socks",
			Settings: map[string]interface{}{
				"udp": true,
			},
		},
		{
			Port:     "1081",
			Protocol: "http",
			Settings: map[string]interface{}{},
		},
	}
	config.Outbound = []OutboundDetourConfig{
		{},
	}
	err = json.Unmarshal(jsonBytes, &config.Outbound[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonBytes, err = json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("❌ Error converting to JSON:", err)
		return
	}
	fmt.Println(string(jsonBytes))
}
