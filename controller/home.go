package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func MainController(context *gin.Context) {
	context.HTML(http.StatusOK, "Front/index.html", gin.H{})
}