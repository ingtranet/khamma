package khamma

import (
	"errors"
	"runtime"
)

var handle int

func Initialze(libPath string, rscDir string, optDir string) error {
	err := LoadLibrary(libPath)
	if err != nil {
		return err
	}

	handle = KhaiiiOpen(rscDir, optDir)
	if handle == 0 {
		return errors.New("Opening resource failed")
	}

	return nil
}

func InitializeWithDefault() error {
	var libPath string

	switch osName := runtime.GOOS; osName {
	case "darwin":
		libPath = "libkhaiii.dylib"
	case "linux":
		libPath = "libkhaiii.so"
	default:
		return errors.New("OS not supported: " + osName)
	}
	return Initialze(libPath, "", "")
}

func Analyze(text string, optStr string) []*KhaiiiWord {
	result := make([]*KhaiiiWord, 0)
	firstWord := KhaiiiAnalyze(handle, text, optStr)
	for word := firstWord; word != nil; word = word.next {
		result = append(result, word.ToGoStruct())
	}
	KhaiiiFreeResults(handle, firstWord)
	return result
}
