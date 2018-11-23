package main

import "flag"

type cmdParams struct {
	configPath     string
	isPrintVersion bool
}

func parseCmdParams() (params *cmdParams) {
	params = &cmdParams{}
	flag.BoolVar(&params.isPrintVersion, "v", false, "print version")
	flag.StringVar(&params.configPath, "c", "config.json", "specify config file")
	flag.Parse()
	return params
}
