package routes

import (
	"arth/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetView(context *gin.Context) {
	path := context.Param("path")
	pathString := models.GenerateViewPath(path)
	file, err := os.ReadFile(pathString)

	if err != nil {
		fmt.Println("Error Opening File")
		fmt.Println(err)
		context.JSON(http.StatusNotFound, file)
		return
	}

	var viewData any
	err = json.Unmarshal(file, &viewData)

	if err != nil {
		fmt.Println("Error Parsing View Data")
		fmt.Println(err)
		context.JSON(http.StatusNotFound, gin.H{"error": err, "message:": "cannot parse view data"})
		return
	}
	context.JSON(http.StatusOK, viewData)

}
