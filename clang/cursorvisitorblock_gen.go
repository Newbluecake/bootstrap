package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

type CursorVisitorBlock struct {
	c C.CXCursorVisitorBlock
}
