package main

import (
	"fmt"
	"net/http"

	"adilkhan.go-clean-architechture/pkg/store/postgres"
	"adilkhan.go-clean-architechture/services/contact/internal/delivery"
	"adilkhan.go-clean-architechture/services/contact/internal/repository"
	usecase "adilkhan.go-clean-architechture/services/contact/internal/useCase"
)

func main() {
	dcp := &postgres.DbConnParams{
		Host:     "localhost",
		Port:     5432,
		User:     "architecture",
		Password: "clean123",
		DbName:   "go_clean_architechture",
	}

	db, err := postgres.OpenDB(dcp)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Pool.Close()

	repo := repository.New(db.Pool)
	delivery := delivery.New()
	usecase := usecase.New(repo)

	_ = usecase

	fmt.Println("app started")
	http.ListenAndServe("localhost:4000", delivery.Mux)
}
