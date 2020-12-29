package parser

import (
	"encoding/json"
	"strconv"

	"github.com/genkami/watson"
)

func encodeAsInt(raw []byte) int64 {
	i, err := strconv.ParseInt(string(raw), 10, 64)

	if err != nil {
		return 0
	}

	return i
}

func encodeAsString(raw []byte) string {
	return string(raw)
}

func encodeAsFloat(raw []byte) float64 {
	i, err := strconv.ParseFloat(string(raw), 64)

	if err != nil {
		return 0.0
	}

	return i
}

func encodeAsBool(raw []byte) bool {
	b, err := strconv.ParseBool(string(raw))

	if err != nil {
		return false
	}

	return b
}

func encodeAsObject(raw []byte) map[string]interface{} {
	m := make(map[string]interface{})
	json.Unmarshal(raw, &m)
	return m
}

func encodeAsArray(raw []byte) []interface{} {
	var arr []interface{}
	json.Unmarshal(raw, &arr)
	return arr
}

// EncodeAs will create a WATSON-encoded byte array based on the provided data
// type.
func EncodeAs(raw []byte, dataType DataType) ([]byte, error) {
	switch dataType {
	case DTInt:
		input := encodeAsInt(raw)
		return watson.Marshal(&input)
	case DTString:
		input := encodeAsString(raw)
		return watson.Marshal(&input)
	case DTFloat:
		input := encodeAsFloat(raw)
		return watson.Marshal(&input)
	case DTBool:
		input := encodeAsBool(raw)
		return watson.Marshal(&input)
	case DTObject:
		input := encodeAsObject(raw)
		return watson.Marshal(&input)
	case DTArray:
		input := encodeAsArray(raw)
		return watson.Marshal(&input)
	}

	return watson.Marshal(&raw)
}
