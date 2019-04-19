package database

import (
	"errors"

	"github.com/jinzhu/gorm"

	m "github.com/avcwisesa/foreign_currency/model"
)

type Database interface {
	Migrate(interface{})
	AddExchangeRate(m.ExchangeRate) (m.ExchangeRate, error)
	AddTrackedExchange(m.TrackedExchange) (m.TrackedExchange, error)
}

type database struct {
	client *gorm.DB
}

func New(client *gorm.DB) Database {
	return &database{client: client}
}

func (d *database) Migrate(i interface{}) {
	d.client.AutoMigrate(i)
}

func (d *database) AddExchangeRate(exchangeRate m.ExchangeRate) (m.ExchangeRate, error) {

	if err := d.client.Where(&m.ExchangeRate{
		From: exchangeRate.From,
		To: exchangeRate.To,
		Date: exchangeRate.Date,
	}).First(&m.ExchangeRate{}).Error; err == nil {
		return m.ExchangeRate{}, errors.New("Overlapping data exists!")
	}

	d.client.Create(&exchangeRate)

	d.client.Where(&m.ExchangeRate{
		From: exchangeRate.From,
		To: exchangeRate.To,
		Date: exchangeRate.Date,
	}).First(&exchangeRate)

	return exchangeRate, nil
}

func (d *database) AddTrackedExchange(trackedExchange m.TrackedExchange) (m.TrackedExchange, error) {

	if err := d.client.Where(&m.TrackedExchange{
		From: trackedExchange.From,
		To: trackedExchange.To,
		User: trackedExchange.User,
	}).First(&m.TrackedExchange{}).Error; err == nil {
		return m.TrackedExchange{}, errors.New("Overlapping data exists!")
	}

	d.client.Create(&trackedExchange)

	d.client.Where(&m.TrackedExchange{
		From: trackedExchange.From,
		To: trackedExchange.To,
		User: trackedExchange.User,
	}).First(&trackedExchange)

	return trackedExchange, nil
}