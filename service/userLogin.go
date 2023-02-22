package service

import (
	"crypto/sha256"
	"douyin2023/repository"
	"encoding/hex"
	"errors"
)

type UserLoginRequest struct {
	Name     string `query:"username"`
	Password string `query:"password"`
}

type UserLoginResponse struct {
	UserID int
	Token  string
}

func UserLogin(request *UserLoginRequest) (*UserLoginResponse, error) {
	err := userLoginCheckRequest(request)
	if err != nil {
		return nil, err
	}
	response, err := userLoginPackResponse(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func userLoginCheckRequest(request *UserLoginRequest) error {
	if len(request.Name) == 0 || len(request.Password) == 0 ||
		len(request.Name) > 32 || len(request.Password) > 32 {
		return errors.New("invalid param")
	}
	return nil
}

func userLoginPackResponse(request *UserLoginRequest) (*UserLoginResponse, error) {
	passwordHashByte := sha256.Sum256([]byte("a" + request.Password))
	passwordHash := hex.EncodeToString(passwordHashByte[:])
	uid, err := repository.UserSelectNamePassword(request.Name, passwordHash)
	if err != nil {
		return nil, err
	}
	var response UserLoginResponse
	response.UserID = uid
	response.Token = uidToToken(uid)
	return &response, nil
}
