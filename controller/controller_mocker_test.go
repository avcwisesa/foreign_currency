package controller_test

import (
	"errors"
	"time"

	m "github.com/avcwisesa/foreign_currency/model"
)

type DBMock struct{}

func (dm *DBMock) Migrate(i interface{}) {
	return
}

func (dm *DBMock) AddExchangeRate(exchangeRate m.ExchangeRate) (m.ExchangeRate, error){
	if exchangeRate.From == "USD" {
		return m.ExchangeRate{}, errors.New("Error while Querying from database!")
	}

	return exchangeRate, nil
}

func (dm *DBMock) AddTrackedExchange(trackedExchange m.TrackedExchange) (m.TrackedExchange, error){
	if trackedExchange.User == "test1" {
		return m.TrackedExchange{}, errors.New("Error while Querying from database!")
	}

	return trackedExchange, nil
}

func (dm *DBMock) GetExchangeRate(from string, to string, date time.Time) (m.ExchangeRate, error) {
	return m.ExchangeRate{}, nil
}

func (dm *DBMock) GetTrackedExchangeList(user string) ([]m.TrackedExchange, error) {
	return nil, nil
}

func (dm *DBMock) DeleteTrackedExchange(from string, to string, user string) ([]m.TrackedExchange, error) {
	return nil, nil
}