package main

import (
	"context"
	"flag"
	"net/http"
	"encoding/json"
	"fmt"
)

func main() {

	var kubeconfig string
	var port int

	ctx := context.TODO()
	isLeader := false

	flag.StringVar(&kubeconfig, "kubeconfig", "", "The location of the kubeconfig file")
	flag.IntVar(&port, "port", 8080, "The port for the HTTP server")
	flag.Parse()

	go startElectionProcess(ctx, kubeconfig, func(elected bool) {
		isLeader = elected
	})

	http.HandleFunc("/isLeader", func(res http.ResponseWriter, req *http.Request) {
		response, _ := json.Marshal(Response{
			IsLeader: isLeader,
		})

		res.Header().Add("Content-Type", "application/json")
		res.Write(response)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}

type Response struct {
	IsLeader bool `json:"isLeader"`
}
