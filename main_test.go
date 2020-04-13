package main

import "testing"

func Test_getAddress(t *testing.T) {
	type args struct {
		hostName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"localhost", args{hostName: "localhost"}, "[\"127.0.0.1\"]", false},
		{"badname", args{hostName: "badname"}, "", true},
		{"ipaddress", args{hostName: "10.0.0.1"}, "[\"10.0.0.1\"]", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getAddress(tt.args.hostName)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
