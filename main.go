package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ausaf007/uniswap-tracker/database"
	"github.com/ausaf007/uniswap-tracker/handlers"
	"github.com/ausaf007/uniswap-tracker/services"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"time"
)

// loadFlags is a helper function that sets the verbosity level of the program through user specified flags
func loadFlags() {
	isVerbose := flag.Bool("verbose", false, "Specifies verbosity of logs. True means Info Level. "+
		"False means Warn Level")
	flag.Parse()
	if *isVerbose {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}

type Config struct {
	ServerConfig  ServerConfig  `json:"server_config"`
	TrackerConfig TrackerConfig `json:"tracker_config"`
}

type ServerConfig struct {
	DatabaseName string `json:"database_name"`
	Port         string `json:"port"`
}

type TrackerConfig struct {
	EthClientURL  string `json:"eth_client_url"`
	PoolAddress   string `json:"pool_address"`
	PauseDuration int    `json:"pause_duration"`
}

func loadConfig(filename string) (Config, error) {
	var config Config

	configFile, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// gracefulShutdown is a helper function that stops the program
// When user presses Ctrl+C, the connection is closed and the program closes
func gracefulShutdown(client *ethclient.Client) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	go func() {
		<-sigs
		client.Close()
		fmt.Println("Uniswap Tracker stopped successfully.")
		os.Exit(0)
	}()
}

// TODO: Remove log.Fatal, implement retries for ethclient and database
func main() {
	loadFlags()

	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Initialize Database
	db, err := database.InitDatabase(config.ServerConfig.DatabaseName)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Create an ethclient.Client
	ethClient, err := ethclient.Dial(config.TrackerConfig.EthClientURL)
	if err != nil {
		log.Fatal("Failed to connect to Ethereum client:", err)
	}

	gracefulShutdown(ethClient)

	app := fiber.New()
	service, err := services.NewTrackingService(ethClient, db)
	if err != nil {
		log.Error("Failed to start NewTrackingService", err)
	}

	handler := handlers.NewPoolHandler(service)

	// Begin tracking the specified pool in a separate goroutine
	go func() {
		for {
			err := service.Tracker(config.TrackerConfig.PoolAddress)
			if err != nil {
				log.Error("Error encountered in Tracking:", err)
			}
			time.Sleep(time.Duration(config.TrackerConfig.PauseDuration) * time.Millisecond)
		}
	}()

	app.Get("/v1/api/pool/:pool_id", handler.PoolDataHandler)
	app.Get("/v1/api/pool/:pool_id/historic", handler.HistoricPoolDataHandler)

	app.Listen(":" + config.ServerConfig.Port)
}
