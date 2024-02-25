package handler_test

import (
	"os"
	"testing"

	"github.com/fchristian1/helper_GO"
)

func TestFile(t *testing.T) {
	//arrange
	jsonString := `{"test": "data"}`

	tempFile, err := os.CreateTemp("", "example.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.WriteString(jsonString)
	byteTemp := []byte(jsonString)
	//act
	byteTest, _ := helper_GO.FILE_LoadFileAsByte(tempFile.Name())
	_, err = helper_GO.FILE_LoadFileAsByte("")
	//assert
	if string(byteTemp) != string(byteTest) {
		t.Errorf("Expected %s, got %s", string(byteTemp), string(byteTest))
	}
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
