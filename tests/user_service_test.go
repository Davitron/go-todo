package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-todo/config"
	"go-todo/core/db"
	user2 "go-todo/users"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

var user *user2.UserService

func setCreateUserRouter(db *gorm.DB, body *bytes.Buffer, cfg *config.Configurations) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()
	user.DB = db
	user.Config = cfg
	r.POST("/users", user.CreateUser)
	req, err := http.NewRequest(http.MethodGet, "/users", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil
}

func Test_Create_User(t *testing.T) {
	Config, _ = cfg.Init("../.")
	TestDB = db.InitDB(Config.Database)
	a := assert.New(t)
	{
		testUser := user2.User{}
		reqBody, err := json.Marshal(testUser)
		if err != nil {
			a.Error(err)
		}
		req, w, err := setCreateUserRouter(TestDB, bytes.NewBuffer(reqBody), Config)
		if err != nil {
			a.Error(err)
		}
		a.Equal(http.MethodPost, req.Method, "HTTP request method error")
		a.Equal(http.StatusBadRequest, w.Code, "HTTP request status code error")
	}
}
