package run

import (
	"fmt"
	"testing"
)

func TestGetGT(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantPic []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				id: "67780",
			},
			wantPic: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPic, _ := GetGT(tt.args.id)
			for i := range gotPic {
				fmt.Println(i, gotPic[i])
			}
			err := DownloadPic("成都", gotPic)
			if err != nil {
				fmt.Println(err)
			}

		})
	}
}
