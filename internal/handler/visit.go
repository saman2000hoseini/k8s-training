package handler

import (
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const KEY = "visits"

type Visit struct {
	Store *redis.Client
}

func (h Visit) Get(c echo.Context) error {
	if err := h.Store.Incr(KEY).Err(); err != nil {
		log.Errorf("incr key failed: %s", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	count, err := h.Store.Get(KEY).Result()
	if err != nil {
		log.Errorf("get key failed: %s", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	host, err := os.Hostname()

	return c.JSON(http.StatusOK, map[string]string{
		"hostname": host,
		"visits":   count,
	})
}

// Register registers the routes of URL handler on given group.
func (h Visit) Register(g *echo.Group) {
	g.GET("", h.Get)
}
