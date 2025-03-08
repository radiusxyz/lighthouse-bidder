package main

import (
	"context"
	"flag"
	"github.com/radiusxyz/lighthouse-bidder/config"
	"github.com/radiusxyz/lighthouse-bidder/manager"
)

func ParseFlag() {
	flag.String("bidder.address", "", "Bidder Address")
	flag.String("rollup.id", "", "Rollup ID")
	flag.Parse()
}

func GetFlag(paramName string) string {
	return flag.Lookup(paramName).Value.(flag.Getter).Get().(string)
}

func main() {
	ParseFlag()
	bidderAddress := GetFlag("bidder.address")
	rollupId := GetFlag("rollup.id")

	conf := config.New()
	m, err := manager.New(conf, bidderAddress, rollupId)
	if err != nil {
		panic(err)
	}

	globalCtx := context.Background()
	m.Start(globalCtx)
	select {}
}
