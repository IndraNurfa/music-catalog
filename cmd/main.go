package main

import (
	"log"
	"net/http"

	"github.com/IndraNurfa/music-catalog/internal/configs"
	membershipsHandler "github.com/IndraNurfa/music-catalog/internal/handler/memberships"
	tracksHandler "github.com/IndraNurfa/music-catalog/internal/handler/tracks"
	"github.com/IndraNurfa/music-catalog/internal/models/memberships"
	membershipsRepo "github.com/IndraNurfa/music-catalog/internal/repository/memberships"
	"github.com/IndraNurfa/music-catalog/internal/repository/spotify"
	membershipsSvc "github.com/IndraNurfa/music-catalog/internal/service/memberships"
	"github.com/IndraNurfa/music-catalog/internal/service/tracks"
	"github.com/IndraNurfa/music-catalog/pkg/httpclient"
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

	httpClient := httpclient.NewClient(&http.Client{})

	spotifyOutbound := spotify.NewSpotifyOutbound(cfg, httpClient)

	membershipRepo := membershipsRepo.NewRepository(db)

	membershipSvc := membershipsSvc.NewService(cfg, membershipRepo)
	trackSvc := tracks.NewService(spotifyOutbound)

	membershipsHandler := membershipsHandler.NewHandler(r, membershipSvc)
	membershipsHandler.RegisterRoute()

	tracksHandler := tracksHandler.NewHandler(r, trackSvc)
	tracksHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
