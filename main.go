package main

import (
	"context"
	"flag"
	"github.com/radiusxyz/lighthouse-bidder/config"
	"github.com/radiusxyz/lighthouse-bidder/manager"
	"strings"
)

func ParseFlag() {
	flag.String("bidder.address", "", "Bidder Address")
	flag.String("rollup.ids", "", "Rollup IDs")
	flag.Parse()
}

func GetFlag(paramName string) string {
	return flag.Lookup(paramName).Value.(flag.Getter).Get().(string)
}

func main() {
	ParseFlag()
	bidderAddress := GetFlag("bidder.address")
	rawRollupIds := GetFlag("rollup.ids")
	rollupIds := strings.Fields(rawRollupIds)

	conf := config.New()
	m, err := manager.New(conf, bidderAddress, rollupIds)
	if err != nil {
		panic(err)
	}

	globalCtx := context.Background()
	m.Start(globalCtx)
	select {}
}
