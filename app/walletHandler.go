package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"strconv"

	// "crypto/sha1"
	// "encoding/hex"

	// "github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/miguelmota/go-ethereum-hdwallet"
	// "github.com/tyler-smith/go-bip39"
	// "github.com/ethereum/go-ethereum/common/hexutil"
	// "github.com/ethereum/go-ethereum/crypto"
) 

var mnemonic = "fog rabbit brisk search valve tilt chat marine front throw uniform tree spare boat casino"

func CreateWallet(w http.ResponseWriter,r *http.Request){
	// seed := bip39.NewSeed(mnemonic,"")
	// fmt.Println(seed)
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		// log.Fatal(err)
	}

	pathChecking := "m/44'/60'/0'/0/0"
	paths := strings.Split(pathChecking, "/")


	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		// log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947

	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account, err = wallet.Derive(path, false)
	if err != nil {
		// log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) 

}

func pathChecking(){
	pathChecking := "m/44'/60'/0'/0/0"
	paths := strings.Split(pathChecking, "/")
	fmt.Println(paths)
	lastElement := paths[len(paths)-1]
	lE, _ := strconv.Atoi(lastElement)
	if lE == 0 {
		lE = +1
	}
	ch := strconv.Itoa(lE)
	paths[len(paths)-1] = ch

	th := strings.Join(paths, "/")

	fmt.Println(th)
}



func writeReponse(w http.ResponseWriter,code int,data interface{}){
	w.Header().Add("Content-Type","application/json")
		w.WriteHeader(code)
		if err:=json.NewEncoder(w).Encode(data);err!=nil{
			panic(err)
		}
}