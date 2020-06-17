package k8s

import (
	. "gogin/handler"
	"gogin/model"

	"github.com/gin-gonic/gin"
)

func Monitoring(c *gin.Context) {

	model.Monitoring(c.Param("ns"))
	SendResponse(c, nil, nil)
}
