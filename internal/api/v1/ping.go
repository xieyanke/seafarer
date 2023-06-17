package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ping struct{}

func NewPing() *Ping {
	return &Ping{}
}

// ping godoc
// @Summary ping...
// @Description do ping
// @Tags Demo
// @Accept plain
// @produce plain
// @Success 200 {string} pong
// @Router	/ping [get]
func (p *Ping) Get(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
