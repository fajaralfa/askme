package service_test

import (
	"strings"
	"testing"

	hp "github.com/fajaralfa/askme/internal/service/hash"
)

var hash hp.Hash = hp.Hash{}

func TestMakeHash(t *testing.T) {
	text := "rahasia"
	dk, err := hash.Make(text)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(dk, "$") {
		t.Fatal("hashed result format wrong:", dk)
	}
}

func TestVerifyHashSuccess(t *testing.T) {
	password := "rahasia"
	dk, _ := hash.Make(password)
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
