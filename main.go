package main

import (
    "fmt"
    "gopkg.in/telegram-bot-api.v4"
    "log"
    "os"
    "encoding/json"
)

type Config struct {
    TelegramBotToken string
}

func main() {
    file, _ := os.Open("config.json")
    decoder := json.NewDecoder(file)
    configuration := Config{}
    err := decoder.Decode(&configuration)
    if err != nil {
       log.Panic(err)
    }
    fmt.Println(configuration.TelegramBotToken)
}
