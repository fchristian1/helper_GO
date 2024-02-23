package helper_GO

import "github.com/google/uuid"

func UUID_NewAsString() string {
	return uuid.New().String()
}

func UUID_NameNewString(name string) string {
	return name + "_" + uuid.New().String()
}
