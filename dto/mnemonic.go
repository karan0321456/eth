package dto

import (
	"github.com/eth/errs"
	"github.com/tyler-smith/go-bip39"
)



func GenerateMnemonic()(string,*errs.AppError){
	entropy,err := bip39.NewEntropy(256)
	if err != nil {
		return "",&errs.AppError{Code:500,Message: "mnemonic is not created"}
	}
	mnemonic,err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "",&errs.AppError{Code:500,Message: "mnemonic is not created"}
	}

	return mnemonic,nil
}