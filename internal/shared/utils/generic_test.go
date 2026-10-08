package utils

import (
	"reflect"
	"testing"
)

func TestPaginationOffsetLimit(t *testing.T) {
	type args struct {
		page  int
		limit int
	}
	tests := []struct {
		name       string
		args       args
		wantOffset int
		wantLimit  int
	}{
		{
			name: "Test case 1 - page: 1 limit: 10 => offset: 00 limit: 10",
			args: args{
				page:  1,
				limit: 10,
			},
			wantOffset: 0,
			wantLimit:  10,
		},
		{
			name: "Test case 2 - page: 0 limit: 10 => offset: 0 limit: 10",
			args: args{
				page:  0,
				limit: 10,
			},
			wantOffset: 0,
			wantLimit:  10,
		},
		{
			name: "Test case 3 - page: 2 limit: 10 => offset: 10 limit: 20",
			args: args{
				page:  2,
				limit: 10,
			},
			wantOffset: 10,
			wantLimit:  20,
		},
		{
			name: "Test case 4 - page: 3 limit: 10 => offset: 20 limit: 30",
			args: args{
				page:  3,
				limit: 10,
			},
			wantOffset: 20,
			wantLimit:  30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := PaginationOffsetLimit(tt.args.page, tt.args.limit)
			if got != tt.wantOffset {
				t.Errorf("PaginationOffsetLimit() got = %v, want %v", got, tt.wantOffset)
			}
			if got1 != tt.wantLimit {
				t.Errorf("PaginationOffsetLimit() got1 = %v, want %v", got1, tt.wantLimit)
			}
		})
	}
}

func TestRandomInt(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		//want int
	}{
		{
			name: "Test case 1 - min: 0 max: 10",
			args: args{
				min: 0,
				max: 10,
			},
		},
		{
			name: "Test case 2 - min: 0 max: 15",
			args: args{
				min: 2,
				max: 15,
			},
		},
		{
			name: "Test case 13 - min: 0 max: 14",
			args: args{
				min: 13,
				max: 14,
			},
		},
		{
			name: "Test case -1 - min: 0 max: 5",
			args: args{
				min: -1,
				max: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomInt(tt.args.min, tt.args.max)

			if got < tt.args.min || got > tt.args.max {
				t.Errorf("Wrong Value for RandomInt() = %v min = %v max = %v", got, tt.args.min, tt.args.max)
			}
		})
	}
}

func TestDecodeBase64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Test case 1 - s: dGVzdA==",
			args: args{
				s: "dGVzdA==",
			},
			want:    []byte("test"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeBase64(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeBase64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}
