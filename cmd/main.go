package main

import (
	"log"

	"github.com/aziz8860/forum-api/internal/configs"
	"github.com/aziz8860/forum-api/internal/handlers/memberships"
	"github.com/aziz8860/forum-api/pkg/internalsql"
	"github.com/gin-gonic/gin"

	membershipRepo "github.com/aziz8860/forum-api/internal/repository/memberships"
	membershipSvc "github.com/aziz8860/forum-api/internal/service/memberships"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs/"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Gagal inisiasi config")
	}
	cfg = configs.Get()
	log.Println("config", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiasi database", err)
	}

	membershipRepo := membershipRepo.NewRepository(db)
	membershipService := membershipSvc.NewService(membershipRepo)

	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
