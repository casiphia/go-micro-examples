package main

import (
	"reflect"
	"testing"
)

func Test_getConfig(t *testing.T) {
	tests := []struct {
		name string
		want Config
	}{
		{
			"/micro/config",
			Config{
				NetWork: "172.30.0.0/16",
				Backend: "vxlan",
			}, // TODO: Add test cases.
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
