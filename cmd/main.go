package main

import (
	"awesomeProject2/internal/repo"
	"awesomeProject2/internal/service"
	"awesomeProject2/internal/transport/rest"
	"awesomeProject2/pkg/hash"
	"awesomeProject2/pkg/postgres"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db, err := postgres.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	hasher := hash.NewSHA1Hasher("salt")
	tokensRepo := repo.NewTokens(db)
	usersRepo := repo.NewUsers(db)
	clientRepo := repo.NewClientPostgres(db)
	agentRepo := repo.NewAgentPostgres(db)
	ownerRepo := repo.NewOwnerPostgres(db)
	usersService := service.NewUsers(usersRepo, tokensRepo, hasher, []byte("sample"))
	handler := rest.NewHandler(usersService, clientRepo, agentRepo, ownerRepo)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8000),
		Handler: handler.InitRouter(),
	}
	log.Println("Server Started")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
