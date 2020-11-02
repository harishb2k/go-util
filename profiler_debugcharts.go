package goxutil

import (
    "fmt"
    "github.com/gorilla/handlers"
    _ "github.com/gorilla/handlers"
    _ "github.com/mkevac/debugcharts"
    "log"
    "net/http"
    _ "net/http/pprof"
)

func LaunchProfileDebugChart(port int) {
    go func() {
        log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handlers.CompressHandler(http.DefaultServeMux)))
    }()
    log.Printf("ProfileDebugChart - Running at http://localhost:%d/debug/charts/", port)
}

