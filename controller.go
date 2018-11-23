package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Index(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	ua := req.UserAgent()
	fmt.Fprint(res, "Hello World!"+"\n")
	fmt.Fprint(res, id)
	fmt.Fprint(res, ua)
}
