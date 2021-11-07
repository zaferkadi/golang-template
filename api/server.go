package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/template-go-server/db/sqlc"
)

type Server struct {
	router *gin.Engine
	store  db.Store
}

// NewServer create a new HTTP server and setup routing
func NewServer(store db.Store) (*Server, error) {

	server := &Server{
		store: store,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.GET("/", server.getHello)
	router.POST("/authors", server.createAuthor)
	router.GET("/authors/:id", server.getAuthor)
	router.GET("/authors", server.listAuthors)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
