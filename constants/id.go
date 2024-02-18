package constants

import (
	"strings"

	"github.com/oklog/ulid/v2"
)

type DataPrefix string

const (
	// user related prefixes
	User DataPrefix = "user_"

	// session related prefixes
	Session DataPrefix = "sess_"
)

func (dp DataPrefix) String() string {
	return string(dp)
}

func (dp DataPrefix) IsValid(s string) bool {
	return strings.HasPrefix(s, string(dp)) && len(s) == len(string(dp))+ulid.EncodedSize
}

func GenerateDataPrefixWithULID[T DataPrefix](prefixType T) string {
	return string(prefixType) + ulid.Make().String()
}
