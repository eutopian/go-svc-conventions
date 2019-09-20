package components

import (
	"github.com/labstack/echo/v4"
	"github.com/geomyidia/go-svc-conventions/components/config"
	"github.com/geomyidia/reverb"
	logger "github.com/sirupsen/logrus"
)

// Base component collection
type Base struct {
	Config *config.Config
	Logger *logger.Logger
}

// TestBase component that keeps stdout clean
type TestBase struct {
	Config *config.Config
}

// Default component collection
type Default struct {
	Base
	HTTPD  *echo.Echo
	GRPCD  *reverb.Reverb
}

// Add more components here that have more or less than what's done above. This
// is useful for testing or runnning in different binaries/executables, etc.
