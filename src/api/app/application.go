package app

import (
	"flay-api-v3.0/src/api/infraestructure/dependencies"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

const (
	port = ":8080"
)

func Start() {
	router = gin.Default()
	handlers := dependencies.Start()
	configureMappings(router, handlers)
	_ = router.Run(port)
}
