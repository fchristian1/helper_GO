package helper_GO

import "encoding/json"

func JSON_UnmarshalToInterface[T any](b []byte) T {
	var i T
	err := json.Unmarshal(b, &i)
	if err != nil {
		panic(err)
	}
	return i
}
