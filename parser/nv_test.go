// created by D. "Mordok" Fedorov

package parser

import (
	"log"
	"os"
	"testing"
)

func TestPrivParseNVSMI(t *testing.T) {
	type args struct {
		testFile string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				testFile: "../test/nvidia.log",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := os.ReadFile(tt.args.testFile)
			if err != nil {
				t.Errorf("TestParseNVSMI(). %v", err)
			}
			parseNVSMI(string(data), "node1")
		})
	}
}

func TestParseNVSMI(t *testing.T) {
	type args struct {
		testFile string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				testFile: "../test/nvidia.log",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := os.ReadFile(tt.args.testFile)
			if err != nil {
				t.Errorf("TestParseNVSMI(). %v", err)
			}
			res := ParseNVSMI(string(data), "node1")
			for _, v := range res {
				log.Printf("%s\n", v)
			}
		})
	}
}
