// package server

// import (
// 	"log"

// 	"cms.golang.tnb.api/server/routes"
// 	"github.com/gin-gonic/gin"
// )

// type Server struct {
// 	port   string
// 	server *gin.Engine
// }

// func NewServer() Server {
// 	return Server{
// 		port:   "8000",
// 		server: gin.Default(),
// 	}
// }

// func (s *Server) Run() {
// 	router := routes.ConfigRoutes(s.server)
// 	log.Printf("Server running at port: %v", s.port)
// 	log.Fatal(router.Run(":" + s.port))
// }
