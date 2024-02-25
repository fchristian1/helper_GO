package handler_test

import (
	"encoding/json"
	"testing"
)

func TestJSON(t *testing.T) {
	//arrange
	jsonString := `{"test": "data"}`
	type TestStruct struct {
		Test string `json:"test"`
	}
	var testStruct TestStruct
	//act
	err := json.Unmarshal([]byte(jsonString), &testStruct)
	//assert
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}
	if testStruct.Test != "data" {
		t.Errorf("Expected 'data', got %s", testStruct.Test)
	}
}
