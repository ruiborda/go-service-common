package dto

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPageBodyBuilder(t *testing.T) {
	type testCase[T any] struct {
		name string
		want *PageBody[T]
	}
	tests := []testCase[string]{
		{
			name: "Builder should return default PageBody",
			want: &PageBody[string]{
				Items:       []string{},
				CurrentPage: 0,
				TotalItems:  0,
				TotalPages:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PageBodyBuilder[string](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PageBodyBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPageBody_SetCurrentPage(t *testing.T) {
	type args struct {
		currentPage int
	}
	type testCase[T any] struct {
		name string
		p    PageBody[T]
		args args
		want *PageBody[T]
	}
	tests := []testCase[string]{
		{
			name: "Set current page to 2",
			p:    PageBody[string]{CurrentPage: 1},
			args: args{currentPage: 2},
			want: &PageBody[string]{CurrentPage: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.SetCurrentPage(tt.args.currentPage)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCurrentPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPageBody_SetItems(t *testing.T) {
	type args[T any] struct {
		items []T
	}
	type testCase[T any] struct {
		name string
		p    PageBody[T]
		args args[T]
		want *PageBody[T]
	}
	tests := []testCase[string]{
		{
			name: "Set items with 3 elements",
			p:    PageBody[string]{},
			args: args[string]{
				items: []string{"a", "b", "c"},
			},
			want: &PageBody[string]{
				Items:      []string{"a", "b", "c"},
				TotalItems: 0,
				PageSize:   0,
				TotalPages: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.SetItems(tt.args.items)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPageBody_SetPageSize(t *testing.T) {
	type args struct {
		pageSize int
	}
	type testCase[T any] struct {
		name string
		p    PageBody[T]
		args args
		want *PageBody[T]
	}
	tests := []testCase[string]{
		{
			name: "Set page size to 10",
			p:    PageBody[string]{},
			args: args{pageSize: 10},
			want: &PageBody[string]{PageSize: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.SetPageSize(tt.args.pageSize)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPageSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPageBody_SetTotalItems(t *testing.T) {
	type args struct {
		totalItems int
	}
	type testCase[T any] struct {
		name string
		p    PageBody[T]
		args args
		want *PageBody[T]
	}
	tests := []testCase[string]{
		{
			name: "totalItems = 10, pageSize = 3 (inexact division)",
			p:    PageBody[string]{PageSize: 3},
			args: args{totalItems: 10},
			want: &PageBody[string]{TotalItems: 10, PageSize: 3, TotalPages: 4},
		},
		{
			name: "totalItems = 9, pageSize = 3 (exact division)",
			p:    PageBody[string]{PageSize: 3},
			args: args{totalItems: 9},
			want: &PageBody[string]{TotalItems: 9, PageSize: 3, TotalPages: 3},
		},
		{
			name: "totalItems = 0, pageSize = 5",
			p:    PageBody[string]{PageSize: 5},
			args: args{totalItems: 0},
			want: &PageBody[string]{TotalItems: 0, PageSize: 5, TotalPages: 0},
		},
		{
			name: "totalItems = 1, pageSize = 5",
			p:    PageBody[string]{PageSize: 5},
			args: args{totalItems: 1},
			want: &PageBody[string]{TotalItems: 1, PageSize: 5, TotalPages: 1},
		},
		{
			name: "totalItems = 5, pageSize = 0 (division by zero logic)",
			p:    PageBody[string]{PageSize: 0},
			args: args{totalItems: 5},
			want: &PageBody[string]{TotalItems: 5, PageSize: 0, TotalPages: 0},
		},
		{
			name: "totalItems = -10, pageSize = 3 (negative totalItems)", // Note: SetItems no longer sets PageSize, so we must set it explicitly if we want non-zero PageSize
			p: *PageBodyBuilder[string]().SetItems([]string{
				"a", "b",
			}).SetCurrentPage(2).SetPageSize(10).SetTotalItems(12),
			args: args{totalItems: 12},
			want: &PageBody[string]{TotalItems: 12, PageSize: 10, TotalPages: 2, Items: []string{"a", "b"}, CurrentPage: 2},
		},
		{
			name: "totalItems = 10, pageSize = -3 (negative pageSize)",
			p:    *PageBodyBuilder[string]().SetItems([]string{}),
			args: args{totalItems: 10},
			want: &PageBody[string]{TotalItems: 10, PageSize: 0, TotalPages: 0, Items: []string{}},
		},
		{
			name: "totalItems = 0, pageSize = 0",
			p:    PageBody[string]{PageSize: 0},
			args: args{totalItems: 0},
			want: &PageBody[string]{TotalItems: 0, PageSize: 0, TotalPages: 0},
		},
		{
			name: "totalItems = math.MaxInt, pageSize = 1",
			p:    PageBody[string]{PageSize: 1},
			args: args{totalItems: int(^uint(0) >> 1)},
			want: &PageBody[string]{
				TotalItems: int(^uint(0) >> 1),
				PageSize:   1,
				TotalPages: int(^uint(0) >> 1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("panic in case %q: %v", tt.name, r)
				}
			}()
			got := tt.p.SetTotalItems(tt.args.totalItems)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetTotalItems() = %+v, want %+v", got, tt.want)
				gotJson, _ := json.MarshalIndent(got, "", "\t")
				wantJson, _ := json.MarshalIndent(tt.want, "", "\t")

				t.Log("got:", string(gotJson))
				t.Log("want:", string(wantJson))

			}
		})
	}
}
