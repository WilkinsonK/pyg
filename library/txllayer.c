#include "txllayer.h"

void CGO_PyConfig_InitFieldRefs(CGO_PyConfig* config) {
    // Bail if reference map has already been
    // completed.
    if (config->cInstanceMapped != 0) return;
    config->cInstance = (PyConfig*)malloc(sizeof(PyConfig));

    config->Argv                  = &config->cInstance->argv;
    config->BaseExecPrefix        = &config->cInstance->base_exec_prefix;
    config->BaseExecutable        = &config->cInstance->base_executable;
    config->BasePrefix            = &config->cInstance->base_prefix;
    config->BufferedStdio         = &config->cInstance->buffered_stdio;
    config->BytesWarning          = &config->cInstance->bytes_warning;
    config->CheckHashPYCsMode     = &config->cInstance->check_hash_pycs_mode;
    config->ConfigureCStdio       = &config->cInstance->configure_c_stdio;
    config->DevMode               = &config->cInstance->dev_mode;
    config->DumpRefs              = &config->cInstance->dump_refs;
    config->ExecPrefix            = &config->cInstance->exec_prefix;
    config->Executable            = &config->cInstance->executable;
    config->FaultHandler          = &config->cInstance->faulthandler;
    config->FilesystemEncoding    = &config->cInstance->filesystem_encoding;
    config->FilesystemErrors      = &config->cInstance->filesystem_errors;
    config->Home                  = &config->cInstance->home;
    config->ImportTime            = &config->cInstance->import_time;
    config->Inspect               = &config->cInstance->inspect;
    config->InstallSignalHandlers = &config->cInstance->install_signal_handlers;
    config->Interactive           = &config->cInstance->interactive;
    config->Isolated              = &config->cInstance->isolated;
    config->MallocStats           = &config->cInstance->malloc_stats;
    config->ModuleSearchPaths     = &config->cInstance->module_search_paths;
    config->ModuleSearchPathsSet  = &config->cInstance->module_search_paths_set;
    config->OptimizationLevel     = &config->cInstance->optimization_level;
    config->ParseArgv             = &config->cInstance->parse_argv;
    config->ParserDebug           = &config->cInstance->parser_debug;
    config->PathConfigWarnings    = &config->cInstance->pathconfig_warnings;
    config->Prefix                = &config->cInstance->prefix;
    config->ProgramName           = &config->cInstance->program_name;
    config->PyCachePrefix         = &config->cInstance->pycache_prefix;
    config->PythonPathEnv         = &config->cInstance->pythonpath_env;
    config->Quiet                 = &config->cInstance->quiet;
    config->RunCommand            = &config->cInstance->run_command;
    config->RunFileName           = &config->cInstance->run_filename;
    config->RunModule             = &config->cInstance->run_module;
    config->ShowRefCount          = &config->cInstance->show_ref_count;
    config->SiteImport            = &config->cInstance->site_import;
    config->SkipSourceFirstLine   = &config->cInstance->skip_source_first_line;
    config->StdioEncoding         = &config->cInstance->stdio_encoding;
    config->StdioErrors           = &config->cInstance->stdio_errors;
    config->TraceMalloc           = &config->cInstance->tracemalloc;
    config->UseEnvironment        = &config->cInstance->use_environment;
    config->UseHashSeed           = &config->cInstance->use_hash_seed;
    config->UserSiteDirectory     = &config->cInstance->user_site_directory;
    config->Verbose               = &config->cInstance->verbose;
    config->WarnOptions           = &config->cInstance->warnoptions;
    config->WriteByteCode         = &config->cInstance->write_bytecode;
    config->XOptions              = &config->cInstance->xoptions;

    #if PY_EMBED_COMPAT(309)
    config->PlatLibDir            = &config->cInstance->platlibdir;
    #endif
    #if PY_EMBED_COMPAT(310)
    config->OrigArgv              = config->cInstance->orig_argv;
    config->WarnDefaultEncoding   = &config->cInstance->warn_default_encoding;
    #endif
    #if PY_EMBED_COMPAT(311)
    config->CodeDebugRanges       = &config->cInstance->code_debug_ranges;
    config->SafePath              = &config->cInstance->safe_path;
    #endif
    #if PY_EMBED_COMPAT(312)
    config->IntMaxStrDigits       = &config->cInstance->int_max_str_digits;
    config->PerfProfiling         = &config->cInstance->perf_profiling;
    #endif
    #ifdef MS_WINDOWS
    config->LegacyWindowsStdio    = &config->cInstance->legacy_windows_stdio;
    #endif

    config->cInstanceMapped = 1;
}

