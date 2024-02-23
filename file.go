package helper_GO

import "os"

func FILE_LoadFileAsByte(path string) []byte {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bytes
}
