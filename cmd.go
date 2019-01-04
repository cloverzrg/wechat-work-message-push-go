package main

import "flag"

type cmdParams struct {
    isPrintVersion bool
}

func parseCmdParams() (params *cmdParams) {
    params = &cmdParams{}
    flag.BoolVar(&params.isPrintVersion, "v", false, "print version")
    flag.Parse()
    return params
}
