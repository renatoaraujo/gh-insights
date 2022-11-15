package charts

import (
	"context"
	"log"
	"math"
	"net/http"
	"reflect"
	"time"

	"renatoaraujo/gh-insights/pkg/infrastructure"
)

func LeadTimeChartHandler(ctx context.Context, w http.ResponseWriter, db *infrastructure.Database) {
	lc := NewLineChart()
	lc.SetTitle("Lead time", "Time to close issue from the time it's opened")
	lc.SetTooltip("item", "mousemove")

	t := time.Now()
	var avgPerMonth []float64
	var head []string

	for i := 0; i < 12; i++ {
		previousMonth := t.AddDate(0, -i, 0)
		issues, err := db.GetIssuesClosedByMonthAndYear(ctx, int(previousMonth.Month()), previousMonth.Year())
		if err != nil {
			log.Fatal(err)
		}

		var totalHours float64
		for _, issue := range issues {
			totalHours = totalHours + issue.TimeOpenedMinutes
		}

		monthAvg := totalHours / float64(len(issues))
		if math.IsNaN(monthAvg) {
			monthAvg = float64(0)
		}

		avgPerMonth = append(avgPerMonth, monthAvg)
		head = append(head, previousMonth.Month().String())
	}

	reverse(head)
	reverse(avgPerMonth)

	lc.SetXAxis(head)
	series := Series{Values: avgPerMonth}
	lc.AddSeries("Lead time per month", series)
	lc.ExampleLineChart()
	lc.Render(w)
}

func PullsThroughputChartHandler(ctx context.Context, w http.ResponseWriter, db *infrastructure.Database) {
	lc := NewLineChart()
	lc.SetTitle("Pull Requests Throughput", "Number of Pull Requests open and Pull Requests closed over time")
	lc.SetTooltip("item", "mousemove")

	t := time.Now()
	var opened []int
	var closed []int
	var head []string

	for i := 0; i < 18; i++ {
		previousMonth := t.AddDate(0, -i, 0)
		pullsOpened, err := db.GetOpenedPullsByMonthAndYear(ctx, int(previousMonth.Month()), previousMonth.Year())
		if err != nil {
			log.Fatal(err)
		}

		pullsClosed, err := db.GetClosedPullsByMonthAndYear(ctx, int(previousMonth.Month()), previousMonth.Year())
		if err != nil {
			log.Fatal(err)
		}

		closed = append(closed, len(pullsClosed))
		opened = append(opened, len(pullsOpened))
		head = append(head, previousMonth.Month().String())
	}

	reverse(head)
	reverse(opened)
	reverse(closed)

	lc.SetXAxis(head)
	seriesOpened := Series{Values: opened}
	lc.AddSeries("Pull Requests opened", seriesOpened)
	seriesClosed := Series{Values: closed}
	lc.AddSeries("Pull Requests closed", seriesClosed)
	lc.ExampleLineChart()
	lc.Render(w)
}

func reverse(input interface{}) {
	inputLen := reflect.ValueOf(input).Len()
	inputMid := inputLen / 2
	inputSwap := reflect.Swapper(input)

	for i := 0; i < inputMid; i++ {
		j := inputLen - i - 1

		inputSwap(i, j)
	}
}
