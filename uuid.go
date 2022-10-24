package strgen

import (
	"strings"

	"github.com/google/uuid"
)

func UUID() string {
	return uuid.New().String()
}

func UUIDOnlyChar() string {
	return strings.ReplaceAll(UUID(), "-", "")
}
