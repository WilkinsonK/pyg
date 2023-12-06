package pyg

// #cgo pkg-config: python3-embed
// #include "txllayer.h"
import "C"

// Initialize a new Go PyConfig object.
func ConfigNew() *PyConfig {
	return &PyConfig{CInstance: CPyConfig{}}
}

// Initialize a new Go PyPreConfig object.
func PreConfigNew() *PyPreConfig {
	return &PyPreConfig{CInstance: CPyPreConfig{}}
}

// Release configuration memory.
// https://docs.python.org/3/c-api/init_config.html#c.PyConfig_Clear
func (config *PyConfig) Clear() {
	C.CGO_PyConfig_Clear(&config.CInstance)
}

// Initialize configuration with the isolated
// configuration.
// https://docs.python.org/3/c-api/init_config.html#c.PyConfig_InitIsolatedConfig
func (config *PyConfig) InitIsolatedConfig() {
	C.CGO_PyConfig_InitIsolatedConfig(&config.CInstance)
}

// Initialize configuration with the Python
// configuration.
// https://docs.python.org/3/c-api/init_config.html#c.PyConfig_InitPythonConfig
func (config *PyConfig) InitPythonConfig() {
	C.CGO_PyConfig_InitPythonConfig(&config.CInstance)
}

// Read all Python configuration. Fields which are
// already initialized are left unchanged. Fields
// for path configuration are no longer calculated
// or modified when calling this function, as of
// Python 3.11.
//
// The PyConfig_Read function only parses
// `PyConfig.argv` arguments once:
// `PyConfig.parse_argv` is set to 2 after
// arguments are parsed. Since Python arguments
// are stripped from `PyConfig.argv`, parsing
// arguments twice would parse the application
// options as Python options.
// https://docs.python.org/3/c-api/init_config.html#c.PyConfig_Read
func (config *PyConfig) Read() PyStatus {
	ret := C.CGO_PyConfig_Read(&config.CInstance)
	return StatusNew(ret)
}

// Set command line arguments. (argv member of
// config) from the argv list of wide character
// strings
// https://docs.python.org/3/c-api/init_config.html#c.PyConfig_SetArgv
func (config *PyConfig) SetArgv(argv []CPyWideString) PyStatus {
	ret := C.CGO_PyConfig_SetArgv(&config.CInstance, Clong(len(argv)), &argv[0])
	return StatusNew(ret)
}

// Set command line arguments (argv member of
// config) from the argv list of bytes strings.
// Decode bytes using Py_DecodeLocale().
// https://docs.python.org/3/c-api/init_config.html#c.PyConfig_SetBytesArgv
func (config *PyConfig) SetBytesArgv(argv []string) PyStatus {
	cArgV := make([]*Cchar, len(argv))
	for idx, arg := range argv {
		carg := C.CString(arg)
		defer CFree(carg)
		cArgV[idx] = carg
	}
	cArgC := Clong(len(argv))

	ret := C.CGO_PyConfig_SetBytesArgv(&config.CInstance, cArgC, &cArgV[0])
	return StatusNew(ret)
}

// Decode str using Py_DecodeLocale() and set the
// result into configStr. Preinitialize Python if
// needed.
// https://docs.python.org/3/c-api/init_config.html#c.PyConfig_SetBytesString
func (config *PyConfig) SetBytesString(
	configStr *CPyWideString, str string) PyStatus {

	cStr := C.CString(str)
	defer CFree(cStr)

	ret := C.CGO_PyConfig_SetBytesString(&config.CInstance, configStr, cStr)
	return StatusNew(ret)
}

// Set property as boolean.
func (config *PyConfig) SetBoolean(prop *Cint, val bool) {
	*(prop) = Bool2CInt(val)
}

// Set property as integer.
func (config *PyConfig) SetInteger(prop *Cint, val int) {
	*(prop) = Cint(val)
}

// Copy wide character string into configStr.
// Preinitialize Python if needed.
// https://docs.python.org/3/c-api/init_config.html#c.PyConfig_SetString
func (config *PyConfig) SetString(
	prop *CPyWideString, val CPyWideString) PyStatus {

	ret := C.CGO_PyConfig_SetString(&config.CInstance, prop, val)
	return StatusNew(ret)
}

// Set the list of wide strings list to length
// and items. Preinitialize Python if needed.
// https://docs.python.org/3/c-api/init_config.html#c.PyConfig_SetWideStringList
func (config *PyConfig) SetWideStringList(
	list *CPyWideStringList,
	items []CPyWideString) PyStatus {

	ret := C.CGO_PyConfig_SetWideStringList(
		&config.CInstance,
		list,
		Clong(len(items)),
		&items[0])
	return StatusNew(ret)
}

// Free any memory associated with preconfig.
func (config *PyPreConfig) Clear() {
	C.CGO_PyPreConfig_Clear(&config.CInstance)
}

// Initialize the preconfiguration with Python
// Configuration.
// https://docs.python.org/3/c-api/init_config.html#c.PyPreConfig_InitPythonConfig
func (config *PyPreConfig) InitPythonConfig() {
	C.CGO_PyPreConfig_InitPythonConfig(&config.CInstance)
}

// Initialize the preconfiguration with Isolated
// Configuration.
// https://docs.python.org/3/c-api/init_config.html#c.PyPreConfig_InitIsolatedConfig
func (config *PyPreConfig) InitIsolatedConfig() {
	C.CGO_PyPreConfig_InitIsolatedConfig(&config.CInstance)
}

// Set property as integer.
func (config *PyPreConfig) SetInteger(prop *Cint, val int) {
	*(prop) = Cint(val)
}
