
package main

import (
	"github.com/benobro/go-ezmail"
	"math/rand"
	"time"
)

func main(){ 
	yourcode := RandomString(50)
	ezmail.SendMail("benobrien705@gmail.com", "nqjslvzgamidktur", "smtp.gmail.com", 587, "ben.obrien000@gmail.com", "test", "test ", yourcode)
}

func RandomString(strlen int) string {
        rand.Seed(time.Now().UTC().UnixNano())
        const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
        result := make([]byte, strlen)
        for i := 0; i < strlen; i++ {
                result[i] = chars[rand.Intn(len(chars))]
        }
        return string(result)
}
