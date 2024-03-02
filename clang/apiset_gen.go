package clang

// #include "./clang-c/Documentation.h"
// #include "go-clang.h"
import "C"

// APISet cXAPISet is an opaque type that represents a data structure containing all the API information for a given translation unit. This can be used for a single symbol symbol graph for a given symbol.
type APISet struct {
	c C.CXAPISet
}

// DisposeAPISet dispose of an APISet.
//
// The provided CXAPISet can not be used after this function is called.
func (apis APISet) Dispose() {
	C.clang_disposeAPISet(apis.c)
}
