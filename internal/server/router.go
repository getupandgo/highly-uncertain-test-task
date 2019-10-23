package server

import (
	"github.com/getupandgo/highly-uncertain-test-task/internal/model"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	CalculateSubstitution(rule string, input model.Input) (string, float32, error)
}

type Server struct {
	Engine *gin.Engine
	ctrl   Controller
}

func Init(ctrl Controller) *Server {
	router := gin.New()

	s := Server{Engine: router, ctrl: ctrl}

	substitutionRouter := router.Group("/substitution")
	substitutionRouter.POST("/", s.BaseSubstitution)
	substitutionRouter.POST("/custom/:id", s.CustomSubstitution)

	return &s
}
