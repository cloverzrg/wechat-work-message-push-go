package config

import "flag"

type cmdParams struct {
	IsPrintVersion bool
}

func ParseCmdParams() (params *cmdParams) {
	params = &cmdParams{}
	flag.BoolVar(&params.IsPrintVersion, "v", false, "print version")
	flag.Parse()
	return params
}
