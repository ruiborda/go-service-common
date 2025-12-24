package dto

import (
	"github.com/ruiborda/go-service-common/types"
	"reflect"
	"testing"
)

func TestDefaultPageRequest(t *testing.T) {
	type args struct {
		request *PageRequest
	}
	tests := []struct {
		name string
		args args
		want *PageRequest
	}{
		{
			name: "Default values for nil request",
			args: args{request: nil},
			want: &PageRequest{
				PageNumber: types.Pointer(1),
				PageSize:   types.Pointer(10),
				Search:     nil,
				Sort:       nil,
				Filters:    nil,
			},
		},
		{
			name: "Default values for empty request",
			args: args{request: &PageRequest{}},
			want: &PageRequest{
				PageNumber: types.Pointer(1),
				PageSize:   types.Pointer(10),
				Search:     nil,
				Sort:       nil,
				Filters:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultPageRequest(tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultPageRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
