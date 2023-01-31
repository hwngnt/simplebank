package util

import {
	"math/rand"
	"strings"
	"time"
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max values.
func RandomInt(min, max int64) int64 {
	return min + rand.Int63(max - min + 1)
}

//RandomString generates a random string of leng n.
func RandomString (n int64) string {
	var sb string.Builder
	k := len(alphabet)

	for i:=0; i<n; i++{
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

//RandomOwner generates a random owner name 
func RandomOwner() string {
	return RandomString(6)
}

//RandomMoney generates a random amount of money
func RandomMoney() int64{
	return RandomInt(0, 1000)
}

//RandomCurrency generates a random currency code
func RandomCurrency() string{
	currencies := []string{"USD", "EUR", "CAD", "VND"}
	n := len(currencies)
	return currencies[RandomInt(0,n)]
}