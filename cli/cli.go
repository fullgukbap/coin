package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/fullgukbap/coin/explorer"
	"github.com/fullgukbap/coin/rest"
)

func usage() {
	fmt.Printf("Welcome to NomadCoin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("explorer:   Set the PORT of the server\n")
	fmt.Printf("rest:       Choose between 'html', and 'rest'\n\n")
	runtime.Goexit()
}

func Start() {

	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	model := flag.String("mode", "rest", "Choose between 'html', and 'rest'")

	flag.Parse()

	switch *model {
	case "rest":
		rest.Start(*port)
	case "port":
		explorer.Start(*port)
	default:
		usage()
	}
}
