package engine

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var onceApp sync.Once
var appInstance *application

type application struct {
	engine *gin.Engine
}

func GetEngine() *gin.Engine {
	onceApp.Do(func() {
		appInstance = &application{
			engine: gin.Default(),
		}
	})
	return appInstance.engine
}
