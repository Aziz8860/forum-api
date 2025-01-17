package main

import (
	"log"

	"github.com/aziz8860/forum-api/internal/configs"
	"github.com/aziz8860/forum-api/internal/handlers/memberships"
	"github.com/aziz8860/forum-api/internal/handlers/posts"
	"github.com/aziz8860/forum-api/pkg/internalsql"
	"github.com/gin-gonic/gin"

	membershipRepo "github.com/aziz8860/forum-api/internal/repository/memberships"
	membershipSvc "github.com/aziz8860/forum-api/internal/service/memberships"

	postRepo "github.com/aziz8860/forum-api/internal/repository/posts"
	postSvc "github.com/aziz8860/forum-api/internal/service/posts"
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

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiasi database", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	postHandler := posts.NewHandler(r, postService)
	postHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
