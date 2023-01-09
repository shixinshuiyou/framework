package netx_test

import (
	"framework/netx"
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

func TestArch(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := netx.Arch(); got != tt.want {
				t.Errorf("Arch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHostname(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := netx.Hostname(); got != tt.want {
				t.Errorf("Hostname() = %v, want %v", got, tt.want)
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
