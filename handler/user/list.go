package user

import (
	"gogin/model"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func List(c *gin.Context) {
	log.Info("List function called.")
	model.ListUser()

}
