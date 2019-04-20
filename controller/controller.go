package controller

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	d "github.com/avcwisesa/foreign_currency/database"
	m "github.com/avcwisesa/foreign_currency/model"
)

type Controller interface {
	GetExchangeRateHist(*gin.Context)
	GetTrackedExchanges(*gin.Context)
	AddExchangeRate(*gin.Context)
	AddTrackedExchange(*gin.Context)
	DeleteTrackedExchange(*gin.Context)
}

type controller struct {
	Database d.Database
}

func New(database d.Database) Controller {
	return &controller{
		Database: database,
	}
}

func (c *controller) AddExchangeRate(ctx *gin.Context) {
	select {
	case <-ctx.Request.Context().Done():
		ctx.JSON(408, nil)
		return
	default:
	}

	date, err:= time.Parse(time.RFC3339, ctx.Query("date"))
	if err != nil {
		ctx.JSON(400, "Use RFC3339 date format!")
		return
	}
	year, month, day := date.Date()
	dateNormalized := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	rate, err := strconv.ParseFloat(ctx.Query("rate"), 64)
	if err != nil {
		ctx.JSON(400, "Rate format must be float!")
		return
	}

	exchangeRate := m.ExchangeRate{
		From: ctx.Query("from"),
		To: ctx.Query("to"),
		Date: dateNormalized,
		Rate: rate,
	}

	exchangeRate, err = c.Database.AddExchangeRate(exchangeRate)
	if err != nil {
		ctx.JSON(409, nil)
		return
	}

	ctx.JSON(200, exchangeRate)
	return
}

func (c *controller) AddTrackedExchange(ctx *gin.Context) {
	select {
	case <-ctx.Request.Context().Done():
		ctx.JSON(408, nil)
		return
	default:
	}

	trackedExchange := m.TrackedExchange{
		From: ctx.Query("from"),
		To: ctx.Query("to"),
		User: ctx.Query("user"),
	}

	trackedExchange, err := c.Database.AddTrackedExchange(trackedExchange)
	if err != nil {
		ctx.JSON(409, nil)
		return
	}

	ctx.JSON(200, trackedExchange)
	return
}

func (c *controller) GetExchangeRateHist(ctx *gin.Context) {
	select {
	case <-ctx.Request.Context().Done():
		ctx.JSON(408, nil)
		return
	default:
	}

	from := ctx.Query("from")
	to := ctx.Query("to")

	exchangeRateHist, err := c.Database.GetExchangeRateHist(from, to)
	if err != nil {
		ctx.JSON(204, "Insufficient data")
	}

	ctx.JSON(200, exchangeRateHist)
	return
}

func (c *controller) GetTrackedExchanges(ctx *gin.Context) {
	ctx.JSON(404, nil)
	return
}

func (c *controller) DeleteTrackedExchange(ctx *gin.Context) {
	ctx.JSON(404, nil)
	return
}
