package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type List struct {
	Rates Rates `json:"rates"`
}
type Rates struct {
	EUR float64 `json:"EUR"`
	USD float64 `json:"RUB"`
}

func main() {
	id := os.Getenv("APP_ID")
	urlRequest := "https://openexchangerates.org/api/latest.json?app_id=" + id
	for {
		resp, err := http.Get(urlRequest)
		if err != nil {
			log.Printf("Error: %s", err)
		}

		rates := List{}
		if err = json.NewDecoder(resp.Body).Decode(&rates); err != nil {
			log.Printf("Error: %s", err)
		}
		if err = resp.Body.Close(); err != nil {
			log.Printf("Error: %s", err)
		}

		writeToFile(fmt.Sprintf("€=%.2f $=%.2f", rates.Rates.USD/rates.Rates.EUR, rates.Rates.USD))
		time.Sleep(1 * time.Hour)
	}
}

func writeToFile(str string) {
	file, err := os.Create("/data/currency.txt")
	if err != nil {
		log.Printf("Error: %s", err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Printf("Error: %s", err)
		}
	}()

	if _, err = file.WriteString(str); err != nil {
		log.Printf("Error: %s", err)
	}
}
