package utils

import (
	"github.com/fatih/structs"
)

func ConvertToMap(data interface{}) map[string]interface{} {
	// return data.(map[string]interface{})
	// result := make(map[string]interface{})
	// fmt.Println("data=", data)
	// fmt.Println("typed data=", data.(map[string]interface{}))
	// for k, v := range data.(map[string]interface{}) {
	// 	result[k] = v
	// }
	// return result

	return structs.Map(data)
}
