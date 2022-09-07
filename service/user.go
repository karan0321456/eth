package service

import (
	"github.com/eth/domain"
	"github.com/eth/dto"
	"github.com/eth/errs"
	"github.com/google/uuid"
)


type UserService interface{
	User(string)(*dto.UserResponse ,*errs.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepositoryDb
}

func(d DefaultUserService) User(password string)(*dto.UserResponse ,*errs.AppError){
	mnemonic,err := dto.GenerateMnemonic()
	if err != nil {
		errs.NewStatusInternalServerError("mnemonic is not created")
		return nil,err
	}

	userId := uuid.NewString()
	error := d.repo.AddUser(userId,mnemonic,password)
	if error != nil {
		return nil,error
	}
	return &dto.UserResponse{Mnemonic:mnemonic},nil
}

func NewUserService(repository domain.UserRepositoryDb)DefaultUserService{
	return DefaultUserService{repository}
}