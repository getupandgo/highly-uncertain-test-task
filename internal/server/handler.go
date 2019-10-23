package server

import (
	"github.com/getupandgo/highly-uncertain-test-task/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	H string
	K float32
}

func (s Server) BaseSubstitution(c *gin.Context) {
	i := model.Input{}
	if err := c.BindJSON(&i); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	state, out, err := s.ctrl.CalculateSubstitution("base", i)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{
		H: state,
		K: out,
	})
}

func (s Server) CustomSubstitution(c *gin.Context) {
	i := model.Input{}
	if err := c.BindJSON(&i); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	ruleNumber := c.Param("id")
	state, out, err := s.ctrl.CalculateSubstitution("custom"+ruleNumber, i)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{
		H: state,
		K: out,
	})
}
