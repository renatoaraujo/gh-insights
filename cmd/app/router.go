package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"renatoaraujo/gh-insights/pkg/charts"
	"renatoaraujo/gh-insights/pkg/infrastructure"
)

type App struct {
	Router http.Handler
}

func NewApp(ctx context.Context, db *infrastructure.Database) App {
	return App{
		Router: buildRouter(ctx, db),
	}
}

func buildRouter(ctx context.Context, db *infrastructure.Database) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Welcome home")
	})

	router.HandleFunc("/lead-time", func(writer http.ResponseWriter, _ *http.Request) {
		charts.LeadTimeChartHandler(ctx, writer, db)
	}).Methods("GET")

	router.HandleFunc("/pulls-throughput", func(writer http.ResponseWriter, _ *http.Request) {
		charts.PullsThroughputChartHandler(ctx, writer, db)
	}).Methods("GET")

	return router
}
