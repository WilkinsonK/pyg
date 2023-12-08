package main

import (
	"testing"

	pyg "github.com/WilkinsonK/pyg/lib"
)

// Validates that 2 separate pointer variables
// point to the same address. This ensures that
// if the value is changed from one, the other
// will also change.
func testAttrP[P comparable](t *testing.T, name string, outer *P, inner *P) {
	if !(outer == inner) {
		t.Errorf(
			"top-level object attr %s is not mapped to C instance attr",
			name)
	}
}

func testEqualS[P string](t *testing.T, a, b, message string) {
	if !(a == b) {
		t.Errorf("expected [%s] got [%s]. %s", a, b, message)
	}
}

func TestConfigSetup(t *testing.T) {
	config := pyg.ConfigNew()
	defer config.Clear()

	config.InitIsolatedConfig()
	if !config.CInstanceMapped {
		t.Error("C instance attributes were not mapped to top-level object")
	}

	testAttrP(t, "Argv", config.Argv, config.CInstance.Argv)
	testAttrP(t, "BaseExecPrefix", config.BaseExecPrefix, config.CInstance.BaseExecPrefix)
	testAttrP(t, "BaseExecutable", config.BaseExecutable, config.CInstance.BaseExecutable)
	testAttrP(t, "BasePrefix", config.BasePrefix, config.CInstance.BasePrefix)
	testAttrP(t, "BufferedStdio", config.BufferedStdio, config.CInstance.BufferedStdio)
	testAttrP(t, "BytesWarning", config.BytesWarning, config.CInstance.BytesWarning)
	testAttrP(t, "CheckHashPYCsMode", config.CheckHashPYCsMode, config.CInstance.CheckHashPYCsMode)
	testAttrP(t, "ConfigureCStdio", config.ConfigureCStdio, config.CInstance.ConfigureCStdio)
	testAttrP(t, "DevMode", config.DevMode, config.CInstance.DevMode)
	testAttrP(t, "DumpRefs", config.DumpRefs, config.CInstance.DumpRefs)
	testAttrP(t, "ExecPrefix", config.ExecPrefix, config.CInstance.ExecPrefix)
	testAttrP(t, "Executable", config.Executable, config.CInstance.Executable)
	testAttrP(t, "FaultHandler", config.FaultHandler, config.CInstance.FaultHandler)
	testAttrP(t, "FilesystemEncoding", config.FilesystemEncoding, config.CInstance.FilesystemEncoding)
	testAttrP(t, "FilesystemErrors", config.FilesystemErrors, config.CInstance.FilesystemErrors)
	testAttrP(t, "Home", config.Home, config.CInstance.Home)
	testAttrP(t, "ImportTime", config.ImportTime, config.CInstance.ImportTime)
	testAttrP(t, "Inspect", config.Inspect, config.CInstance.Inspect)
	testAttrP(t, "InstallSignalHandlers", config.InstallSignalHandlers, config.CInstance.InstallSignalHandlers)
	testAttrP(t, "Interactive", config.Interactive, config.CInstance.Interactive)
	testAttrP(t, "Isolated", config.Isolated, config.CInstance.Isolated)
	testAttrP(t, "MallocStats", config.MallocStats, config.CInstance.MallocStats)
	testAttrP(t, "PythonPathEnv", config.PythonPathEnv, config.CInstance.PythonPathEnv)
	testAttrP(t, "ModuleSearchPaths", config.ModuleSearchPaths, config.CInstance.ModuleSearchPaths)
	testAttrP(t, "ModuleSearchPathsSet", config.ModuleSearchPathsSet, config.CInstance.ModuleSearchPathsSet)
	testAttrP(t, "OptimizationLevel", config.OptimizationLevel, config.CInstance.OptimizationLevel)
	testAttrP(t, "ParseArgv", config.ParseArgv, config.CInstance.ParseArgv)
	testAttrP(t, "ParserDebug", config.ParserDebug, config.CInstance.ParserDebug)
	testAttrP(t, "PathConfigWarnings", config.PathConfigWarnings, config.CInstance.PathConfigWarnings)
	testAttrP(t, "Prefix", config.Prefix, config.CInstance.Prefix)
	testAttrP(t, "ProgramName", config.ProgramName, config.CInstance.ProgramName)
	testAttrP(t, "PyCachePrefix", config.PyCachePrefix, config.CInstance.PyCachePrefix)
	testAttrP(t, "Quiet", config.Quiet, config.CInstance.Quiet)
	testAttrP(t, "RunCommand", config.RunCommand, config.CInstance.RunCommand)
	testAttrP(t, "RunFileName", config.RunFileName, config.CInstance.RunFileName)
	testAttrP(t, "RunModule", config.RunModule, config.CInstance.RunModule)
	testAttrP(t, "ShowRefCount", config.ShowRefCount, config.CInstance.ShowRefCount)
	testAttrP(t, "SiteImport", config.SiteImport, config.CInstance.SiteImport)
	testAttrP(t, "SkipSourceFirstLine", config.SkipSourceFirstLine, config.CInstance.SkipSourceFirstLine)
	testAttrP(t, "StdioEncoding", config.StdioEncoding, config.CInstance.StdioEncoding)
	testAttrP(t, "StdioErrors", config.StdioErrors, config.CInstance.StdioErrors)
	testAttrP(t, "TraceMalloc", config.TraceMalloc, config.CInstance.TraceMalloc)
	testAttrP(t, "UseEnvironment", config.UseEnvironment, config.CInstance.UseEnvironment)
	testAttrP(t, "UseHashSeed", config.UseHashSeed, config.CInstance.UseHashSeed)
	testAttrP(t, "UserSiteDirectory", config.UserSiteDirectory, config.CInstance.UserSiteDirectory)
	testAttrP(t, "Verbose", config.Verbose, config.CInstance.Verbose)
	testAttrP(t, "WarnOptions", config.WarnOptions, config.CInstance.WarnOptions)
	testAttrP(t, "WriteByteCode", config.WriteByteCode, config.CInstance.WriteByteCode)
	testAttrP(t, "XOptions", config.XOptions, config.CInstance.XOptions)
}

func TestConfigSetInteger(t *testing.T) {
	config := pyg.ConfigNew()
	defer config.Clear()
	config.InitIsolatedConfig()

	expected, got := 3, 0
	config.SetInteger(config.Verbose, expected)
	got = int(*config.CInstance.Verbose)
	if !(expected == got) {
		t.Errorf(
			"expected %d got %d. Changes from top level should reflect in lower",
			expected,
			got)
	}
}

func TestConfigSetBytesString(t *testing.T) {
	config := pyg.ConfigNew()
	defer config.Clear()
	config.InitIsolatedConfig()

	expected, got := "PyGi", ""
	config.SetBytesString(config.ProgramName, expected)
	got = pyg.WString2String(*config.ProgramName)
	testEqualS(t, expected, got, "Changes from top level should reflect in lower")

}

func TestConfigSetString(t *testing.T) {
	config := pyg.ConfigNew()
	defer config.Clear()
	config.InitIsolatedConfig()

	expected, got := "PyGi_2", ""
	config.SetString(config.ProgramName, pyg.String2WideString(expected))
	got = pyg.WString2String(*config.ProgramName)
	testEqualS(t, expected, got, "Changes from top level should reflect in lower")
}

func TestStringTxl2WideString(t *testing.T) {
	original := "Patty cake, patty cake"
	// Transform a string into a wide string, and
	// back again.
	got := pyg.WString2String(pyg.String2WideString(original))
	testEqualS(t, original, got, "String translation should be 1:1")
}
