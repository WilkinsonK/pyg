package pygmy

// #cgo pkg-config: python3-embed
// #include "txllayer.h"
import "C"
import "fmt"

// Get a new instance of Python interpreter with
// preset config and preconfig objects.
func PygmyNew(config *PyConfig, preconfig *PyPreConfig) *Pygmy {
	if config == nil {
		config = ConfigNew()
		config.InitPythonConfig()
	}
	if preconfig == nil {
		preconfig = PreConfigNew()
		preconfig.InitPythonConfig()
	}
	return &Pygmy{Config: config, PreConfig: preconfig}
}

// Get a new instance of Python interpreter.
func PygmyNewPython() *Pygmy {
	return PygmyNew(nil, nil)
}

// Get a new isolated instance of Python
// interpreter.
func InterpreterNewIsolated() *Pygmy {
	preconfig, config := PreConfigNew(), ConfigNew()
	preconfig.InitIsolatedConfig()
	config.InitIsolatedConfig()
	return PygmyNew(config, preconfig)
}

func (inter *Pygmy) checkInitialized() error {
	if !inter.IsInitialized() {
		return fmt.Errorf("interpreter not yet initialized")
	}
	return nil
}

func (inter *Pygmy) checkInitializedS(val string, err error) (string, error) {
	if err := inter.checkInitialized(); err != nil {
		return "", err
	}
	return val, err
}

func (inter *Pygmy) FinalizeEx() int {
	inter.Config.Clear()
	inter.PreConfig.Clear()
	return int(C.Py_FinalizeEx())
}

// Initialize Python from config configuration.
// https://docs.python.org/3/c-api/init_config.html#c.Py_InitializeFromConfig
func (inter *Pygmy) InitializeFromConfig() PyStatus {
	ret := C.CGO_Py_InitializeFromConfig(&inter.Config.CInstance)
	return PyStatusNew(ret)
}

// Whether the Python interpreter has been
// initialized or not.
// https://docs.python.org/3/c-api/init.html#c.Py_IsInitialized
func (inter *Pygmy) IsInitialized() bool {
	return CInt2Bool(C.Py_IsInitialized())
}

// Return information about the sequence number
// and build date and time of the current Python
// interpreter.
// https://docs.python.org/3/c-api/init.html#c.Py_GetBuildInfo
func (inter *Pygmy) GetBuildInfo() string {
	return C.GoString(C.Py_GetBuildInfo())
}

// Return an indication of the compiler used to
// build the current Python version.
// https://docs.python.org/3/c-api/init.html#c.Py_GetCompiler
func (inter *Pygmy) GetCompiler() string {
	return C.GoString(C.Py_GetCompiler())
}

// Return the official copyright for the current
// Python version.
// https://docs.python.org/3/c-api/init.html#c.Py_GetCopyright
func (inter *Pygmy) GetCopyright() string {
	return C.GoString(C.Py_GetCopyright())
}

// Return the exec-prefix for installed
// platform-dependent files.
// https://docs.python.org/3/c-api/init.html#c.Py_GetExecPrefix
func (inter *Pygmy) GetExecPrefix() (string, error) {
	return inter.checkInitializedS(WString2String(C.Py_GetExecPrefix()))
}

// Return the default module search path.
// https://docs.python.org/3/c-api/init.html#c.Py_GetPath
func (inter *Pygmy) GetPath() (string, error) {
	return inter.checkInitializedS(WString2String(C.Py_GetPath()))
}

// Return the platform identifier for the current
// platform.
// https://docs.python.org/3/c-api/init.html#c.Py_GetPlatform
func (inter *Pygmy) GetPlatform() string {
	return C.GoString(C.Py_GetPlatform())
}

// Return the prefix for installed platform-independent files.
// https://docs.python.org/3/c-api/init.html#c.Py_GetPrefix
func (inter *Pygmy) GetPrefix() (string, error) {
	return inter.checkInitializedS(WString2String(C.Py_GetPrefix()))
}

// Return the full program name of the Python
// executable.
// https://docs.python.org/3/c-api/init.html#c.Py_GetProgramFullPath
func (inter *Pygmy) GetProgramFullPath() (string, error) {
	return inter.checkInitializedS(WString2String(C.Py_GetProgramFullPath()))
}

// Return the program name set by Python
// configuration.
// https://docs.python.org/3/c-api/init.html#c.Py_GetProgramName
func (inter *Pygmy) GetProgramName() (string, error) {
	return inter.checkInitializedS(WString2String(C.Py_GetProgramName()))
}

// Return the default "home", that is, the value
// set by configuration.
// https://docs.python.org/3/c-api/init.html#c.Py_GetPythonHome
func (inter *Pygmy) GetPythonHome() (string, error) {
	return inter.checkInitializedS(WString2String(C.Py_GetPythonHome()))
}

// Return the version of this Python interpreter.
// https://docs.python.org/3/c-api/init.html#c.Py_GetVersion
func (inter *Pygmy) GetVersion() string {
	return C.GoString(C.Py_GetVersion())
}

// Preinitialize Python from preconfig
// preconfiguration.
// https://docs.python.org/3/c-api/init_config.html#c.Py_PreInitialize
func (inter *Pygmy) PreInitialize() PyStatus {
	ret := C.CGO_PreInitialize(&inter.PreConfig.CInstance)
	return PyStatusNew(ret)
}

// Preinitialize Python from preconfig
// preconfiguration. Parse argv command line
// arguments (wide strings) if parse_argv of
// PyPreConfig is non-zero.
// https://docs.python.org/3/c-api/init_config.html#c.Py_PreInitializeFromArgs
func (inter *Pygmy) PreInitializeFromArgs(argv []CPyWideString) PyStatus {
	ret := C.CGO_Py_PreInitializeFromArgs(
		&inter.PreConfig.CInstance,
		Clong(len(argv)),
		&argv[0])
	return PyStatusNew(ret)
}

// Preinitialize Python from preconfig
// preconfiguration. Parse argv command line
// arguments (bytes strings) if parse_argv of
// PyPreConfig is non-zero.
// https://docs.python.org/3/c-api/init_config.html#c.Py_PreInitializeFromBytesArgs
func (inter *Pygmy) PreInitializeFromBytesArgs(argv []string) PyStatus {
	cArgV := make([]*Cchar, len(argv))
	for idx, arg := range argv {
		carg := C.CString(arg)
		defer CFree(carg)
		cArgV[idx] = carg
	}

	ret := C.CGO_Py_PreInitializeFromBytesArgs(
		&inter.PreConfig.CInstance,
		Clong(len(argv)),
		&cArgV[0])
	return PyStatusNew(ret)
}

// Execute the command (PyConfig.run_command), the
// script (PyConfig.run_filename) or the module
// (PyConfig.run_module) specified on the command
// line or in the configuration.
// https://docs.python.org/3/c-api/init_config.html#c.Py_RunMain
func (inter *Pygmy) RunMain() (int, error) {
	if err := inter.checkInitialized(); err != nil {
		return 0, err
	}
	return int(C.Py_RunMain()), nil
}

func (inter *Pygmy) SetArgv(argv []string) PyStatus {
	return inter.Config.SetBytesArgv(argv)
}
