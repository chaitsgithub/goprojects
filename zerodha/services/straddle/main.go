package main

import (
    "fmt"
    "log"
    "math/rand"
    "time"

    "github.com/zerodha/gokiteconnect/v4"
)

// Replace with your Zerodha API access token
var apiKey = "YOUR_API_KEY"

func main() {
    // Create a new Kite Connect instance
    kite := kiteconnect.New(apiKey)

    // Get current timestamp (adjusted for 9:20 AM)
    now := time.Now()
    targetTime := time.Date(now.Year(), now.Month(), now.Day(), 9, 20, 0, 0, now.Location())

    // Check if it's past 9:20 AM, exit if so
    if now.After(targetTime) {
        log.Println("Market already past 9:20 AM, exiting")
        return
    }

    // Wait until 9:20 AM
    time.Sleep(time.Until(targetTime))

    // Get Bank Nifty options
    instruments, err := getBankNiftyOptions(kite)
    if err != nil {
        log.Fatal(err)
    }

    // Place orders
    for _, instrument := range instruments {
        err := placeStraddleOrder(kite, instrument, 0.35)
        if err != nil {
            log.Printf("Error placing straddle order for %s: %v\n", instrument.Tradingsymbol, err)
        }
    }
}

func getBankNiftyOptions(kite *kiteconnect.KiteConnect) ([]kiteconnect.Instrument, error) {
    // Get options for current expiry
    expiry := time.Now().Add(time.Hour * 24 * 7).Format("yyyy-mm-dd") // Adjust for expiry as needed

    // Fetch instruments using kite.Instruments
    instruments, err := kite.Instruments(kiteconnect.InstrumentTypeOptions, kiteconnect.ExchangeNSE)
    if err != nil {
        return nil, err
    }

    // Filter for Bank Nifty options expiring on the target expiry
    var niftyOptions []kiteconnect.Instrument
    for _, instrument := range instruments {
        if instrument.Expiry == expiry && instrument.Tradingsymbol == "NIFTY 50" {
            niftyOptions = append(niftyOptions, instrument)
        }
    }

    return niftyOptions, nil
}

func placeStraddleOrder(kite *kiteconnect.KiteConnect, instrument kiteconnect.Instrument, stopLoss float64) error {
    // Fetch ATM strike price for Call and Put options
    strike, err := getATMStrike(kite, instrument.Exchange, instrument.Tradingsymbol, instrument.Expiry)
    if err != nil {
        return err
    }

    // Randomly choose between Call or Put leg to sell first (optional for order diversification)
    if rand.Intn(2) == 0 {
        err = placeOrder(kite, instrument.Exchange, instrument.Tradingsymbol, strike, kiteconnect.TransactionTypeSell, kiteconnect.ProductMis, stopLoss)
        if err != nil {
            return err
        }
    } else {
        err = placeOrder(kite, instrument.Exchange, instrument.Tradingsymbol, strike, kiteconnect.TransactionTypeSell, kiteconnect.ProductMis, stopLoss)
        if err != nil {
            return err
        }
    }

    // Place opposite leg order (Call if Put was sold first, vice versa)
    var product kiteconnect.Product
    if strike.Strike == instrument.Strike {
        product = kiteconnect.ProductCNC
    } else {
        product = kiteconnect.ProductMIS
    }
    err = placeOrder(kite, instrument.Exchange, instrument.Tradingsymbol, strike, kiteconnect.TransactionTypeSell, product, stopLoss)
    if err != nil {
        return err
    }

    return nil
}

func placeOrder(kite *kiteconnect.KiteConnect, exchange, symbol string, strike float64, transactionType kiteconnect.TransactionType, product kiteconnect.Product, stopLoss float64) error {
    // Calculate stop-loss trigger price based on option price and stop-loss percentage
    stopLossTrigger := strike * (1 - stopLoss)

    // Create separate orders for main leg and stop-loss
    order := &kiteconnect.Order{
        Exchange:        exchange,
