package charts

import (
	"net/http"
	"reflect"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

type LineChart struct {
	line *charts.Line
}

type Series struct {
	Values interface{}
}

func NewLineChart() LineChart {
	return LineChart{
		line: charts.NewLine(),
	}
}

func (lc LineChart) SetTitle(title, sub string) {
	lc.line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    title,
			Subtitle: sub,
		}),
	)
}

func (lc LineChart) SetTooltip(trigger, triggerOn string) {
	lc.line.SetGlobalOptions(
		charts.WithTooltipOpts(opts.Tooltip{
			Show:      true,
			Trigger:   trigger,
			TriggerOn: triggerOn,
		}),
	)
}

func (lc LineChart) SetXAxis(x []string) {
	lc.line.SetXAxis(x)
}

func (lc LineChart) AddSeries(name string, series Series) {
	items := make([]opts.LineData, 0)

	switch reflect.TypeOf(series.Values).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(series.Values)
		for i := 0; i < s.Len(); i++ {
			items = append(items, opts.LineData{Value: s.Index(i).Interface(), Symbol: "circle"})
		}
	default:
	}

	lc.line.AddSeries(name, items)
}

func (lc LineChart) ExampleLineChart() {
	lc.line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
	)

	lc.line.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true, ConnectNulls: true}))
}

func (lc LineChart) Render(w http.ResponseWriter) {
	lc.line.Render(w)
}
