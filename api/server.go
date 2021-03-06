package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/template-go-server/db/sqlc"
	"github.com/template-go-server/token"
	"github.com/template-go-server/util"
)

type Server struct {
	config     util.Config
	router     *gin.Engine
	tokenMaker token.Maker
	store      db.Store
}

// NewServer create a new HTTP server and setup routing
func NewServer(store db.Store, config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.GET("/", server.getHello)

	router.GET("/authors/:id", server.getAuthor)
	router.GET("/authors", server.listAuthors)
	router.POST("/authors", server.createAuthor)

	router.GET("/genres", server.listGenres)
	router.GET("/genres/:id", server.getGenre)
	router.POST("/genres", server.createGenre)

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
