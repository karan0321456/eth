package domain

import (
	"github.com/eth/errs"
	"github.com/eth/logger"


	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type UserRepositoryDb struct{
	client *sqlx.DB
}


var path  = "m/44H/60H/0H/0/0"

func (d UserRepositoryDb)AddUser(UserID,Mnemonic,passowrd string)(*errs.AppError){
	psqlInsert := "INSERT INTO users(user_id,mnemonic,password,path)values($,$,$,$)"
	_,err := d.client.Exec(psqlInsert,UserID,Mnemonic,passowrd,path) 
	if err!=nil{
		logger.Error("Error While creating new account"+err.Error())
		return errs.NewUnexpectedError("Unexpected error from database")
	}
	return nil
}

func NewUserRepositoryDb(dbCLient *sqlx.DB)UserRepositoryDb{
	return UserRepositoryDb{dbCLient}
}