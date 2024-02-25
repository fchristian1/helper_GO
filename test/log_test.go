package handler_test

import (
	"testing"

	"github.com/fchristian1/helper_GO"
)

func TestLog(t *testing.T) {
	//arrange
	logLevel := "info"
	message := "test"
	//act
	i := helper_GO.LogInfo(logLevel, message)
	w := helper_GO.LogWarning(logLevel, message)
	e := helper_GO.LogError(logLevel, message)
	d := helper_GO.LogDebug(logLevel, message)
	//assert
	if !i {
		t.Errorf("Expected true, got %t", i)
	}
	if w {
		t.Errorf("Expected false, got %t", w)
	}
	if e {
		t.Errorf("Expected false, got %t", e)
	}
	if d {
		t.Errorf("Expected false, got %t", d)
	}

}

func TestLogPrint(t *testing.T) {
	//arrange
	name := "info"
	message := "test"
	//act
	m := helper_GO.LogPrint(name, message)
	//assert
	if m == "" {
		t.Errorf("Expected not empty, got %s", m)
	}
}

func TestSaveLog(t *testing.T) {
	//arrange
	filename := "test.txt"
	message := "test"
	//act
	helper_GO.SaveLog(filename, message)
	//assert
	if !helper_GO.FILE_Exists(filename) {
		t.Errorf("Expected true, got %t", helper_GO.FILE_Exists(filename))
	}
}
