package pyg

// #cgo pkg-config: python3-embed
// #include "txllayer.h"
import "C"
import (
	"strings"
	"unicode/utf16"
	"unsafe"
)

const maxRunes uint64 = 1<<30 - 1

type CPyCompilerFlags = C.PyCompilerFlags
type CPyConfig = C.CGO_PyConfig
type CPyObject = C.PyObject
type CPyPreConfig = C.CGO_PyPreConfig
type CPySSizeT = C.Py_ssize_t
type CPyStatus = C.PyStatus
type CPyWideString = *C.wchar_t
type CPyWideStringList = C.PyWideStringList
type Cchar = C.char
type CFILE = C.FILE
type Cint = C.int
type Clong = C.long
type Cvoid = C.void
type Cwchar = C.wchar_t

// Representative object of some object from the
// Python CAPI. Allows for related CPython methods
// to grouped under an object in Go per the object
// (or objects) they modify.
type PyArtifact[T comparable] struct {
	// Internal C object that this artifact points
	// to/represents.
	CInstance       T
	cInstanceMapped bool
}

type PyStatus PyArtifact[CPyStatus]
type PyPreConfig struct {
	PyArtifact[CPyPreConfig]
	CInstance       CPyPreConfig
	CInstanceMapped bool
	CPyPreConfig
}
type PyConfig struct {
	PyArtifact[CPyConfig]
	CInstance       CPyConfig
	CInstanceMapped bool
	CPyConfig
}

type Pyg struct {
	PreConfig *PyPreConfig
	Config    *PyConfig
}

// Translate a boolean to a C type integer.
func Bool2CInt(obj bool) Cint {
	if obj {
		return 1
	} else {
		return 0
	}
}

func CFree[T comparable](obj *T) {
	C.free(unsafe.Pointer(obj))
}

// CInt2Bool
//
// Translate a C type integer into a boolean.
func CInt2Bool(obj Cint) bool {
	return obj > 0
}

// WString2String
//
// Translate a wide char string into a Go string.
func WString2String(value CPyWideString) string {
	wptr := unsafe.Pointer(value)
	size := C.wcslen((CPyWideString)(wptr)) * 2
	ret := string(utf16.Decode((*[maxRunes]uint16)(wptr)[:size:size]))

	// Have to clean the decoded string to remove
	// 0 byte values embedded in it.
	badChar := string(byte(0))
	for strings.Contains(ret, badChar) {
		idx := strings.Index(ret, badChar)
		ret = ret[:idx] + ret[idx+1:]
	}

	return ret
}

// String2WideString
//
// Translate a Go string into a wide char string.
func String2WideString(value string) CPyWideString {
	return (CPyWideString)(unsafe.Pointer(&utf16.Encode([]rune(value))[0]))
}
