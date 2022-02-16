package tests

import (
	"github.com/gin-gonic/gin"
	"go-todo/models"
	"go-todo/pkg/settings"
	"go-todo/routers"
	"os"
	"testing"
)

var r *gin.Engine

func init() {
	settings.InitSettings()
	models.Setup()
}

func testMain(m *testing.M) int {
	r = routers.InitRourter()
	return m.Run()
}

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}
