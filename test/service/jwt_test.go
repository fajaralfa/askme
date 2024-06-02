package service_test

import (
	"testing"

	"github.com/fajaralfa/askme/internal/service/jwt"
)

func TestCreateJwt(t *testing.T) {
	token, err := jwt.Create("fajaralfa@gmail.com")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("token:", token)
}

func TestParseJwt(t *testing.T) {
	tokenString, err := jwt.Create("fajaralfa@gmail.com")
	if err != nil {
		t.Fatal(err)
	}

	token, err := jwt.Parse(tokenString)
	if err != nil {
		t.Fatal(err)
	} else if claims, ok := token.Claims.(*jwt.Claims); ok {
		t.Log(claims.Email)
	} else {
		t.Fatal("Unknown claim types")
	}
}
