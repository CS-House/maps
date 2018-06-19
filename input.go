package main

import (
	"crypto/sha512"
	"fmt"
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"
)

type jsonObject struct {
	AccountName  string   `json:"AccountName"`
	ClientTokens []string `json:"ClientTokens"`
	EmailAddress string   `json:"EmailAddress"`
	Password     string   `json:"Password"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (jo *jsonObject) fillJsonObject() {
	jo.AccountName = randSeq(10)
	append(jo.ClientTokens, uuid.Must(uuid.NewV4()))
	append(jo.ClientTokens, uuid.Must(uuid.NewV4()))
	append(jo.ClientTokens, uuid.Must(uuid.NewV4()))
	jo.EmailAddress = randSeq(10) + "@admin.com"
	jo.Password = sha512.New()
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	fmt.Println(randSeq(10))

	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)

	jsonObject := &jsonObject{}

}
