package hash

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// dk is derived key

func makeWithSalt(text string, salt []byte) string {
	dk := argon2.Key([]byte(text), salt, 1, 64*1024, 4, 32)

	b64salt := base64.RawStdEncoding.EncodeToString(salt)
	b64dk := base64.RawStdEncoding.EncodeToString(dk)

	result := fmt.Sprintf("%v$%s", string(b64salt), string(b64dk)) // format: salt$derivedKey
	return result
}

func Make(text string) (string, error) {
	salt := make([]byte, 8)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	dk := makeWithSalt(text, salt)
	return dk, nil
}

func Verify(text, savedDk string) bool {
	salt := strings.Split(savedDk, "$")[0]
	saltBytes, _ := base64.RawStdEncoding.DecodeString(salt)

	currentDk := makeWithSalt(text, saltBytes)
	return currentDk == savedDk
}
