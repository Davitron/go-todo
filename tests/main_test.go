package tests

import (
	"go-todo/api"
	cfg "go-todo/config"
	"go-todo/core/db"
	"gorm.io/gorm"
	"os"
	"testing"
)

var app *api.Server
var Config *cfg.Configurations
var TestDB *gorm.DB

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	Config, _ = cfg.Init("../.")
	TestDB = db.InitDB(Config.Database)
	return m.Run()
}
