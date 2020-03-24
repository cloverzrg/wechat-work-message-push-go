package api

import (
	"github.com/gin-gonic/gin"
)

func Start() (err error) {
	r := gin.New()
	SetRoute(r)
	err = r.Run("0.0.0.0:80")
	return err
}
