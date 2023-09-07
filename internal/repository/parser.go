package repository

import (
	"bytes"
	"fmt"
	"strings"
)

type OrderedDict struct {
	keys  []string
	value map[string]interface{}
}

func NewOrderedMap() *OrderedDict {
	return &OrderedDict{
		keys:  make([]string, 0),
		value: make(map[string]interface{}),
	}
}

func (om *OrderedDict) Add(key string, value interface{}) {
	om.keys = append(om.keys, key)
	om.value[key] = value
}

func KvlmParse(raw []byte, start int, dct *OrderedDict) (*OrderedDict, int) {
	if dct == nil {
		dct = NewOrderedMap()
	}

	spxIdx := bytes.IndexByte(raw[start:], ' ')
	nwlIdx := bytes.IndexByte(raw[start:], '\n')

	if spxIdx < 0 || nwlIdx < spxIdx {
		return kvlmParseMessage(raw, start, dct)
	}

	key := string(raw[start : start+spxIdx])
	end := kvlmFindValueEnd(raw, start+spxIdx)

	value := string(raw[start+spxIdx+1 : end])
	value = strings.Replace(value, "\n ", "\n", 1)
	dct.Add(key, value)

	return KvlmParse(raw, end+1, dct)
}

func kvlmFindValueEnd(raw []byte, start int) int {

	for i := start; i < len(raw)-1; i++ {
		if raw[i] == '\n' && raw[i+1] != ' ' {
			return i
		}
	}

	return len(raw)
}

func kvlmParseMessage(raw []byte, start int, od *OrderedDict) (*OrderedDict, int) {
	od.Add("", string(raw[start+1:]))
	return od, len(raw)
}
