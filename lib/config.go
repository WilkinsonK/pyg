package pyg

// #cgo pkg-config: python3-embed
// #include "txllayer.h"
import "C"

// Initialize a new Go PyConfig object.
func ConfigNew() *PyConfig {
	return &PyConfig{CInstance: CPyConfig{}, CInstanceMapped: false}
}

// Initialize a new Go PyPreConfig object.
func PreConfigNew() *PyPreConfig {
	return &PyPreConfig{CInstance: CPyPreConfig{}}
}

func (config *PyConfig) initFieldRefs() {
	if config.CInstanceMapped {
		return
	}

	config.Argv = config.CInstance.Argv
	config.BaseExecPrefix = config.CInstance.BaseExecPrefix
	config.BaseExecutable = config.CInstance.BaseExecutable
	config.BasePrefix = config.CInstance.BasePrefix
	config.BufferedStdio = config.CInstance.BufferedStdio
	config.BytesWarning = config.CInstance.BytesWarning
	config.CheckHashPYCsMode = config.CInstance.CheckHashPYCsMode
	config.ConfigureCStdio = config.CInstance.ConfigureCStdio
	config.DevMode = config.CInstance.DevMode
	config.DumpRefs = config.CInstance.DumpRefs
	config.ExecPrefix = config.CInstance.ExecPrefix
	config.Executable = config.CInstance.Executable
	config.FaultHandler = config.CInstance.FaultHandler
	config.FaultHandler = config.CInstance.FaultHandler
	config.FilesystemEncoding = config.CInstance.FilesystemEncoding
	config.FilesystemErrors = config.CInstance.FilesystemErrors
	config.Home = config.CInstance.Home
	config.ImportTime = config.CInstance.ImportTime
	config.Inspect = config.CInstance.Inspect
	config.InstallSignalHandlers = config.CInstance.InstallSignalHandlers
	config.Interactive = config.CInstance.Interactive
	config.Isolated = config.CInstance.Isolated
	config.MallocStats = config.CInstance.MallocStats
	config.PythonPathEnv = config.CInstance.PythonPathEnv
	config.ModuleSearchPaths = config.CInstance.ModuleSearchPaths
	config.ModuleSearchPathsSet = config.CInstance.ModuleSearchPathsSet
	config.OptimizationLevel = config.CInstance.OptimizationLevel
	config.ParseArgv = config.CInstance.ParseArgv
	config.ParserDebug = config.CInstance.ParserDebug
	config.PathConfigWarnings = config.CInstance.PathConfigWarnings
	config.Prefix = config.CInstance.Prefix
	config.ProgramName = config.CInstance.ProgramName
	config.PyCachePrefix = config.CInstance.PyCachePrefix
	config.Quiet = config.CInstance.Quiet
	config.RunCommand = config.CInstance.RunCommand
	config.RunFileName = config.CInstance.RunFileName
	config.RunModule = config.CInstance.RunModule
	config.ShowRefCount = config.CInstance.ShowRefCount
	config.SiteImport = config.CInstance.SiteImport
	config.SkipSourceFirstLine = config.CInstance.SkipSourceFirstLine
	config.StdioEncoding = config.CInstance.StdioEncoding
	config.StdioErrors = config.CInstance.StdioErrors
	config.TraceMalloc = config.CInstance.TraceMalloc
	config.UseEnvironment = config.CInstance.UseEnvironment
	config.UseHashSeed = config.CInstance.UseHashSeed
	config.UserSiteDirectory = config.CInstance.UserSiteDirectory
	config.Verbose = config.CInstance.Verbose
	config.WarnOptions = config.CInstance.WarnOptions
	config.WriteByteCode = config.CInstance.WriteByteCode
	config.XOptions = config.CInstance.XOptions

	config.CInstanceMapped = !config.CInstanceMapped
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
	config.initFieldRefs()
}

// Initialize configuration with the Python
// configuration.
// https://docs.python.org/3/c-api/init_config.html#c.PyConfig_InitPythonConfig
func (config *PyConfig) InitPythonConfig() {
	C.CGO_PyConfig_InitPythonConfig(&config.CInstance)
	config.initFieldRefs()
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
func (config *PyConfig) SetBoolean(prop *Cint, val bool) PyStatus {
	*(prop) = Bool2CInt(val)
	return StatusOk()
}

// Set property as integer.
func (config *PyConfig) SetInteger(prop *Cint, val int) PyStatus {
	*(prop) = Cint(val)
	return StatusOk()
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

func (config *PyPreConfig) initFieldRefs() {
	if config.CInstanceMapped {
		return
	}

	config.Allocator = config.CInstance.Allocator
	config.ConfigureLocale = config.CInstance.ConfigureLocale
	config.CoerceCLocale = config.CInstance.CoerceCLocale
	config.DevMode = config.CInstance.DevMode
	config.Isolated = config.CInstance.Isolated
	config.ParseArgv = config.CInstance.ParseArgv
	config.UseEnvironment = config.CInstance.UseEnvironment
	config.UTF8Mode = config.CInstance.UTF8Mode

	config.CInstanceMapped = !config.CInstanceMapped
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
	config.initFieldRefs()
}

// Initialize the preconfiguration with Isolated
// Configuration.
// https://docs.python.org/3/c-api/init_config.html#c.PyPreConfig_InitIsolatedConfig
func (config *PyPreConfig) InitIsolatedConfig() {
	C.CGO_PyPreConfig_InitIsolatedConfig(&config.CInstance)
	config.initFieldRefs()
}

// Set property as boolean.
func (config *PyPreConfig) SetBoolean(prop *Cint, val bool) PyStatus {
	*(prop) = Bool2CInt(val)
	return StatusOk()
}

// Set property as integer.
func (config *PyPreConfig) SetInteger(prop *Cint, val int) PyStatus {
	*(prop) = Cint(val)
	return StatusOk()
}
