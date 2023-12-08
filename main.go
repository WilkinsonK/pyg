package main

import (
	"os"

	pyg "github.com/WilkinsonK/pyg/lib"
)

func EvalStatus(status pyg.PyStatus) {
	if status.IsException() {
		if status.IsExit() {
			os.Exit(status.ExitCode())
		}
		status.ExitStatusException()
	}
}

func main() {
	pygi := pyg.PygNewPython()
	defer pygi.FinalizeEx()

	EvalStatus(pygi.SetPreConfigInteger(pygi.PreConfig.ParseArgv, 1))
	EvalStatus(pygi.SetPreConfigInteger(pygi.PreConfig.DevMode, 1))
	EvalStatus(pygi.PreInitialize())

	EvalStatus(pygi.SetConfigArgv(os.Args))
	EvalStatus(pygi.SetConfigInteger(pygi.Config.Verbose, 2))
	EvalStatus(pygi.SetConfigString(pygi.Config.ProgramName, "PyGi"))
	EvalStatus(pygi.InitializeFromConfig())

	pygi.RunString("print('Hello World')")
}
