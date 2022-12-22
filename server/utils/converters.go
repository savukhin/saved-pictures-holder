package utils

import (
	"encoding/json"
)

func ConvertToMap(data interface{}) (map[string]interface{}, error) {
	// return data.(map[string]interface{})
	// result := make(map[string]interface{})
	// fmt.Println("data=", data)
	// fmt.Println("typed data=", data.(map[string]interface{}))
	// for k, v := range data.(map[string]interface{}) {
	// 	result[k] = v
	// }
	// return result

	// return structs.Map(data)

	jsoned, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(jsoned, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
