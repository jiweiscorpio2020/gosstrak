// Copyright (c) 2017 Iori Mizutani
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package filtering

import (
	"reflect"
	"testing"
)

func TestList_MarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		list    *List
		want    []byte
		wantErr bool
	}{
		{
			"simple marshal",
			&List{
				&ExactMatch{"3", NewFilter("0011", 0)},
				&ExactMatch{"3-0", NewFilter("00110000", 0)},
			},
			[]byte{24, 12, 0, 21, 69, 110, 103, 105, 110, 101, 58, 102, 105, 108, 116, 101, 114, 105, 110, 103, 46, 76, 105, 115, 116, 3, 4, 0, 4, 4, 12, 0, 1, 51, 113, 255, 129, 3, 1, 1, 12, 70, 105, 108, 116, 101, 114, 79, 98, 106, 101, 99, 116, 1, 255, 130, 0, 1, 7, 1, 6, 83, 116, 114, 105, 110, 103, 1, 12, 0, 1, 4, 83, 105, 122, 101, 1, 4, 0, 1, 6, 79, 102, 102, 115, 101, 116, 1, 4, 0, 1, 10, 66, 121, 116, 101, 70, 105, 108, 116, 101, 114, 1, 10, 0, 1, 8, 66, 121, 116, 101, 77, 97, 115, 107, 1, 10, 0, 1, 10, 66, 121, 116, 101, 79, 102, 102, 115, 101, 116, 1, 4, 0, 1, 8, 66, 121, 116, 101, 83, 105, 122, 101, 1, 4, 0, 0, 0, 19, 255, 130, 1, 4, 48, 48, 49, 49, 1, 8, 2, 1, 63, 1, 1, 15, 2, 2, 0, 6, 12, 0, 3, 51, 45, 48, 23, 255, 130, 1, 8, 48, 48, 49, 49, 48, 48, 48, 48, 1, 16, 2, 1, 48, 1, 1, 0, 2, 2, 0},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.list.MarshalBinary()
			if (err != nil) != tt.wantErr {
				t.Errorf("List.MarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.MarshalBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_UnmarshalBinary(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		list    *List
		args    args
		wantErr bool
	}{
		{
			"simple unmarshal",
			&List{
				&ExactMatch{"3", NewFilter("0011", 0)},
				&ExactMatch{"3-0", NewFilter("00110000", 0)},
			},
			args{
				[]byte{24, 12, 0, 21, 69, 110, 103, 105, 110, 101, 58, 102, 105, 108, 116, 101, 114, 105, 110, 103, 46, 76, 105, 115, 116, 3, 4, 0, 4, 4, 12, 0, 1, 51, 113, 255, 129, 3, 1, 1, 12, 70, 105, 108, 116, 101, 114, 79, 98, 106, 101, 99, 116, 1, 255, 130, 0, 1, 7, 1, 6, 83, 116, 114, 105, 110, 103, 1, 12, 0, 1, 4, 83, 105, 122, 101, 1, 4, 0, 1, 6, 79, 102, 102, 115, 101, 116, 1, 4, 0, 1, 10, 66, 121, 116, 101, 70, 105, 108, 116, 101, 114, 1, 10, 0, 1, 8, 66, 121, 116, 101, 77, 97, 115, 107, 1, 10, 0, 1, 10, 66, 121, 116, 101, 79, 102, 102, 115, 101, 116, 1, 4, 0, 1, 8, 66, 121, 116, 101, 83, 105, 122, 101, 1, 4, 0, 0, 0, 19, 255, 130, 1, 4, 48, 48, 49, 49, 1, 8, 2, 1, 63, 1, 1, 15, 2, 2, 0, 6, 12, 0, 3, 51, 45, 48, 23, 255, 130, 1, 8, 48, 48, 49, 49, 48, 48, 48, 48, 1, 16, 2, 1, 48, 1, 1, 0, 2, 2, 0},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.list.UnmarshalBinary(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("List.UnmarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestList_AnalyzeLocality(t *testing.T) {
	type args struct {
		id     []byte
		prefix string
		lm     *LocalityMap
	}
	tests := []struct {
		name string
		list *List
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.AnalyzeLocality(tt.args.id, tt.args.prefix, tt.args.lm)
		})
	}
}

func TestList_Search(t *testing.T) {
	type args struct {
		id []byte
	}
	tests := []struct {
		name        string
		list        *List
		args        args
		wantMatches []string
	}{
		{
			"0011xxxx on []byte{60, 128}",
			&List{
				&ExactMatch{"3", NewFilter("0011", 0)},
			},
			args{[]byte{60, 128}},
			[]string{"3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMatches := tt.list.Search(tt.args.id); !reflect.DeepEqual(gotMatches, tt.wantMatches) {
				t.Errorf("List.Search() = %v, want %v", gotMatches, tt.wantMatches)
			}
		})
	}
}

func TestBuildList(t *testing.T) {
	type args struct {
		sub Subscriptions
	}
	tests := []struct {
		name string
		args args
		want *List
	}{
		{
			"BuildList testing...",
			args{
				Subscriptions{
					"0011":         &Info{0, "3", 10, nil},
					"1111":         &Info{0, "15", 2, nil},
					"00110000":     &Info{0, "3-0", 5, nil},
					"001100110000": &Info{0, "3-3-0", 5, nil},
				},
			},
			&List{
				&ExactMatch{"3", NewFilter("0011", 0)},
				&ExactMatch{"3-0", NewFilter("00110000", 0)},
				&ExactMatch{"3-3-0", NewFilter("001100110000", 0)},
				&ExactMatch{"15", NewFilter("1111", 0)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildList(tt.args.sub); !reflect.DeepEqual(got, tt.want) {
				for i, em := range *got {
					if !reflect.DeepEqual(em.filter, (*tt.want)[i].filter) {
						t.Errorf("(*BuildList())[%v].filter = \n%v, want \n%v", i, em.filter, (*tt.want)[i].filter)
					} else if em.notificationURI != (*tt.want)[i].notificationURI {
						t.Errorf("(*BuildList())[%v].notificationURI = \n%v, want \n%v", i, em.notificationURI, (*tt.want)[i].notificationURI)
					}
				}
			}
		})
	}
}

func TestList_Dump(t *testing.T) {
	tests := []struct {
		name string
		list *List
		want string
	}{
		{"List.Dump() test", &List{
			&ExactMatch{"3", NewFilter("0011", 0)},
			&ExactMatch{"3-0", NewFilter("00110000", 0)},
		}, "--0011(0 4) 3\n--00110000(0 8) 3-0\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.Dump(); got != tt.want {
				t.Errorf("List.Dump() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}