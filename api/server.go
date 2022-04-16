package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/tornvallalexander/goreddit/db/sqlc"
)

// Server serves HTTP requests for GoReddit
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server with routing
func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}

	server.setupRouter()
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.GET("/users/:username", server.getUser)

	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
