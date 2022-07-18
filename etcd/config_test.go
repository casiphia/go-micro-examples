/*
 * @Descripttion: 828 project
 * @Version: 1.0.0
 * @Author: wangxiaodiao
 * @Email: 413586280@qq.com
 * @Date: 2022-07-18 10:31:52
 * @LastEditors: wangxiaodiao
 * @LastEditTime: 2022-07-18 11:47:44
 * @FilePath: /go-micro-examples/etcd/config_test.go
 */
package main

import (
	"reflect"
	"testing"
)

func Test_getConfig(t *testing.T) {
	tests := []struct {
		name string
		want Demo
	}{
		{
			"/micro/config",
			Demo{
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

func Test_getJSON(t *testing.T) {
	type args struct {
		pr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"getJSON",
			args{
				pr: "demo",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := getJSON(tt.args.pr); (err != nil) != tt.wantErr {
				t.Errorf("getJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func Test_watch(t *testing.T) {
	type args struct {
		pr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"watch",
			args{
				pr: "demo",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := watch(tt.args.pr); (err != nil) != tt.wantErr {
				t.Errorf("watch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
