package util

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)
}

// generates random integer btw min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomFloat(min, max float64) float64 {
	randNum := rand.Float64()*max + min

	return randNum
}

// generates random string
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
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCategory() string {
	category := []string{"food", "car", "phone", "laptop"}
	n := len(category)
	return category[rand.Intn(n)]
}

func RandomPaymentMethod() string {
	paymentMethod := []string{"bank transfer", "card"}
	n := len(paymentMethod)
	return paymentMethod[rand.Intn(n)]
}

func NewNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func NewNullInt(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: i != 0}
}

func NewNullBool(i bool) sql.NullBool {
	return sql.NullBool{Bool: i, Valid: i != false}
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "NGN"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomUUIDR() (id *uuid.UUID, err error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &uuid, nil
}

func Test() uuid.UUID {
	id, _ := RandomUUIDR()

	return *id
}

func RandomRole() string {
	roles := []string{"merchant", "admin", "buyer"}
	n := len(roles)
	return roles[rand.Intn(n)]
}

func RandomCountry() string {
	roles := []string{"USA", "UK", "Canada", "Nigeria", "Germany"}
	n := len(roles)
	return roles[rand.Intn(n)]
}

func ConvertStringToUUID(id string) (uuid.UUID, error) {
	uuid, err := uuid.Parse(id)

	return uuid, err
}
