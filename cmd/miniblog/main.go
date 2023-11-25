package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/spf13/pflag"
)

const helpText = `Usage: main [flags] arg [arg...]
This is a very simple app framework(do nothing).
Flags:`

var (
	addr = pflag.String("addr", ":8088", "The address to listen to.")
	help = pflag.BoolP("help", "h", false, "show help message")

	usage = func() {
		fmt.Println(helpText)
		flag.PrintDefaults()
	}
)

func main() {
	pflag.Usage = usage
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello world")
	})

	server := http.Server{Addr: *addr}
	fmt.Printf(`starting http server at %s`, *addr)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
