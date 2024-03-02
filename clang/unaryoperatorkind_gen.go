package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// UnaryOperatorKind describes the kind of unary operators.
type UnaryOperatorKind uint32

const (
	// UnaryOperator_Invalid this value describes cursors which are not unary operators.
	UnaryOperator_Invalid UnaryOperatorKind = C.CXUnaryOperator_Invalid
	// UnaryOperator_PostInc postfix increment operator.
	UnaryOperator_PostInc = C.CXUnaryOperator_PostInc
	// UnaryOperator_PostDec postfix decrement operator.
	UnaryOperator_PostDec = C.CXUnaryOperator_PostDec
	// UnaryOperator_PreInc prefix increment operator.
	UnaryOperator_PreInc = C.CXUnaryOperator_PreInc
	// UnaryOperator_PreDec prefix decrement operator.
	UnaryOperator_PreDec = C.CXUnaryOperator_PreDec
	// UnaryOperator_AddrOf address of operator.
	UnaryOperator_AddrOf = C.CXUnaryOperator_AddrOf
	// UnaryOperator_Deref dereference operator.
	UnaryOperator_Deref = C.CXUnaryOperator_Deref
	// UnaryOperator_Plus plus operator.
	UnaryOperator_Plus = C.CXUnaryOperator_Plus
	// UnaryOperator_Minus minus operator.
	UnaryOperator_Minus = C.CXUnaryOperator_Minus
	// UnaryOperator_Not not operator.
	UnaryOperator_Not = C.CXUnaryOperator_Not
	// UnaryOperator_LNot lNot operator.
	UnaryOperator_LNot = C.CXUnaryOperator_LNot
	// UnaryOperator_Real "__real expr" operator.
	UnaryOperator_Real = C.CXUnaryOperator_Real
	// UnaryOperator_Imag "__imag expr" operator.
	UnaryOperator_Imag = C.CXUnaryOperator_Imag
	// UnaryOperator_Extension __extension__ marker operator.
	UnaryOperator_Extension = C.CXUnaryOperator_Extension
	// UnaryOperator_Coawait c++ co_await operator.
	UnaryOperator_Coawait = C.CXUnaryOperator_Coawait
)

// GetUnaryOperatorKindSpelling retrieve the spelling of a given CXUnaryOperatorKind.
func (uok UnaryOperatorKind) Spelling() string {
	o := cxstring{C.clang_getUnaryOperatorKindSpelling(C.enum_CXUnaryOperatorKind(uok))}
	defer o.Dispose()

	return o.String()
}

func (uok UnaryOperatorKind) String() string {
	return uok.Spelling()
}
