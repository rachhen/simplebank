package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/rachhen/simple_bank/db/sqlc"
)

// Server serves HTTP requests for our bank service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer create a new HTTP Server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.PATCH("/accounts/:id", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	server.router = router
	return server
}

// Start runs on HTTP server on the specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
