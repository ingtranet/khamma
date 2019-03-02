package khamma

/*
#cgo LDFLAGS: -ldl
#include <stdlib.h>
#include "bindings.h"
*/
import "C"

// Go struct for khaiii_morph_t
type KhaiiiMorph struct {
	Lex      string `json:"lex"`
	Tag      string `json:"tag"`
	Begin    int    `json:"begin"`
	Length   int    `json:"length"`
	Reserved string `json:"-"`
}

// Go struct for khaiii_word_t
type KhaiiiWord struct {
	Begin    int            `json:"begin"`
	Length   int            `json:"length"`
	Reserved string         `json:"-"`
	Morphs   []*KhaiiiMorph `json:"morphs"`
}

// Go converters for khaiii_morph_t
func (morphType *C.khaiii_morph_t) Lex() string {
	return C.GoString(morphType.lex)
}

func (morphType *C.khaiii_morph_t) Tag() string {
	return C.GoString(morphType.tag)
}

func (morphType *C.khaiii_morph_t) Begin() int {
	return int(morphType.begin)
}

func (morphType *C.khaiii_morph_t) Length() int {
	return int(morphType.length)
}

func (morphType *C.khaiii_morph_t) Reserved() string {
	return C.GoString(&morphType.reserved[0])
}

func (morphType *C.khaiii_morph_t) Next() *C.khaiii_morph_t {
	return morphType.next
}

func (morphType *C.khaiii_morph_t) ToGoStruct() *KhaiiiMorph {
	return &KhaiiiMorph{
		Lex:      morphType.Lex(),
		Tag:      morphType.Tag(),
		Begin:    morphType.Begin(),
		Length:   morphType.Length(),
		Reserved: morphType.Reserved()}
}

// Go converters for khaiii_word_t
func (wordType *C.khaiii_word_t) Begin() int {
	return int(wordType.begin)
}

func (wordType *C.khaiii_word_t) Length() int {
	return int(wordType.length)
}

func (wordType *C.khaiii_word_t) Reserved() string {
	return C.GoString(&wordType.reserved[0])
}

func (wordType *C.khaiii_word_t) Morphs() *C.khaiii_morph_t {
	return wordType.morphs
}

func (wordType *C.khaiii_word_t) Next() *C.khaiii_word_t {
	return wordType.next
}

func (wordType *C.khaiii_word_t) ToGoStruct() *KhaiiiWord {
	morphs := make([]*KhaiiiMorph, 0)
	for morph := wordType.morphs; morph != nil; morph = morph.next {
		morphs = append(morphs, morph.ToGoStruct())
	}
	return &KhaiiiWord{
		Begin:    wordType.Begin(),
		Length:   wordType.Length(),
		Reserved: wordType.Reserved(),
		Morphs:   morphs}
}
