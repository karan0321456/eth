package app

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/eth/domain"
	"github.com/eth/errs"
	"github.com/eth/logger"
	"github.com/eth/service"
	"github.com/jmoiron/sqlx"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

var infuraUrl = "https://mainnet.infura.io/v3/977ca2d19b68405885d4aafbc40ad7d6"

func Start(){
	r :=mux.NewRouter()

	dbClient := getDbClient()
	//wiring
	newRepositoryDb := domain.NewUserRepositoryDb(dbClient)
	us := UserHandler{service.NewUserService(newRepositoryDb)}
	
	r.HandleFunc("/createWallet",us.register)
	http.ListenAndServe("localhost:8003",r)
}

func EthClient()*ethclient.Client{
	client,err:=ethclient.Dial(infuraUrl)
	if err != nil{
		errs.NewUnexpectedError("Ether Client Not Created")
	}
	defer client.Close()
	return client
}

func getDbClient() *sqlx.DB{
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	db:= os.Getenv("DB")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)

	client, err := sqlx.Open(db, dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 10)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	logger.Info("Database is Connected")
	return client
}


