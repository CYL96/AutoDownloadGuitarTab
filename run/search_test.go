package run

import (
	"fmt"
	"testing"
)

func TestGetSearchResult(t *testing.T) {
	type args struct {
		search string
		p      string
	}
	tests := []struct {
		name     string
		args     args
		wantData []SearchResultExt
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				search: "成都",
				p:      "1",
			},
			wantData: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, _ := GetSearchResult(tt.args.search, tt.args.p)
			for i := range gotData {
				fmt.Println(i+1, "名称：", gotData[i].Name, "id:", gotData[i].Id)
			}
		})
	}
}
