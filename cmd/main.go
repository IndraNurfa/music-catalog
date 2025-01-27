package main

import (
	"log"

	"github.com/IndraNurfa/music-catalog/internal/configs"
	membershipsHandler "github.com/IndraNurfa/music-catalog/internal/handler/memberships"
	"github.com/IndraNurfa/music-catalog/internal/models/memberships"
	membershipsRepo "github.com/IndraNurfa/music-catalog/internal/repository/memberships"
	membershipsSvc "github.com/IndraNurfa/music-catalog/internal/service/memberships"
	"github.com/IndraNurfa/music-catalog/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs/"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}

	cfg = configs.Get()
	log.Println("config", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database, error: %v", err)
	}

	db.AutoMigrate(&memberships.User{})

	r := gin.Default()

	membershipRepo := membershipsRepo.NewRepository(db)

	membershipSvc := membershipsSvc.NewService(cfg, membershipRepo)

	membershipsHandler := membershipsHandler.NewHandler(r, membershipSvc)
	membershipsHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
