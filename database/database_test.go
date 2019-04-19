package database_test

import (
	"os"
	"time"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	d "github.com/avcwisesa/foreign_currency/database"
	m "github.com/avcwisesa/foreign_currency/model"
)

type DatabaseSuite struct {
	suite.Suite
	database d.Database
}

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, &DatabaseSuite{})
}

func (ds *DatabaseSuite) SetupSuite() {
	client, _ := gorm.Open("sqlite3", "test.sqlite")

	ds.database = d.New(client)

	ds.database.Migrate(&m.ExchangeRate{})
	ds.database.Migrate(&m.TrackedExchange{})

	exchangeRate := m.ExchangeRate{
		From: "IDR",
		To: "USD",
		Date: time.Date(2019, time.April, 1, 0, 0, 0, 0, time.UTC),
		Rate: 0.123,
	}

	ds.database.AddExchangeRate(exchangeRate)
}

func (ds *DatabaseSuite) TearDownSuite() {
	os.Remove("test.sqlite")
}

func (ds *DatabaseSuite) TestAddExchangeRateOnSuccess() {
	exchangeRate := m.ExchangeRate{
		From: "IDR",
		To: "USD",
		Date: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		Rate: 0.123,
	}

	_, err := ds.database.AddExchangeRate(exchangeRate)

	assert.Nil(ds.T(), err, "Error should be nil!")
}

func (ds *DatabaseSuite) TestAddExchangeRateOnDuplicateError() {
	exchangeRate := m.ExchangeRate{
		From: "IDR",
		To: "USD",
		Date: time.Date(2019, time.April, 1, 0, 0, 0, 0, time.UTC),
		Rate: 0.123,
	}

	_, err := ds.database.AddExchangeRate(exchangeRate)

	assert.NotNil(ds.T(), err, "Error should not be nil!")
}