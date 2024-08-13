package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PingController struct {
}

func NewPingController() *PingController {
	c := &PingController{}

	return c
}

func (c *PingController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":     "pong",
		"status":      "up",
		"retrievedAt": time.Now(),
	})
}
