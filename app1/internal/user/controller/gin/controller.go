package gin

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	List(c *gin.Context)
}
