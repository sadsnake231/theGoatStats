package routes

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSeasonStats() gin.HandlerFunc {
	return func(c *gin.Context){
		c.JSON(http.StatusTeapot, "")
	}
}

func GetLastMatchStats() gin.HandlerFunc {
	return func(c *gin.Context){
		c.JSON(http.StatusTeapot, "")
	}
}
