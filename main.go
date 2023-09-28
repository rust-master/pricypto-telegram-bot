package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	sendTopCoinsMessage(bot)

	// Set up a timer to send the message every 1 minute
	ticker := time.NewTicker(10 * time.Second)

	// Use a goroutine to repeatedly send the message
	go func() {
		for range ticker.C {
			sendTopCoinsMessage(bot)
		}
	}()

	// Block the main goroutine to keep the program running
	select {}
}

func sendTopCoinsMessage(bot *tgbotapi.BotAPI) {

	urlEth := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=ethereum"
	urlBtc := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=bitcoin"
	urlSol := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=solana"
	urlAda := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=cardano"
	urlDot := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=polkadot"
	urlXlm := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=stellar"
	urlAtom := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=cosmos"

	priceBtc, erre := fetchPriceData(urlBtc)
	if erre != nil {
		log.Printf("Error price fetching: %v", erre)
	}

	priceEth, erre := fetchPriceData(urlEth)
	if erre != nil {
		log.Printf("Error price fetching: %v", erre)
	}

	priceSol, erre := fetchPriceData(urlSol)
	if erre != nil {
		log.Printf("Error price fetching: %v", erre)
	}

	priceAda, erre := fetchPriceData(urlAda)
	if erre != nil {
		log.Printf("Error price fetching: %v", erre)
	}

	priceDot, erre := fetchPriceData(urlDot)
	if erre != nil {
		log.Printf("Error price fetching: %v", erre)
	}

	priceXlm, erre := fetchPriceData(urlXlm)
	if erre != nil {
		log.Printf("Error price fetching: %v", erre)
	}

	priceAtom, erre := fetchPriceData(urlAtom)
	if erre != nil {
		log.Printf("Error price fetching: %v", erre)
	}

	// Prices with emojis and bold formatting
	btcPrice := fmt.Sprintf("💰 *BTC*: $ %.2f", priceBtc[0].CurrentPrice)
	ethPrice := fmt.Sprintf("💰 *ETH*: $ %.2f", priceEth[0].CurrentPrice)
	solPrice := fmt.Sprintf("💰 *SOL*: $ %.2f", priceSol[0].CurrentPrice)
	adaPrice := fmt.Sprintf("💰 *ADA*: $ %.2f", priceAda[0].CurrentPrice)
	dotPrice := fmt.Sprintf("💰 *DOT*: $ %.2f", priceDot[0].CurrentPrice)
	xlmPrice := fmt.Sprintf("💰 *XLM*: $ %.2f", priceXlm[0].CurrentPrice)
	atomPrice := fmt.Sprintf("💰 *ATOM*: $ %.2f", priceAtom[0].CurrentPrice)

	// Construct the message with Markdown formatting
	messageFormat := fmt.Sprintf(`
    *Top Coins Prices* 📈

   %s
   %s
   %s
   %s
   %s
   %s
   %s

   🚀 Stay tuned for more updates! 🌕
   `, btcPrice, ethPrice, solPrice, adaPrice, dotPrice, xlmPrice, atomPrice)

	msg := tgbotapi.NewMessageToChannel("@top_coins_price_alerts", messageFormat)
	msg.ParseMode = "Markdown"

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
