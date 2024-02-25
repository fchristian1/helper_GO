package helper_GO

import "os"

// FILE_LoadFileAsByte loads a file as a byte array.
//
// It takes a path string as a parameter and returns a byte array.
func FILE_LoadFileAsByte(path string) ([]byte, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// FILE_SaveFileAsByte saves a byte array as a file.
//
// It takes a path string and a byte array as parameters and returns an error.
func FILE_SaveFileAsByte(path string, bytes []byte) error {
	err := os.WriteFile(path, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

// FILE_Exists checks if a file exists.
//
// It takes a path string as a parameter and returns a boolean.
func FILE_Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// FILE_DeleteFile deletes a file.
//
// It takes a path string as a parameter and returns an error.
func FILE_DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

// FILE_DeleteFolder deletes a folder.
//
// It takes a path string as a parameter and returns an error.
func FILE_DeleteFolder(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}

// FILE_CreateFolder creates a folder.
//
// It takes a path string as a parameter and returns an error.
func FILE_CreateFolder(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	return nil
}

// FILE_CreateFile creates a file.
//
// It takes a path string as a parameter and returns an error.
func FILE_CreateFile(path string) error {
	_, err := os.Create(path)
	if err != nil {
		return err
	}
	return nil
}
