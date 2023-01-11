package netx_test

import (
	"github.com/shixinshuiyou/framework/netx"
	"testing"
)

func TestIp2long(t *testing.T) {
	type args struct {
		ipstr string
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "common",
			args: args{ipstr: "192.0.34.166"},
			want: 3221234342,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := netx.Ip2long(tt.args.ipstr); got != tt.want {
				t.Errorf("Ip2long() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLong2ip(t *testing.T) {
	type args struct {
		ipLong uint32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "common", args: args{ipLong: 3221234342}, want: "192.0.34.166"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := netx.Long2ip(tt.args.ipLong); got != tt.want {
				t.Errorf("Long2ip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIDC(t *testing.T) {
	t.Log(netx.IDC())
}

func TestInternalIp(t *testing.T) {
	t.Log(netx.InternalIp())
}

func TestOs(t *testing.T) {
	t.Log(netx.Os())
}
