package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/template-go-server/db/sqlc"
	"github.com/template-go-server/util"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type createUserResponse struct {
	Username string `json:"username" `
	Fullname string `json:"fullname" `
	Email    string `json:"email" `
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		FullName:       req.Fullname,
		HashedPassword: hashedPassword,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}

		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := createUserResponse{
		Username: user.Username,
		Fullname: user.FullName,
		Email:    user.Email,
	}
	ctx.JSON(http.StatusOK, resp)
}
