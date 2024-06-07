package gears

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"{{cookiecutter.project_name}}/store/database"
)

func HandleError(c *gin.Context, err error) {
	if msg := database.IsPostgresError(err); msg != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
}
