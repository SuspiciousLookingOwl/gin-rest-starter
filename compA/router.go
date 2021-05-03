package compA

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(api *gin.RouterGroup) {
	router := api.Group("/compA")

	router.GET("/", GetHelloWorld)
	router.GET("/:x/:y", GetSum)
}

// ShowAccount godoc
// @Summary Hello there!
// @Description Returns "Hello World"
// @ID get-hello-world
// @Success 200 {string} string
// @Router /compA [get]
func GetHelloWorld(c *gin.Context) {
	c.String(200, "Hello World")
}

// Example custom defined struct for JSON Response
type Result struct {
	Result int `json:"result"`
}

// Sum 2 number godoc
// @Summary Sum 2 numbers
// @Description Get sum of 2 numbers
// @ID get-sum-by-int
// @Produce  json
// @Param x path int true "First number"
// @Param y path int true "Second number"
// @Success 200 {object} Result
// @Failure 400
// @Router /compA/{x}/{y} [get]
func GetSum(c *gin.Context) {
	x, xErr := strconv.Atoi(c.Params.ByName("x"))
	y, yErr := strconv.Atoi(c.Params.ByName("y"))
	if xErr != nil || yErr != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": Sum(x, y),
	})
}
