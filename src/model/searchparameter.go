package model

// TODO: think in better way to pass parameters to the queries
type SearchParameter struct {
	Field  string
	Values []interface{}
}
