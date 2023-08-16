package common

import (
	"encoding/json"
	"fmt"
)

func JsonPretty(v any) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("fail json")
		return
	}
	fmt.Printf("%s\n", data)
}
