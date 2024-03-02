package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// IndexOptions index initialization options.
//
// 0 is the default value of each member of this struct except for Size.
// Initialize the struct in one of the following three ways to avoid adapting
// code each time a new member is added to it:
//
//	CXIndexOptions Opts;
//	memset(&Opts, 0, sizeof(Opts));
//	Opts.Size = sizeof(CXIndexOptions);
//
// or explicitly initialize the first data member and zero-initialize the rest:
//
//	CXIndexOptions Opts = { sizeof(CXIndexOptions) };
//
// or to prevent the -Wmissing-field-initializers warning for the above version:
//
//	CXIndexOptions Opts{};
//	Opts.Size = sizeof(CXIndexOptions);
type IndexOptions struct {
	c C.CXIndexOptions
}

// Size the size of struct Size used for option versioning.
//
// Always initialize this member to sizeof(Size), or assign
// sizeof(Size) to it right after creating a Size object.
func (io IndexOptions) Size() uint32 {
	return uint32(io.c.Size)
}

// ThreadBackgroundPriorityForIndexing a CXChoice enumerator that specifies the indexing priority policy. \sa CXGlobalOpt_ThreadBackgroundPriorityForIndexing
func (io IndexOptions) ThreadBackgroundPriorityForIndexing() uint8 {
	return uint8(io.c.ThreadBackgroundPriorityForIndexing)
}

// ThreadBackgroundPriorityForEditing a CXChoice enumerator that specifies the editing priority policy. \sa CXGlobalOpt_ThreadBackgroundPriorityForEditing
func (io IndexOptions) ThreadBackgroundPriorityForEditing() uint8 {
	return uint8(io.c.ThreadBackgroundPriorityForEditing)
}

// PreambleStoragePath the path to a directory, in which to store temporary PCH files. If null or
// empty, the default system temporary directory is used. These PCH files are
// deleted on clean exit but stay on disk if the program crashes or is killed.
//
// This option is ignored if \a StorePreamblesInMemory is non-zero.
//
// Libclang does not create the directory at the specified path in the file
// system. Therefore it must exist, or storing PCH files will fail.
func (io IndexOptions) PreambleStoragePath() string {
	return C.GoString(io.c.PreambleStoragePath)
}

// InvocationEmissionPath specifies a path which will contain log files for certain libclang invocations. A null value implies that libclang invocations are not logged.
func (io IndexOptions) InvocationEmissionPath() string {
	return C.GoString(io.c.InvocationEmissionPath)
}
