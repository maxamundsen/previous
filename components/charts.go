package components

type ChartOptions struct {
	Title string
}

type BarChartOptions struct {
	Base ChartOptions
	XLabel string
	YLabel string
}

func BarChart(opts BarChartOptions) Node {
	return nil
}