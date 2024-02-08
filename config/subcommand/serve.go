package subcommand

import (
	"flag"
	"github.com/rahmaninsani/dns-api/helper"
	"os"
)

type Serve struct {
	Address string `json:"address"`
	File    string `json:"file"`
}

func (s *Serve) GetFlag() *Serve {
	serve := flag.NewFlagSet("serve", flag.ExitOnError)
	address := serve.String("address", "localhost:8080", "The address where the API server listens (optional)")
	file := serve.String("file", "./data.json", "The path to the file storing the data (optional)")

	err := serve.Parse(os.Args[2:])
	helper.PanicIfError(err)

	return &Serve{
		Address: *address,
		File:    *file,
	}
}
