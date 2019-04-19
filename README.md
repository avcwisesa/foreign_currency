# Foreign Currency

This back-end APIs are to be used by front-end engineers to develop an application that store and display foreign exchange rate for currencies on a daily basis.

## Prerequisites & Setup

- Linux machine with root access
- Golang - [Go installation](https://golang.org/doc/install)

```
git clone https://github.com/avcwisesa/foreign_currency
```

## Running Test

For running test, use command below:
```
make test
```

## Running Program

For running program, use command below:
```
go run main.go
```

## Domain Model

### Daily Exchange Rate
| Attribute | Type | Description |
| --- | --- | --- |
| `from` | string | Currency code for origin currency |
| `to` | string | Currency code for target currency |
| `date` | datetime | Date on which exchange rate was recorded |
| `rate` | float | Exchange rate on specified date |

### User Tracked Exchange Rate
| Attribute | Type | Description |
| --- | --- | --- |
| `from` | string | Currency code for origin currency |
| `to` | string | Currency code for target currency |
| `user` | string | User who track the exchange rate |

## API Documentation

### User wants to input daily exchange rate data `POST /exchangeRate/add`

| Attribute | Type | Required |
| --- | --- | --- |
| `from` | string | yes |
| `to` | string | yes |
| `date` | datetime | yes |
| `rate` | float | yes |

datetime input are in RFC3339 format

### User has a list of exchange rates to be tracked `GET /trackedExchange`

| Attribute | Type | Required |
| --- | --- | --- |
| `user` | string | yes |

### User wants to see the exchange rate trend from the most recent 7 data points `GET /exchangeRate`

| Attribute | Type | Required |
| --- | --- | --- |
| `from` | string | yes |
| `to` | string | yes |

### User wants to add an exchange rate to the list `POST /trackedExchange/add`

| Attribute | Type | Required |
| --- | --- | --- |
| `from` | string | yes |
| `to` | string | yes |
| `user` | string | yes |

### User wants to remove an exchange rate from the list `DELETE /trackedExchange/delete`

| Attribute | Type | Required |
| --- | --- | --- |
| `from` | string | yes |
| `to` | string | yes |
| `user` | string | yes |

## Assumptions Used in Development
- all required input is always given
- frontend always give 00:00:00.000Z for time part
- dialy exchange rate can only be deleted directly from database

## TODO
- add DB model documentation
