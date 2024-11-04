package pkg_models

import (
	jtoken "github.com/golang-jwt/jwt/v4"
)

type Token struct {
	UserID     int    `json:"userId"`
	LoginID    string `json:"loginId"`
	Email      string `json:"email"`
	CurrencyID int    `json:"currencyId"`
	LanguageID int    `json:"languageId"`
	Timezone   string `json:"timezone"`
	SessionID  string `json:"sessionId"`
	Expiry     int64  `json:"exp"`
	jtoken.RegisteredClaims
}
