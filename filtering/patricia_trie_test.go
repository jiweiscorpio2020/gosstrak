// Copyright (c) 2017 Iori Mizutani
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package filtering

import (
	"testing"
)

func Test_lcp(t *testing.T) {
	type args struct {
		l []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0001,0010,0011,0000", args{[]string{"0001", "0010", "0011", "0000"}}, "00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lcp(tt.args.l); got != tt.want {
				t.Errorf("lcp() = %v, want %v", got, tt.want)
			}
		})
	}
}
