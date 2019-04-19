package controller_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"io"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	c "github.com/avcwisesa/foreign_currency/controller"
)

type ControllerSuite struct {
	suite.Suite
	controller c.Controller
}

func performRequest(r *gin.Engine, method, url string, body io.Reader, ctx context.Context) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, url, body)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, &ControllerSuite{})
}

func (cs *ControllerSuite) SetupSuite() {
	d := &DBMock{}

	cs.controller = c.New(d)
}

func (cs *ControllerSuite) TestAddExchangeRateOnSuccess() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	r := gin.Default()
	r.GET("/exchangeRate/add", cs.controller.AddExchangeRate)
	w := performRequest(
		r,
		"GET",
		"/exchangeRate/add?from=JPY&to=IDR&date=2019-01-01T00:00:00.000Z&rate=0.1231",
		nil,
		ctx,
	)

	assert.Equal(cs.T(), 200, w.Code, "Code should be 200")
}

func (cs *ControllerSuite) TestAddExchangeRateOnTimeoutError() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()

	r := gin.Default()
	r.POST("/exchangeRate/add", cs.controller.AddExchangeRate)
	w := performRequest(
		r,
		"POST",
		"/exchangeRate/add?from=USD&to=IDR&date=2019-01-01T00:00:00.000Z&rate=0.1231",
		nil,
		ctx,
	)

	assert.Equal(cs.T(), 408, w.Code, "Code should be 408")
}

func (cs *ControllerSuite) TestAddExchangeRateOnDuplicateError() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	r := gin.Default()
	r.POST("/exchangeRate/add", cs.controller.AddExchangeRate)
	w := performRequest(
		r,
		"POST",
		"/exchangeRate/add?from=USD&to=IDR&date=2019-01-01T00:00:00.000Z&rate=0.1231",
		nil,
		ctx,
	)

	assert.Equal(cs.T(), 409, w.Code, "Code should be 409")
}

func (cs *ControllerSuite) TestAddExchangeRateOnDateError() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	r := gin.Default()
	r.GET("/exchangeRate/add", cs.controller.AddExchangeRate)
	w := performRequest(
		r,
		"GET",
		"/exchangeRate/add?from=USD&to=IDR&date=2019-01-0100:00:00.000&rate=0.1231",
		nil,
		ctx,
	)

	assert.Equal(cs.T(), 400, w.Code, "Code should be 400")
}

func (cs *ControllerSuite) TestAddExchangeRateOnRateError() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	r := gin.Default()
	r.GET("/exchangeRate/add", cs.controller.AddExchangeRate)
	w := performRequest(
		r,
		"GET",
		"/exchangeRate/add?from=USD&to=IDR&date=2019-01-01T00:00:00.000Z&rate=A31",
		nil,
		ctx,
	)

	assert.Equal(cs.T(), 400, w.Code, "Code should be 400")
}