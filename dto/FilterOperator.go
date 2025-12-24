package dto

type FilterOperator string

const (
	LessThan           FilterOperator = "<"
	LessThanOrEqual    FilterOperator = "<="
	Equals             FilterOperator = "=="
	GreaterThan        FilterOperator = ">"
	GreaterThanOrEqual FilterOperator = ">="
	NotEquals          FilterOperator = "!="
	ArrayContains      FilterOperator = "array-contains"
	ArrayContainsAny   FilterOperator = "array-contains-any"
	In                 FilterOperator = "in"
	NotIn              FilterOperator = "not-in"
)
