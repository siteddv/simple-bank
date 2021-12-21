package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet       = "abcdefghijklmnopqrstuvwxyz"
	alphabetLength = len(alphabet)

	ownerNameLength = 6

	minMoneyAmount = 0
	maxMoneyAmount = 1000
)

var (
	currencies      = []string{"EUR", "USD", "CAD"}
	currenciesCount = len(currencies)
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	res := min + rand.Int63n(max-min+1)
	return res
}

// RandomString generates a random string of specified length
func RandomString(length int) string {
	var sb strings.Builder

	for i := 0; i < length; i++ {
		randCharIndex := rand.Intn(alphabetLength)
		ch := alphabet[randCharIndex]
		sb.WriteByte(ch)
	}

	res := sb.String()
	return res
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	res := RandomString(ownerNameLength)
	return res
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	res := RandomInt(minMoneyAmount, maxMoneyAmount)
	return res
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	randCurrencyIndex := rand.Intn(currenciesCount)

	res := currencies[randCurrencyIndex]
	return res
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
