package main

import (
	"fmt"
	"os"

	"github.com/go-errors/errors"
	"github.com/rickbassham/goapi/middleware"
	"github.com/rickbassham/goapi/router"
	"github.com/rickbassham/logging"

	"github.com/bassham-aws/api-test/environment"
	"github.com/bassham-aws/api-test/handler"
	"github.com/bassham-aws/api-test/routes"
)

var (
	buildVersion = "UNKNOWN"
	buildDate    = "UNKNOWN"
)

func main() {
	statusCode := mainWithStatusCode()
	os.Exit(statusCode)
}

func mainWithStatusCode() (statusCode int) {
	logger := logging.NewLogger(os.Stdout, logging.JSONFormatter{}, logging.LogLevelInfo).WithField("buildVersion", buildVersion).WithField("buildDate", buildDate)
	logger.Info("starting")

	defer func() {
		if r := recover(); r != nil {
			statusCode = 1

			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}

			wrapped := errors.Wrap(err, 2)

			logger.WithField("stackTrace", wrapped.StackFrames()).WithError(err).Error("recovered from panic; exiting")
		}

		logger.WithField("statusCode", statusCode).Info("done")
	}()

	requestLogger := middleware.NewRequestLogger(logger)
	h := handler.NewHandlerService(requestLogger)

	r := routes.NewAPI(h)

	listenAddress := environment.ListenAddress()
	router := router.NewRouter(listenAddress, requestLogger, r)

	logger.WithField("listenAddress", listenAddress).Info("listening...")

	err := router.ListenAndServe()
	if err != nil {
		logger.WithError(err).Error("error in router.ListenAndServe")
		statusCode = 1
		return
	}

	statusCode = 0
	return
}
