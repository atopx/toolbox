package utils

import (
	"testing"
)

func TestMACDecode(t *testing.T) {
	type args struct {
		mac string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "MACDecode",
			args: args{mac: "00155D554FCE"},
			want: "00:15:5D:55:4F:CE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MACDecode(tt.args.mac); got != tt.want {
				t.Errorf("MACDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMACEncode(t *testing.T) {
	type args struct {
		mac string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "MACEncode[-]",
			args: args{mac: "00-15-5D-55-4F-CE"},
			want: "00155D554FCE",
		},
		{
			name: "MACEncode[:]",
			args: args{mac: "00:15:5D:55:4F:CE"},
			want: "00155D554FCE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MACEncode(tt.args.mac); got != tt.want {
				t.Errorf("MACEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}
