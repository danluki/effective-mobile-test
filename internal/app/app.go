package app

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SteelPangolin/go-genderize"
	"github.com/danluki/effective-mobile-test/internal/config"
	delivery "github.com/danluki/effective-mobile-test/internal/delivery/http"
	"github.com/danluki/effective-mobile-test/internal/logger"
	"github.com/danluki/effective-mobile-test/internal/repository"
	"github.com/danluki/effective-mobile-test/internal/server"
	"github.com/danluki/effective-mobile-test/internal/service"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/masonkmeyer/agify"
	"github.com/masonkmeyer/nationalize"
	"github.com/pressly/goose"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// Run initializes whole application.
func Run(configPath string) {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		slog.Error("Error:", slog.Any("err", err))
	}

	if cfg.AppEnv == "development" {
		gin.SetMode("debug")
	}

	gin.SetMode(gin.ReleaseMode)

	runMigrations(cfg)

	db, err := gorm.Open(
		postgres.Open(cfg.DatabaseUrl), &gorm.Config{
			Logger: gormLogger.Default.LogMode(gormLogger.Silent),
		},
	)
	if err != nil {
		panic(err)
	}

	d, _ := db.DB()
	d.SetMaxOpenConns(5)
	d.SetConnMaxIdleTime(1 * time.Minute)

	genderizeClient, err := genderize.NewClient(genderize.Config{
		APIKey:    "",
		UserAgent: "GoGenderize",
	})
	if err != nil {
		panic(err)
	}

	clogger := logger.New(logger.Opts{
		Env:     cfg.AppEnv,
		Service: "app",
	})

	services := service.NewServices(service.Deps{
		Repositories:      repository.NewRepositories(db),
		GenderizeClient:   genderizeClient,
		NationalizeClient: nationalize.NewClient(),
		AgifyClient:       agify.NewClient(),
		Logger:            clogger,
	})

	handlers := delivery.NewHandler(services)
	if handlers == nil {
		panic("Failed to create handlers")
	}

	srv := server.NewServer(cfg, handlers.Init(cfg))
	go func() {
		clogger.Info("Server has been started on:", slog.Any("addr", cfg.HttpServerAddress))

		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	_, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()
}

func runMigrations(cfg *config.Config) {
	pgDB, err := sql.Open("postgres", cfg.DatabaseUrl)
	if err != nil {
		panic(err)
	}
	defer pgDB.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(pgDB, cfg.MigrationsPath); err != nil {
		panic(err)
	}
}
