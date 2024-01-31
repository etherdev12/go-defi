package main

import (
	"context"
	"fmt"
	"log"

	"github.com/etherdev12/go-defi/bclient"
	"github.com/etherdev12/go-defi/config"
	"github.com/pkg/errors"
)

var (
	// Version denotes the compile time build version
	Version string
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(errors.Wrap(err, "load config"))
	}
	client, err := cfg.EthClient(ctx)
	if err != nil {
		log.Fatal(errors.Wrap(err, "eth client"))
	}
	bc, err := bclient.NewClient(ctx, client)
	if err != nil {
		log.Fatal(errors.Wrap(err, "new client"))
	}
	price, err := bc.EthDaiPrice()
	if err != nil {
		log.Fatal(errors.Wrap(err, "get price"))
	}
	fmt.Println(price)
}
