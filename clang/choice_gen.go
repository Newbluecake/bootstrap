package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import "fmt"

type Choice uint32

const (
	// Choice_Default use the default value of an option that may depend on the process environment.
	Choice_Default Choice = C.CXChoice_Default
	// Choice_Enabled enable the option.
	Choice_Enabled = C.CXChoice_Enabled
	// Choice_Disabled disable the option.
	Choice_Disabled = C.CXChoice_Disabled
)

func (c Choice) Spelling() string {
	switch c {
	case Choice_Default:
		return "Choice=Default"
	case Choice_Enabled:
		return "Choice=Enabled"
	case Choice_Disabled:
		return "Choice=Disabled"
	}

	return fmt.Sprintf("Choice unknown %d", int(c))
}

func (c Choice) String() string {
	return c.Spelling()
}
