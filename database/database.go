package database

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"

	m "github.com/avcwisesa/foreign_currency/model"
)

type Database interface {
	Migrate(interface{})
	AddExchangeRate(m.ExchangeRate) (m.ExchangeRate, error)
	AddTrackedExchange(m.TrackedExchange) (m.TrackedExchange, error)
	GetExchangeRate(string, string, time.Time) (m.ExchangeRate, error)
	GetExchangeRateHist(string, string) ([]m.ExchangeRate, error)
	GetTrackedExchangeList(string) ([]m.TrackedExchange, error)
	DeleteTrackedExchange(m.TrackedExchange) ([]m.TrackedExchange, error)
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

func (d *database) GetExchangeRate(from string, to string, date time.Time) (m.ExchangeRate, error) {

	var exchangeRate m.ExchangeRate
	if err := d.client.Where(&m.ExchangeRate{
		From: from,
		To: to,
		Date: date,
	}).First(&exchangeRate).Error; err != nil {
		return m.ExchangeRate{}, err
	}

	return exchangeRate, nil
}

func (d *database) GetExchangeRateHist(from string, to string) ([]m.ExchangeRate, error) {

	var exchangeRateHist []m.ExchangeRate
	if err := d.client.Where(&m.ExchangeRate{
		From: from,
		To: to,
	}).Order("date desc").Limit(7).Find(&exchangeRateHist).Error; err != nil {
		return nil, err
	}

	return exchangeRateHist, nil
}

func (d *database) GetTrackedExchangeList(user string) ([]m.TrackedExchange, error) {

	var trackedExchangeList []m.TrackedExchange
	if err := d.client.Where(&m.TrackedExchange{
		User: user,
	}).Find(&trackedExchangeList).Error; err != nil {
		return nil, err
	}

	return trackedExchangeList, nil
}

func (d *database) DeleteTrackedExchange(trackedExchange m.TrackedExchange) ([]m.TrackedExchange, error) {

	if err := d.client.Delete(&trackedExchange).Error; err != nil {
		return nil, err
	}

	var trackedExchangeList []m.TrackedExchange
	if err := d.client.Where(&m.TrackedExchange{
		User: trackedExchange.User,
	}).Find(&trackedExchangeList).Error; err != nil {
		return nil, err
	}

	return trackedExchangeList, nil
}