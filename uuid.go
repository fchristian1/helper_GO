package helper_GO

import (
	"github.com/google/uuid"
)

func UUID_NewString() string {
	return uuid.New().String()
}

func UUID_NameNewString(name string) string {
	return name + "_" + uuid.New().String()
}

// UUID validation
func UUID_Validate(uuidString string) bool {
	err := uuid.Validate(uuidString)
	return err == nil
}
