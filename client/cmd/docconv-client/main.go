package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/monzo/docconv/client"
)

var (
	path     = flag.String("path", "", "`path` to file to convert")
	endpoint = flag.String("endpoint", "", "docconv HTTP server to use for conversion")
)

func main() {
	flag.Parse()

	if *path == "" {
		fmt.Println("must set -path")
		os.Exit(1)
	}

	var opts []client.Opt
	if *endpoint != "" {
		opts = append(opts, client.WithEndpoint(*endpoint))
	}

	c := client.New(opts...)
	resp, err := client.ConvertPath(c, *path)
	if err != nil {
		fmt.Printf("Could not convert: %v", err)
		os.Exit(1)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(resp); err != nil {
		fmt.Printf("Could not encode response: %v", err)
		os.Exit(1)
	}
}
