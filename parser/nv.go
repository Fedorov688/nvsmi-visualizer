// created by D. "Mordok" Fedorov

package parser

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"
)

type NVlog struct {
	//Timestamp string `json:"#Time,omitempty"`
	GPU      string `json:"gpu,omitempty"`
	PWR      string `json:"pwr,omitempty"`
	GTemp    string `json:"gtemp,omitempty"`
	MTemp    string `json:"mtemp,omitempty"`
	SM       string `json:"sm,omitempty"`
	Mem      string `json:"mem,omitempty"`
	Enc      string `json:"enc,omitempty"`
	Dec      string `json:"dec,omitempty"`
	MClk     string `json:"mclk,omitempty"`
	PClk     string `json:"pclk,omitempty"`
	Hostname string `json:"hostname,omitempty"`
}

type NVlogRes struct {
	Timestamp string `json:"#Time,omitempty"`
	GPU       int    `json:"gpu,omitempty"`
	PWR       int    `json:"pwr,omitempty"`
	GTemp     int    `json:"gtemp,omitempty"`
	MTemp     int    `json:"mtemp,omitempty"`
	SM        int    `json:"sm,omitempty"`
	Mem       int    `json:"mem,omitempty"`
	Enc       int    `json:"enc,omitempty"`
	Dec       int    `json:"dec,omitempty"`
	MClk      int    `json:"mclk,omitempty"`
	PClk      int    `json:"pclk,omitempty"`
	Hostname  string `json:"hostname,omitempty"`
}

func ParseNVSMI(log string, host string) (resJsons [][]byte) {
	res := parseNVSMI(log, host)
	return convertMap2Jsons(res)
}

func parseNVSMI(log string, host string) (res []map[string]string) {
	data := strings.Split(log, "\n")
	if len(data) < 3 {
		return
	}
	keys := parseValue(data[0])
	//metrics := parseValue(data[1])
	//fmt.Println(metrics) TODO( metrics need or not? )
	for _, value := range data[2:] {
		val := parseValue(value)
		if len(val) < 1 {
			continue
		}
		resTmp := make(map[string]string)
		for k, key := range keys {
			resTmp[key] = val[k]
		}
		resTmp["hostname"] = host
		res = append(res, resTmp)
	}
	return
}

func convertMap2Jsons(res []map[string]string) (resJsons [][]byte) {
	for _, valueMap := range res {
		js, err := json.Marshal(valueMap)
		if err != nil {
			log.Printf("convertMap2Jsons err convert to json. %v\n", err)
			continue
		}
		var nvlog NVlog
		err = json.Unmarshal(js, &nvlog)
		if err != nil {
			log.Printf("convertMap2Jsons err convert from json to nvlog. %v\n", err)
			continue
		}
		nvlogRes := NVlogRes{
			Timestamp: time.Now().Format(time.RFC3339),
			GPU:       conv2Int(nvlog.GPU),
			PWR:       conv2Int(nvlog.PWR),
			GTemp:     conv2Int(nvlog.GTemp),
			MTemp:     conv2Int(nvlog.MTemp),
			SM:        conv2Int(nvlog.SM),
			Mem:       conv2Int(nvlog.Mem),
			Enc:       conv2Int(nvlog.Enc),
			Dec:       conv2Int(nvlog.Dec),
			MClk:      conv2Int(nvlog.MClk),
			PClk:      conv2Int(nvlog.PClk),
			Hostname:  nvlog.Hostname,
		}

		js, err = json.Marshal(nvlogRes)
		if err != nil {
			log.Printf("convertMap2Jsons nvlogRes err convert to json. %v\n", err)
			continue
		}
		resJsons = append(resJsons, js)
	}
	return
}

func parseValue(data string) (resKey []string) {
	keys := strings.Split(data, " ")
	for _, value := range keys {
		tmpValue := strings.ReplaceAll(value, " ", "")
		if tmpValue != "" {
			resKey = append(resKey, tmpValue)
		}
	}
	return
}

func conv2Int(data string) int {
	res, err := strconv.Atoi(data)
	if err != nil {
		return 0
	}
	return res
}
