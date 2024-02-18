package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	// return capitalize first letter of each word
	str := RandomString(6) + " " + RandomString(6)
	return cases.Title(language.English).String(str)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
