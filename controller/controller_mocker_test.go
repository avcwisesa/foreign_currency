package controller_test

import (
	"errors"

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