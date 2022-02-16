package tests

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go-todo/dto"
	"go-todo/routers/api"
	"io/ioutil"
	"net/http"
	"testing"
)

//func setCreateUserRouter(body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
//	r := gin.New()
//	r.POST("/users", api.CreateUser)
//	req, err := http.NewRequest(http.MethodPost, "/users", body)
//	if err != nil {
//		return req, httptest.NewRecorder(), err
//	}
//	req.Header.Set("Content-Type", "application/json")
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	return req, w, nil
//}

func Test_CreateUser(t *testing.T) {
	a := assert.New(t)
	user := dto.NewUserRequestDto{
		Username: "test",
		Email:    "test@testmail.com",
		Password: "nfenbfebebthtnggr",
	}
	reqBody, err := json.Marshal(user)
	if err != nil {
		a.Error(err)
	}
	req, w, err := postRouter("/users", bytes.NewBuffer(reqBody), api.CreateUser)
	if err != nil {
		a.Error(err)
	}
	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}
	// test the  dto
	actual := dto.NewUserResponseDto{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}
	expected := dto.NewUserResponseDto{
		Email:    "test@testmail.com",
		Username: "test",
	}
	a.Equal(expected, actual)
}

func Test_CreateUser_with_wrong_email_format(t *testing.T) {
	a := assert.New(t)
	user := dto.NewUserRequestDto{
		Username: "test",
		Email:    "test@testmail",
		Password: "nfenbfebebthtnggr",
	}
	reqBody, err := json.Marshal(user)
	if err != nil {
		a.Error(err)
	}
	req, w, err := postRouter("/users", bytes.NewBuffer(reqBody), api.CreateUser)
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusBadRequest, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}
	actual := map[string]interface{}{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := map[string]interface{}{
		"error": "Key: 'NewUserRequestDto.Email' Error:Field validation for 'Email' failed on the 'email' tag",
	}
	a.Equal(expected, actual)
}

func Test_Auth(t *testing.T) {
	a := assert.New(t)
	user := dto.UserAuthRequestDto{
		Email:    "test@testmail.com",
		Password: "nfenbfebebthtnggr",
	}
	reqBody, err := json.Marshal(user)
	if err != nil {
		a.Error(err)
	}
	req, w, err := postRouter("/users/login", bytes.NewBuffer(reqBody), api.Auth)
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	var response map[string]interface{}

	if err = json.Unmarshal(body, &response); err != nil {
		a.Error(err)
	}

	data, hasData := response["data"].(map[string]interface{})
	a.True(hasData)

	message, hasMessage := data["message"]
	_, hasToken := data["token"]

	a.True(hasMessage)
	a.True(hasToken)
	a.Equal("Authentication successful", message)

}

func Test_Auth_with_incorrect_email(t *testing.T) {
	a := assert.New(t)
	user := dto.UserAuthRequestDto{
		Email:    "wrong@testmail.com",
		Password: "nfenbfebebthtnggr",
	}
	reqBody, err := json.Marshal(user)
	if err != nil {
		a.Error(err)
	}
	req, w, err := postRouter("/users/login", bytes.NewBuffer(reqBody), api.Auth)
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusBadRequest, w.Code, "HTTP request status code error")
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	var response map[string]interface{}

	if err = json.Unmarshal(body, &response); err != nil {
		a.Error(err)
	}

	value, exists := response["error"]
	a.True(exists)
	a.Equal("invalid login details", value)
}
