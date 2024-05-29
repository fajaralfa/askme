package hash

import (
	"encoding/base64"
	"strings"
	"testing"

	"github.com/fajaralfa/askme/internal/service/hash"
)

func TestMakeHash(t *testing.T) {
	text := "rahasia"
	dk, err := hash.Make(text)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("derived key:", dk)
	t.Log("derived key length:", len(dk))
	if !strings.Contains(dk, "$") {
		t.Fatal("hashed result format wrong:", dk)
	}
}

func TestVerifyHashSuccess(t *testing.T) {
	password := "rahasia"
	dk, _ := hash.Make(password)

	salt := strings.Split(dk, "$")[0]
	saltBytes, _ := base64.RawStdEncoding.DecodeString(salt)
	t.Log("salt:", saltBytes)
	t.Log("derived key:", dk)

	if hash.Verify(password, dk) != true {
		t.Fatal("Hash result not match")
	}
}

func TestVerifyHashWrongPass(t *testing.T) {
	password := "rahasia"
	dk, _ := hash.Make(password)
	if hash.Verify("salah", dk) == true {
		t.Fatal("Hash result not match")
	}
}
