package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/codeedu/codebank/domain"
	"github.com/codeedu/codebank/infra/repository"
	"github.com/codeedu/codebank/usecase"
	_ "github.com/lib/pq"
)

func main() {
	db := setupDb()
	defer db.Close()

	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "Math"
	cc.ExpirationMonth = 1
	cc.ExpirationYear = 2022
	cc.CVV = 123
	cc.Limit = 1000
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDb(db)
	err := repo.CreateCreditCard(*cc)

	if err != nil {
		fmt.Println(err)
	}
}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)

	return useCase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
		"db",
		5432,
		"postgres",
		"root",
		"codebank")

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("error connecting to the database")
	}

	return db
}