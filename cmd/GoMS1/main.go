package main

import (
	"fmt"
	"log"
	"net"

	interfaces "github.com/felixcaronpare/GoMS1/pkg/v1"
	repo "github.com/felixcaronpare/GoMS1/pkg/v1/repository"
	usecase "github.com/felixcaronpare/GoMS1/pkg/v1/usecase"
	"gorm.io/gorm"

	dbConfig "github.com/felixcaronpare/GoMS1/internal/db"
	"github.com/felixcaronpare/GoMS1/internal/models"
	handler "github.com/felixcaronpare/GoMS1/pkg/v1/handler/GoMS1"
	"google.golang.org/grpc"
)

func main() {
	// connect to the db
	db := dbConfig.dbConn()
	migrations(db)

	// add a listener address
	lis, err := net.Listen ("tcp", ":5001")
	if err != nil {
		log.Fatalf("ERROR STARTING THE SERVER : %v", err)
	}

	// start the grpc server
	grpcServer := grpc.NewServer()

	accountUsecase := initAccountServer(db)
	handler.NewServer(grpcServer, accountUsecase)

	// start serving to the address
	log.Fatal(grpcServer.Serve(lis))
}

func initAccountServer(db *gorm.DB) interfaces.UsecaseInterface {
	accountRepo := repo.New(db)
	return usecase.New(accountRepo)
}

func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.Account{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migrated")
	}
}