package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-profiles/users"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetUsersEmpty(t *testing.T) {
	router := SetUpRouter()
	router.GET("/users", getUsers)

	request, _ := http.NewRequest("GET", "/users", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	var users []users.UserResponse
	json.Unmarshal(recorder.Body.Bytes(), &users)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Empty(t, users)
}

func TestPostUsers(t *testing.T) {
	router := SetUpRouter()
	router.POST("/users", postUsers)

	users := []users.UserRequest{{
		ID:          "1",
		Name:        "Jane Doe",
		DateOfBirth: "1989-04-29",
		CreatedOn:   0,
	}}

	jsonValue, _ := json.Marshal(users)

	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusCreated, recorder.Code)
}

//func TestPostImage(t *testing.T) {
//	router := SetUpRouter()
//	router.POST("/image", postImage)
//
//	openedFile, _ := os.Open("../image-utils/test-images/tall.jpeg")
//	defer openedFile.Close()
//
//	fileBytes, err := io.ReadAll(openedFile)
//	if err != nil {
//		panic(err)
//	}
//
//	request, _ := http.NewRequest("POST", "/image", bytes.NewBuffer(fileBytes))
//	request.Header.Set("Content-Type", "multipart/form-data")
//	request.Body.Read(fileBytes)
//
//	recorder := httptest.NewRecorder()
//	router.ServeHTTP(recorder, request)
//	assert.Equal(t, http.StatusCreated, recorder.Code)
//}
