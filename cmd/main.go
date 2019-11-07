package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/knative-sample/serving-sdk-demo/cmd/app"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	// Start runner
	cmd := app.NewCommandStartServer()
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	flag.CommandLine.Parse([]string{})

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
