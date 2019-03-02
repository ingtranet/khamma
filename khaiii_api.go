package khamma

/*
#cgo LDFLAGS: -ldl
#include <stdlib.h>
#include "bindings.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

func LoadLibrary(libPath string) error {
	cLibPath := C.CString(libPath)
	defer C.free(unsafe.Pointer(cLibPath))

	result := C.init_funct_ptrs(cLibPath)

	if result != nil {
		//defer C.free(unsafe.Pointer(result))
		return errors.New(C.GoString(result))
	}

	return nil
}

func KhaiiiVersion() string {
	strVer := C.khaiii_version()
	return C.GoString(strVer)
}

func KhaiiiOpen(rscDir string, optDir string) int {
	cRscDir := C.CString(rscDir)
	defer C.free(unsafe.Pointer(cRscDir))
	cOptDir := C.CString(optDir)
	defer C.free(unsafe.Pointer(cOptDir))

	return int(C.khaiii_open(cRscDir, cOptDir))
}

func KhaiiiAnalyze(handle int, input string, optStr string) *C.khaiii_word_t {
	cInput := C.CString(input)
	defer C.free(unsafe.Pointer(cInput))
	cOptStr := C.CString(optStr)
	defer C.free(unsafe.Pointer(cOptStr))

	return C.khaiii_analyze(C.int(handle), cInput, cOptStr)
}

func KhaiiiFreeResults(handle int, results *C.khaiii_word_t) {
	C.khaiii_free_results(C.int(handle), results)
}

func KhaiiiClose(handle int) {
	C.khaiii_close(C.int(handle))
}

func KhaiiiLastError(handle int) string {
	return C.GoString(C.khaiii_last_error(C.int(handle)))
}
