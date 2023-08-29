package utils

import (
	"fmt"
	"reflect"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func EnsureMap(value interface{}) map[string]interface{} {
	if reflect.TypeOf(value).Kind() == reflect.Map {
		result, ok := value.(map[string]interface{})

		if !ok {
			panic("Failed converting to map")
		}

		return result
	}

	panic("Value is not a map")
}

func EnsureKeys(value map[string]interface{}, keys []string) {
	for _, key := range keys {
		_, ok := value[key]

		if !ok {
			fmt.Println(fmt.Sprintf("Missing required key: %s", key))
			fmt.Println("Required keys are:")

			for _, k := range keys {
				fmt.Println(k)
			}

			panic("Exiting due missing key")
		}
	}
}
