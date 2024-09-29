package api

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidate_Pointer(t *testing.T) {
	username := "admin"
	password := "123456"

	err := Validate(
		V("username", &username, "omitempty"),
		V("password", &password, "omitempty"),
	)

	require.Nil(t, err)
}

func TestValidate_NotExist(t *testing.T) {
	username := "admin"
	var password *string
	password = nil

	err := Validate(
		V("username", &username, "omitempty,min=3,max=20"),
		V("password", &password, "omitempty,min=6,max=20"),
	)

	require.Nil(t, err)
}

func TestValidate_Error(t *testing.T) {
	username := "admin"
	password := "123"

	err := Validate(
		V("username", &username, "omitempty,max=2"),
		V("password", &password, "omitempty,min=6,max=20"),
	)

	require.NotNil(t, err)
}

func TestValidate_EmptyUsername(t *testing.T) {
	var username string
	password := "123456"

	err := Validate(
		V("username", &username, "min=3,max=20"),
		V("password", &password, "min=6,max=20"),
	)

	require.NotNil(t, err)
}
