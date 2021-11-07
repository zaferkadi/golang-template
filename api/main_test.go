package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	// this will change the debug mode of gin for testing
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
