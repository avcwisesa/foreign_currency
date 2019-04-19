package controller

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	d "github.com/avcwisesa/foreign_currency/database"
	m "github.com/avcwisesa/foreign_currency/model"
)

type Controller interface {
	AddExchangeRate(*gin.Context)
	AddTrackedExchange(*gin.Context)
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

	rate, err := strconv.ParseFloat(ctx.Query("rate"), 64)
	if err != nil {
		ctx.JSON(400, "Rate format must be float!")
		return
	}

	exchangeRate := m.ExchangeRate{
		From: ctx.Query("from"),
		To: ctx.Query("to"),
		Date: date,
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

