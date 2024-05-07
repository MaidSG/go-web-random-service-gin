package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


type response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

// 获取随机数
func getRandomNumber() int {
	return rand.Intn(100) + 10
}

// 生成表达式
func generateExpression(a,b int) string {
	if rand.Intn(2) == 0 {
		return strconv.Itoa(a) + "+" + strconv.Itoa(b)
	} else {
		return strconv.Itoa(a) + "-" + strconv.Itoa(b)
	}
}


// 创造随机的加减法题目
func getRandomAdditonOrSubtraction() string {
	a := getRandomNumber()
    b := getRandomNumber()
    return generateExpression(a, b)
}

func createResponse(c *gin.Context,success bool, message string, data interface{}) {
     c.JSON(http.StatusOK, response{
		Success: success,
		Message: message,
		Data: data,
	})
}

// 获取随机的加减法题目
func getRandomAdditonOrSubtractionHandler(c *gin.Context) {
	rand_type := c.Param("type")

	switch rand_type {
		case "add-or-sub":
			createResponse(c,true, "success", getRandomAdditonOrSubtraction())
		default:
			createResponse(c,false, "type not found", nil)
	}
	
}

func main() {
	r := gin.Default()
	r.GET("/random-str/:type", getRandomAdditonOrSubtractionHandler)
	r.Run(":8089")
}