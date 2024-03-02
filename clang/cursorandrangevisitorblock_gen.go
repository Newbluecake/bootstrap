package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

type CursorAndRangeVisitorBlock struct {
	c C.CXCursorAndRangeVisitorBlock
}