void CGO_PyConfig_Clear(CGO_PyConfig* config) {
    if (config->cInstanceMapped == 0) return;

    PyConfig_Clear(config->cInstance);
    free(config->cInstance);
    config->cInstanceMapped = 0;
}

void CGO_PyConfig_InitPythonConfig(CGO_PyConfig* config) {
    CGO_PyConfig_InitFieldRefs(config);
    PyConfig_InitPythonConfig(config->cInstance);
}

void CGO_PyConfig_InitIsolatedConfig(CGO_PyConfig* config) {
    CGO_PyConfig_InitFieldRefs(config);
    PyConfig_InitIsolatedConfig(config->cInstance);
}

PyStatus CGO_Py_InitializeFromConfig(const CGO_PyConfig* config) {
    return Py_InitializeFromConfig(config->cInstance);
}

PyStatus CGO_PyConfig_Read(CGO_PyConfig* config) {
    return PyConfig_Read(config->cInstance);
}

PyStatus CGO_PyConfig_SetArgv(
    CGO_PyConfig* config,
    Py_ssize_clean_t argc,
    wchar_t *const *argv) {

    return PyConfig_SetArgv(config->cInstance, argc, argv);
}

PyStatus CGO_PyConfig_SetBytesArgv(
    CGO_PyConfig* config,
    Py_ssize_clean_t argc,
    char *const *argv) {

    return PyConfig_SetBytesArgv(config->cInstance, argc, argv);
}

PyStatus CGO_PyConfig_SetBytesString(
    CGO_PyConfig* config,
    wchar_t** config_str,
    const char* str) {

    return PyConfig_SetBytesString(config->cInstance, config_str, str);
}

PyStatus CGO_PyConfig_SetString(
    CGO_PyConfig* config,
    wchar_t** config_str,
    const wchar_t* str) {

    return PyConfig_SetString(config->cInstance, config_str, str);
}

PyStatus CGO_PyConfig_SetWideStringList(
    CGO_PyConfig* config,
    PyWideStringList* list,
    Py_ssize_clean_t length,
    wchar_t** items) {

    return PyConfig_SetWideStringList(
        config->cInstance,
        list,
        length,
        items);
}

void CGO_PyPreConfig_Clear(CGO_PyPreConfig* config) {
    if (config->cInstanceMapped == 0) return;

    free(config->cInstance);
    config->cInstanceMapped = 0;
}

void CGO_PyPreConfig_InitFieldRefs(CGO_PyPreConfig* config) {
    // Bail if reference map has already been
    // completed.
    if (config->cInstanceMapped != 0) return;

    config->cInstance = (PyPreConfig*)malloc(sizeof(PyPreConfig));
    config->Allocator       = &config->cInstance->allocator;
    config->ConfigureLocale = &config->cInstance->configure_locale;
    config->CoerceCLocale   = &config->cInstance->coerce_c_locale;
    config->DevMode         = &config->cInstance->dev_mode;
    config->Isolated        = &config->cInstance->isolated;
    config->ParseArgv       = &config->cInstance->parse_argv;
    config->UseEnvironment  = &config->cInstance->use_environment;
    config->UTF8Mode        = &config->cInstance->utf8_mode;

    #ifdef MS_WINDOWS
    config->LegacyWindowsFSEncoding = &config->cInstance->legacy_windows_fs_encoding;
    #endif

    config->cInstanceMapped = 1;
}

void CGO_PyPreConfig_InitPythonConfig(CGO_PyPreConfig* config) {
    CGO_PyPreConfig_InitFieldRefs(config);
    PyPreConfig_InitPythonConfig(config->cInstance);
}

void CGO_PyPreConfig_InitIsolatedConfig(CGO_PyPreConfig* config) {
    CGO_PyPreConfig_InitFieldRefs(config);
    PyPreConfig_InitIsolatedConfig(config->cInstance);
}

PyStatus CGO_PreInitialize(const CGO_PyPreConfig* config) {
    return Py_PreInitialize(config->cInstance);
}

PyStatus CGO_Py_PreInitializeFromArgs(
    CGO_PyPreConfig* config,
    Py_ssize_clean_t argc,
    wchar_t** argv) {

    CGO_PyPreConfig_InitFieldRefs(config);
    return Py_PreInitializeFromArgs(config->cInstance, argc, argv);
}

PyStatus CGO_Py_PreInitializeFromBytesArgs(
    CGO_PyPreConfig* config,
    Py_ssize_clean_t argc,
    char **argv) {

    CGO_PyPreConfig_InitFieldRefs(config);
    return Py_PreInitializeFromBytesArgs(config->cInstance, argc, argv);
};
