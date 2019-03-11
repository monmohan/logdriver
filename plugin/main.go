package main

import (
	"fmt"
	"os"

	"github.com/docker/go-plugins-helpers/sdk"
)

func main() {
	cfgVal := os.Getenv("SOME_CFG_FOR_PLUGIN")
	fmt.Fprintln(os.Stdout, "A config value ", cfgVal)
	h := sdk.NewHandler(`{"Implements": ["LoggingDriver"]}`)
	handlers(&h, newDriver())
	if err := h.ServeUnix("mylogdriver", 0); err != nil {
		panic(err)
	}
}
