package http

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/danluki/effective-mobile-test/internal/config"
	v1 "github.com/danluki/effective-mobile-test/internal/delivery/http/v1"
	"github.com/danluki/effective-mobile-test/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.New()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	rateLimitStore := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 5,
	})

	rateLimitMw := ratelimit.RateLimiter(rateLimitStore, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	router.Use(
		gin.Recovery(),
		sloggin.New(logger),
		cors.Default(),
		rateLimitMw,
	)

	err := router.SetTrustedProxies(nil)
	if err != nil {
		return nil
	}

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
