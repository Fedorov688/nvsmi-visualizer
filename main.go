// created by D. "Mordok" Fedorov

package main

import (
	"flag"
	"fmt"
	"github.com/Fedorov688/nvsmi-visualizer/bd"
	"github.com/Fedorov688/nvsmi-visualizer/parser"
	"os"
	"os/exec"
	"time"
)

var esAddr *string
var esPrefix *string
var hostname *string
var pathNVSMI *string
var timeInterval *uint64
var testMode *bool

func init() {

	esAddr = flag.String(
		"esAddress",
		"http://localhost:9200",
		"ElasticSearch address",
	)

	esPrefix = flag.String(
		"esPrefix",
		"nvsmi-",
		"Prefix for ElasticSearch index",
	)
	host, _ := os.Hostname()

	hostname = flag.String(
		"host",
		host,
		"Hostname of this machine",
	)

	pathNVSMI = flag.String(
		"pathNVSMI",
		"nvidia-smi",
		"Path to nvidia-smi command",
	)

	timeInterval = flag.Uint64(
		"n",
		1,
		"seconds to wait between updates",
	)

	testMode = flag.Bool(
		"t",
		false,
		"seconds to wait between updates",
	)

	flag.Parse()
}

func main() {

	var es bd.ES
	es.Address = *esAddr
	es.Init()
	args := []string{"dmod", "-c 1", "-o T"}
	if *testMode {
		*pathNVSMI = "python3"
		args = []string{"test/nvsmi-generator.py"}
	}

	for {
		runCMD(&es, *pathNVSMI, args)
		time.Sleep(time.Duration(*timeInterval) * time.Second)
	}
}

func runCMD(es *bd.ES, command string, args []string) {
	out, err := exec.Command(command, args...).Output()
	if err != nil {
		fmt.Printf("Err cmd.Run(). %v\n", err)
		return
	}
	fmt.Println(string(out))
	res := parser.ParseNVSMI(string(out), *hostname)
	timestamp := time.Now().Format("02-01-2006")
	for _, v := range res {
		es.SendJson(
			*esPrefix+timestamp,
			v,
		)
	}
}
