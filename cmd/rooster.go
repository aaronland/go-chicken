package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/thisisaaronland/go-chicken"
	"github.com/whosonfirst/go-sanitize"
	"io/ioutil"
	"net/http"	
	"os"
)

func main() {

	var host = flag.String("host", "localhost", "The hostname to listen for requests on")
	var port = flag.Int("port", 8080, "The port number to listen for requests on")

	handler := func(rsp http.ResponseWriter, req *http.Request) {

		query := req.URL.Query()

		lang := query.Get("language")
		clucking := query.Get("clucking")

		opts := sanitize.DefaultOptions()
		output, _ := sanitize.SanitizeString(input, opts)

		ch, err := chicken.GetChickenForLanguageTag(lang, clucking)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}
		
		body, err := ioutil.ReadAll(req.Body)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		b := ch.TextToChicken(string(txt))

		rsp.Header().Set("Access-Control-Allow-Origin", "*")
		rsp.Header().Set("Content-Type", "text/chicken")
		
		rsp.Write(b)		
	}
	
	endpoint := fmt.Sprintf("%s:%d", *host, *port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	err := gracehttp.Serve(&http.Server{Addr: endpoint, Handler: mux})

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
