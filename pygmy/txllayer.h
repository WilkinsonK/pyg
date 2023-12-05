#ifndef _PYGMY_TXLLAYER_H
#define _PYGMY_TXLLAYER_H

#define PY_SSIZE_T_CLEAN
#include <wchar.h>
#include <Python.h>

#define PY_EMBED_VERSION 311
#define PY_EMBED_COMPAT(VERS) defined(PY_EMBED_VERSION) && PY_EMBED_VERSION >= VERS

typedef struct CGO_PyConfig {
    PyWideStringList *Argv;
    wchar_t*         *BaseExecPrefix;
    wchar_t*         *BaseExecutable;
    wchar_t*         *BasePrefix;
    int              *BufferedStdio;
    int              *BytesWarning;
    wchar_t*         *CheckHashPYCsMode;
    int              *ConfigureCStdio;
    int              *DevMode;
    int              *DumpRefs;
    wchar_t*         *ExecPrefix;
    wchar_t*         *Executable;
    int              *FaultHandler;
    wchar_t*         *FilesystemEncoding;
    wchar_t*         *FilesystemErrors;
    wchar_t*         *Home;
    int              *ImportTime;
    int              *Inspect;
    int              *InstallSignalHandlers;
    int              *Interactive;
    int              *Isolated;
    int              *MallocStats;
    wchar_t*         *PythonPathEnv;
    PyWideStringList *ModuleSearchPaths;
    int              *ModuleSearchPathsSet;
    int              *OptimizationLevel;
    int              *ParseArgv;
    int              *ParserDebug;
    int              *PathConfigWarnings;
    wchar_t*         *Prefix;
    wchar_t*         *ProgramName;
    wchar_t*         *PyCachePrefix;
    int              *Quiet;
    wchar_t*         *RunCommand;
    wchar_t*         *RunFileName;
    wchar_t*         *RunModule;
    int              *ShowRefCount;
    int              *SiteImport;
    int              *SkipSourceFirstLine;
    wchar_t*         *StdioEncoding;
    wchar_t*         *StdioErrors;
    int              *TraceMalloc;
    int              *UseEnvironment;
    int              *UseHashSeed;
    int              *UserSiteDirectory;
    int              *Verbose;
    PyWideStringList *WarnOptions;
    int              *WriteByteCode;
    PyWideStringList *XOptions;

    #if PY_EMBED_COMPAT(309)
    wchar_t*         *PlatLibDir;
    #endif
    #if PY_EMBED_COMPAT(310)
    PyWideStringList OrigArgv;
    int              *WarnDefaultEncoding;
    #endif
    #if PY_EMBED_COMPAT(311)
    int              *CodeDebugRanges;
    int              *SafePath;
    #endif
    #if PY_EMBED_COMPAT(312)
    int              *IntMaxStrDigits;
    int              *PerfProfiling;
    #endif
    #ifdef MS_WINDOWS
    int *LegacyWindowsStdio;
    #endif

    // Original `PyConfig` object.
    PyConfig* cInstance;
    int       cInstanceMapped;
} CGO_PyConfig;

typedef struct CGO_PyPreConfig {
    int *Allocator;
    int *ConfigureLocale;
    int *CoerceCLocale;
    int *DevMode;
    int *Isolated;
    int *ParseArgv;
    int *UseEnvironment;
    int *UTF8Mode;

    #ifdef MS_WINDOWS
    int *LegacyWindowsFSEncoding;
    #endif

    // Original `PyConfig` object.
    PyPreConfig* cInstance;
    int          cInstanceMapped;
} CGO_PyPreConfig;

void CGO_PyConfig_InitFieldRefs(CGO_PyConfig*);
void CGO_PyConfig_Clear(CGO_PyConfig*);
void CGO_PyConfig_InitPythonConfig(CGO_PyConfig*);
void CGO_PyConfig_InitIsolatedConfig(CGO_PyConfig*);
PyStatus CGO_PyConfig_Read(CGO_PyConfig*);
PyStatus CGO_PyConfig_SetArgv(CGO_PyConfig*, Py_ssize_clean_t, wchar_t* const*);
PyStatus CGO_PyConfig_SetBytesArgv(CGO_PyConfig*, Py_ssize_clean_t, char* const*);
PyStatus CGO_PyConfig_SetBytesString(CGO_PyConfig*, wchar_t**, const char*);
PyStatus CGO_PyConfig_SetString(CGO_PyConfig*, wchar_t**, const wchar_t*);
PyStatus CGO_PyConfig_SetWideStringList(CGO_PyConfig*, PyWideStringList*, Py_ssize_clean_t, wchar_t**);
void CGO_PyPreConfig_Clear(CGO_PyPreConfig*);
void CGO_PyPreConfig_InitFieldRefs(CGO_PyPreConfig*);
void CGO_PyPreConfig_InitPythonConfig(CGO_PyPreConfig*);
void CGO_PyPreConfig_InitIsolatedConfig(CGO_PyPreConfig*);
PyStatus CGO_PreInitialize(const CGO_PyPreConfig*);
PyStatus CGO_Py_InitializeFromConfig(const CGO_PyConfig*);
PyStatus CGO_Py_PreInitializeFromArgs(CGO_PyPreConfig*, Py_ssize_clean_t, wchar_t**);
PyStatus CGO_Py_PreInitializeFromBytesArgs(CGO_PyPreConfig*, Py_ssize_clean_t, char**);
#endif // _PYGMY_TXLLAYER_H
