package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type State struct {
	Name     string
	Senators []string
	Water    float64
	Area     int
}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int8, int16, int32, int64:
			fmt.Printf("param #%d is an int\n", i)
		case uint, uint8, uint16, uint32, uint64:
			fmt.Printf("param #%d is an unsigned int\n", i)
		case nil:
			fmt.Printf("param #%d is nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d's type is unknow\n", i)
		}
	}
}

func jsonObjectAsString(jsonObject map[string]interface{}) string {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	comma := ""
	for key, value := range jsonObject {
		buffer.WriteString(comma)
		switch value := value.(type) {
		case nil:
			fmt.Fprintf(&buffer, "%q: null", key)
		case bool:
			fmt.Fprintf(&buffer, "%q: %t", key, value)
		case float64:
			fmt.Fprintf(&buffer, "%q: %f", key, value)
		case string:
			fmt.Fprintf(&buffer, "%q: %q", key, value)
		case []interface{}:
			fmt.Fprintf(&buffer, "%q: [", key)
			innerComma := ""
			for _, s := range value {
				if s, ok := s.(string); ok {
					fmt.Fprintf(&buffer, "%s%q", innerComma, s)
					innerComma = ", "
				}
			}
			buffer.WriteString("]")
		}
		comma = ", "
	}
	buffer.WriteString("}")
	return buffer.String()
}

func unmar() {
	MA := []byte("{\"name\": \"Massachusetts\", \"area\": 27336, \"water\": 25.7, \"senators\": [\"John Kerry\", \"Scott Brown\"]}")
	var object interface{}
	if err := json.Unmarshal(MA, &object); err != nil {
		fmt.Println(err)
	} else {
		jsonObject := object.(map[string]interface{})
		fmt.Println(jsonObjectAsString(jsonObject))
	}
}

func unmar2() {
	var state State
	MA := []byte("{\"name\": \"Massachusetts\", \"area\": 27336, \"water\": 25.7, \"senators\": [\"John Kerry\", \"Scott Brown\"]}")
	if err := json.Unmarshal(MA, &state); err != nil {
		fmt.Println(err)
	}
	fmt.Println(state)
}

func main() {
	classifier(5, -17.9, "ZIP", nil, true, complex(1, 1))
	unmar2()
}
