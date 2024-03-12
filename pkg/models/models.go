package models

import "time"

type Users struct {
	UserId   int    `json:"user_Id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CurrencyRates struct {
	RatesId           int       `json:"rates_Id"`
	CurrencyName      string    `json:"currency_name"`
	CurrencyShortName string    `json:"currency_short_name"`
	Rate              float64   `json:"rate"`
	LastUpdated       time.Time `json:"last_updated"`
}

type Transactions struct {
	Id             int       `json:"id"`
	UserId         int       `json:"user_id"`
	CurrencyFromId int       `json:"currency_from_id"`
	CurrencyToId   int       `json:"currency_to_id"`
	CurrencyRate   float64   `json:"currency_rate"`
	AmountFrom     float64   `json:"amount_from"`
	AmountTo       float64   `json:"amount_to"`
	ConversionTime time.Time `json:"conversion_time"`
}
