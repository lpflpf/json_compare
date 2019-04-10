package json_compare

import (
	"encoding/json"
	"reflect"
)

func byte2Interface(input []byte) (output interface{}) {
	_ = json.Unmarshal(input, &output)
	return output
}

func Compare(data1 []byte, data2 []byte) bool {
	d1 := byte2Interface(data1)
	d2 := byte2Interface(data2)

	return interfaceCompare(d1, d2)
}

func interfaceCompare(data1, data2 interface{}) bool {
	return reflect.DeepEqual(data1, data2)
}
