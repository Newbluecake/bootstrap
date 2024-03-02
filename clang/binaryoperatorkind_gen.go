package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// BinaryOperatorKind describes the kind of binary operators.
type BinaryOperatorKind uint32

const (
	// BinaryOperator_Invalid this value describes cursors which are not binary operators.
	BinaryOperator_Invalid BinaryOperatorKind = C.CXBinaryOperator_Invalid
	// BinaryOperator_PtrMemD c++ Pointer - to - member operator.
	BinaryOperator_PtrMemD = C.CXBinaryOperator_PtrMemD
	// BinaryOperator_PtrMemI c++ Pointer - to - member operator.
	BinaryOperator_PtrMemI = C.CXBinaryOperator_PtrMemI
	// BinaryOperator_Mul multiplication operator.
	BinaryOperator_Mul = C.CXBinaryOperator_Mul
	// BinaryOperator_Div division operator.
	BinaryOperator_Div = C.CXBinaryOperator_Div
	// BinaryOperator_Rem remainder operator.
	BinaryOperator_Rem = C.CXBinaryOperator_Rem
	// BinaryOperator_Add addition operator.
	BinaryOperator_Add = C.CXBinaryOperator_Add
	// BinaryOperator_Sub subtraction operator.
	BinaryOperator_Sub = C.CXBinaryOperator_Sub
	// BinaryOperator_Shl bitwise shift left operator.
	BinaryOperator_Shl = C.CXBinaryOperator_Shl
	// BinaryOperator_Shr bitwise shift right operator.
	BinaryOperator_Shr = C.CXBinaryOperator_Shr
	// BinaryOperator_Cmp c++ three-way comparison (spaceship) operator.
	BinaryOperator_Cmp = C.CXBinaryOperator_Cmp
	// BinaryOperator_LT less than operator.
	BinaryOperator_LT = C.CXBinaryOperator_LT
	// BinaryOperator_GT greater than operator.
	BinaryOperator_GT = C.CXBinaryOperator_GT
	// BinaryOperator_LE less or equal operator.
	BinaryOperator_LE = C.CXBinaryOperator_LE
	// BinaryOperator_GE greater or equal operator.
	BinaryOperator_GE = C.CXBinaryOperator_GE
	// BinaryOperator_EQ equal operator.
	BinaryOperator_EQ = C.CXBinaryOperator_EQ
	// BinaryOperator_NE not equal operator.
	BinaryOperator_NE = C.CXBinaryOperator_NE
	// BinaryOperator_And bitwise AND operator.
	BinaryOperator_And = C.CXBinaryOperator_And
	// BinaryOperator_Xor bitwise XOR operator.
	BinaryOperator_Xor = C.CXBinaryOperator_Xor
	// BinaryOperator_Or bitwise OR operator.
	BinaryOperator_Or = C.CXBinaryOperator_Or
	// BinaryOperator_LAnd logical AND operator.
	BinaryOperator_LAnd = C.CXBinaryOperator_LAnd
	// BinaryOperator_LOr logical OR operator.
	BinaryOperator_LOr = C.CXBinaryOperator_LOr
	// BinaryOperator_Assign assignment operator.
	BinaryOperator_Assign = C.CXBinaryOperator_Assign
	// BinaryOperator_MulAssign multiplication assignment operator.
	BinaryOperator_MulAssign = C.CXBinaryOperator_MulAssign
	// BinaryOperator_DivAssign division assignment operator.
	BinaryOperator_DivAssign = C.CXBinaryOperator_DivAssign
	// BinaryOperator_RemAssign remainder assignment operator.
	BinaryOperator_RemAssign = C.CXBinaryOperator_RemAssign
	// BinaryOperator_AddAssign addition assignment operator.
	BinaryOperator_AddAssign = C.CXBinaryOperator_AddAssign
	// BinaryOperator_SubAssign subtraction assignment operator.
	BinaryOperator_SubAssign = C.CXBinaryOperator_SubAssign
	// BinaryOperator_ShlAssign bitwise shift left assignment operator.
	BinaryOperator_ShlAssign = C.CXBinaryOperator_ShlAssign
	// BinaryOperator_ShrAssign bitwise shift right assignment operator.
	BinaryOperator_ShrAssign = C.CXBinaryOperator_ShrAssign
	// BinaryOperator_AndAssign bitwise AND assignment operator.
	BinaryOperator_AndAssign = C.CXBinaryOperator_AndAssign
	// BinaryOperator_XorAssign bitwise XOR assignment operator.
	BinaryOperator_XorAssign = C.CXBinaryOperator_XorAssign
	// BinaryOperator_OrAssign bitwise OR assignment operator.
	BinaryOperator_OrAssign = C.CXBinaryOperator_OrAssign
	// BinaryOperator_Comma comma operator.
	BinaryOperator_Comma = C.CXBinaryOperator_Comma
)

// GetBinaryOperatorKindSpelling retrieve the spelling of a given CXBinaryOperatorKind.
func (bok BinaryOperatorKind) Spelling() string {
	o := cxstring{C.clang_getBinaryOperatorKindSpelling(C.enum_CXBinaryOperatorKind(bok))}
	defer o.Dispose()

	return o.String()
}

func (bok BinaryOperatorKind) String() string {
	return bok.Spelling()
}
