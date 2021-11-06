package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getHello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "this worked"})

}
