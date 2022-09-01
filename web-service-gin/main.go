package main

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
	imageutils "user-profiles/image-utils"
	"user-profiles/users"
)

func main() {

	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", postUsers)
	router.POST("/image", postImage)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}

// getUsers handles the GET /users requests by fetching
// the slice of UserResponses and serializes them into JSON
func getUsers(context *gin.Context) {
	convertedUsers := users.GetUsers()
	context.IndentedJSON(http.StatusOK, convertedUsers)
	return
}

// postUsers handles POST /users requests by parsing the JSON-encoded
// request into UserRequest structs that can be converted into UserResponse
// structs and then serialized into JSON for the response
func postUsers(c *gin.Context) {

	var newUsers []users.UserRequest

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(400, err)
	}

	err = json.Unmarshal(body, &newUsers)
	if err != nil {
		c.AbortWithError(400, nil)
		return
	}

	addedUsers := users.AddUsers(newUsers)
	c.IndentedJSON(http.StatusCreated, addedUsers)
}

// postImage handles POST /image requests by opening the raw file bytes
// passed in as part of a multipart/form-data Content-Type request. Will
// return a new PNG file.
func postImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	filetype, err := imageutils.GetFileContentType(file)
	if err != nil {
		panic(err)
	}

	// Get raw file bytes
	openedFile, _ := file.Open()
	defer openedFile.Close()
	fileBytes, err := io.ReadAll(openedFile)
	if err != nil {
		panic(err)
	}

	convertedFile := imageutils.GetNewPNGFileName(filepath.Base(file.Filename))

	switch filetype {
	case "image/jpeg", "image/jpg":
		imageutils.ConvertJPEGToPNGAndResize(fileBytes, convertedFile)
	case "image/png":
		imageutils.ResizePNG(fileBytes, convertedFile)
	default:
		invalidFiletypeError := errors.New("received invalid file type")
		c.AbortWithError(http.StatusBadRequest, invalidFiletypeError)
	}

	c.Header("Content-Type", "application/octet-stream")
	c.File(convertedFile)

	// Clean up the file
	removeErr := os.Remove(convertedFile)
	if removeErr != nil {
		panic(removeErr)
	}
	return
}
