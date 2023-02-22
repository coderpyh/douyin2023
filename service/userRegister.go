package service

import (
	"crypto/sha256"
	"douyin2023/repository"
	"encoding/hex"
	"errors"
)

type UserRegisterRequest struct {
	Name     string `query:"username"`
	Password string `query:"password"`
}

type UserRegisterResponse struct {
	UserID int
	Token  string
}

func UserRegister(request *UserRegisterRequest) (*UserRegisterResponse, error) {
	err := userRegisterCheckRequest(request)
	if err != nil {
		return nil, err
	}
	response, err := userRegisterPackResponse(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func userRegisterCheckRequest(request *UserRegisterRequest) error {
	if len(request.Name) == 0 || len(request.Password) == 0 ||
		len(request.Name) > 32 || len(request.Password) > 32 {
		return errors.New("invalid param")
	}
	return nil
}

// 密码计算加盐hash值存入数据库
func userRegisterPackResponse(request *UserRegisterRequest) (*UserRegisterResponse, error) {
	passwordHashByte := sha256.Sum256([]byte("a" + request.Password))
	passwordHash := hex.EncodeToString(passwordHashByte[:])
	uid, err := repository.UserInsert(request.Name, passwordHash)
	if err != nil {
		return nil, err
	}
	var response UserRegisterResponse
	response.UserID = uid
	response.Token = uidToToken(uid)
	return &response, nil
}
