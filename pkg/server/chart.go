package server

import (
	"net/http"
	"renatoaraujo/gh-insights/pkg/charts"
)

func chartLink(w http.ResponseWriter, _ *http.Request) {
	lc := charts.NewLineChart()
	lc.ExampleLineChart()
	lc.Render(w)
}
