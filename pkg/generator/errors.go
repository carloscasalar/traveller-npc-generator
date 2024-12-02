package generator

//go:generate errgen -type=generatorErrors --out-file=errors_auto.go
type generatorErrors struct {
	EmptyListErr error `errmsg:"invalid %v list: should contain at least one item" vars:"listName string"`
}
