// (C) 2014 Mathias Dalheimer <md@gonium.net>. Code derived from the
// Gorilla WebSocket Demo, which is licensed as follows:
// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"text/template"
)

var configFile = flag.String("config", "defluxio.conf", "configuration file")
var homeTempl = template.Must(template.ParseFiles("home.html"))

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTempl.Execute(w, r.Host)
}

func init() {
	flag.Parse()
	loadConfiguration(*configFile)
}

func main() {
	go h.run()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/api/submit/{meter}", submitReading).Methods("POST")
	r.HandleFunc("/ws", serveWs)
	http.Handle("/", r)
	err := http.ListenAndServe(Cfg.Network.Host, nil)
	if err != nil {
		log.Fatal("Failed to start http server: ", err)
	}
}
