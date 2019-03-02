#include <dlfcn.h>
#include <stdlib.h>
#include "bindings.h"

void* handle;
char* (*ptr_khaiii_version)();
int (*ptr_khaiii_open)(const char*, const char*);
const khaiii_word_t* (*ptr_khaiii_analyze)(int, const char*, const char*);
void (*ptr_khaiii_free_results)(int, const khaiii_word_t*);
void (*ptr_khaiii_close)(int);
const char* (*ptr_khaiii_last_error)(int);

char* init_funct_ptrs(char* libpath) {
	handle = dlopen(libpath, RTLD_LAZY);
    if (handle == NULL) return dlerror();
    
	ptr_khaiii_version = dlsym(handle, "khaiii_version");
    if (ptr_khaiii_version == NULL) return dlerror();

    ptr_khaiii_open = dlsym(handle, "khaiii_open");
    if (ptr_khaiii_open == NULL) return dlerror();

	ptr_khaiii_analyze = dlsym(handle, "khaiii_analyze");
    if (ptr_khaiii_analyze == NULL) return dlerror();

    ptr_khaiii_free_results = dlsym(handle, "khaiii_free_results");
    if (ptr_khaiii_free_results == NULL) return dlerror();
    
    ptr_khaiii_close = dlsym(handle, "khaiii_close");
    if (ptr_khaiii_close == NULL) return dlerror();

    ptr_khaiii_last_error = dlsym(handle, "khaiii_last_error");
    if (ptr_khaiii_last_error == NULL) return dlerror();

    return NULL;
}

char* khaiii_version() {
	return ptr_khaiii_version();
}

int khaiii_open(const char* rsc_dir, const char* opt_str) {
	return ptr_khaiii_open(rsc_dir, opt_str);
}

const khaiii_word_t* khaiii_analyze(int handle, const char* input, const char* opt_str) {
	return ptr_khaiii_analyze(handle, input, opt_str);
}

void khaiii_free_results(int handle, const khaiii_word_t* results) {
    ptr_khaiii_free_results(handle, results);
}

void khaiii_close(int handle) {
    ptr_khaiii_close(handle);
}

const char* khaiii_last_error(int handle) {
    return ptr_khaiii_last_error(handle);
}
