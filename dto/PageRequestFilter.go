package dto

import (
	"encoding/json"
	"reflect"
)

type PageRequestFilter struct {
	Field    string         `json:"field" form:"field"`
	Operator FilterOperator `json:"operator" form:"operator"`
	Value    FilterValue    `json:"value" form:"value"`
}

type FilterValue []byte

func (fv *FilterValue) UnmarshalJSON(b []byte) error {
	*fv = b
	return nil
}

func (fv FilterValue) MarshalJSON() ([]byte, error) {
	return fv, nil
}

func (fv FilterValue) Export() interface{} {
	if len(fv) == 0 {
		return nil
	}

	var val interface{}
	if err := json.Unmarshal(fv, &val); err != nil {
		return nil
	}

	if f, ok := val.(float64); ok {
		if f == float64(int(f)) {
			return int(f)
		}
		return f
	}

	if slice, ok := val.([]interface{}); ok {
		if len(slice) == 0 {
			return slice
		}

		first := slice[0]
		firstType := reflect.TypeOf(first)
		isHomogeneous := true
		for _, item := range slice {
			if reflect.TypeOf(item) != firstType {
				isHomogeneous = false
				break
			}
		}

		if isHomogeneous {
			switch first.(type) {
			case string:
				ss := make([]string, len(slice))
				for i, v := range slice {
					ss[i] = v.(string)
				}
				return ss
			case bool:
				bs := make([]bool, len(slice))
				for i, v := range slice {
					bs[i] = v.(bool)
				}
				return bs
			case float64:
				allInts := true
				for _, v := range slice {
					f := v.(float64)
					if f != float64(int(f)) {
						allInts = false
						break
					}
				}
				if allInts {
					is := make([]int, len(slice))
					for i, v := range slice {
						is[i] = int(v.(float64))
					}
					return is
				}
				fs := make([]float64, len(slice))
				for i, v := range slice {
					fs[i] = v.(float64)
				}
				return fs
			}
		}
		return slice
	}

	return val
}

func ExportAs[T any](fv FilterValue) T {
	return fv.Export().(T)
}
