package config

import "github.com/rahmaninsani/dns-api/config/subcommand"

type config struct {
	Subcommand subcommand.Subcommand `json:"subcommand"`
}

var Config *config

func SetupConfig() {
	var sc subcommand.Subcommand
	sc = *sc.Parse()

	Config = &config{
		Subcommand: sc,
	}
}
