package line

import (
	"github.com/ngyewch/go-chartjs"
	"go.octolab.org/pointer"
	"time"
)

func Example1() (*chartjs.LineChartConfiguration, error) {
	data := []chartjs.Point{
		{
			X: 0,
			Y: pointer.ToFloat64(0),
		},
		{
			X: 1,
			Y: pointer.ToFloat64(20),
		},
		{
			X: 2,
			Y: pointer.ToFloat64(20),
		},
		{
			X: 3,
			Y: pointer.ToFloat64(60),
		},
		{
			X: 4,
			Y: pointer.ToFloat64(60),
		},
		{
			X: 5,
			Y: pointer.ToFloat64(120),
		},
		{
			X: 6,
			Y: nil,
		},
		{
			X: 7,
			Y: pointer.ToFloat64(180),
		},
		{
			X: 8,
			Y: pointer.ToFloat64(120),
		},
		{
			X: 9,
			Y: pointer.ToFloat64(125),
		},
		{
			X: 10,
			Y: pointer.ToFloat64(105),
		},
		{
			X: 11,
			Y: pointer.ToFloat64(110),
		},
		{
			X: 12,
			Y: pointer.ToFloat64(170),
		},
	}
	return &chartjs.LineChartConfiguration{
		Data: &chartjs.LineChartData{
			Datasets: []*chartjs.LineChartDataset{
				{
					Data: data,
					ControllerDatasetOptions: &chartjs.ControllerDatasetOptions{
						Label: "Cubic interpolation (monotone)",
					},
					LineOptions: &chartjs.LineOptions{
						CommonElementOptions: &chartjs.CommonElementOptions{
							BorderColor: "red",
						},
						CubicInterpolationMode: chartjs.CubicInterpolationModeMonotone,
						Tension:                pointer.ToFloat64(0.4),
					},
				},
				{
					Data: data,
					ControllerDatasetOptions: &chartjs.ControllerDatasetOptions{
						Label: "Cubic interpolation",
					},
					LineOptions: &chartjs.LineOptions{
						CommonElementOptions: &chartjs.CommonElementOptions{
							BorderColor: "blue",
						},
						Tension: pointer.ToFloat64(0.4),
					},
				},
				{
					Data: data,
					ControllerDatasetOptions: &chartjs.ControllerDatasetOptions{
						Label: "Linear interpolation",
					},
					LineOptions: &chartjs.LineOptions{
						CommonElementOptions: &chartjs.CommonElementOptions{
							BorderColor: "green",
						},
					},
				},
			},
		},
		Options: &chartjs.LineControllerChartOptions{
			CoreChartOptions: &chartjs.CoreChartOptions{
				MaintainAspectRatio: pointer.ToBool(false),
			},
			Scales: map[string]chartjs.ICartesianScaleType{
				"x": &chartjs.LinearScaleOptions{
					CartesianScaleOptions: &chartjs.CartesianScaleOptions{
						Title: &chartjs.CartesianScaleTitle{
							Display: pointer.ToBool(true),
						},
					},
				},
				"y": &chartjs.LinearScaleOptions{
					CartesianScaleOptions: &chartjs.CartesianScaleOptions{
						Title: &chartjs.CartesianScaleTitle{
							Display: pointer.ToBool(true),
							Text:    "Value",
						},
					},
					BeginAtZero: pointer.ToBool(true),
				},
			},
			Plugins: &chartjs.PluginOptions{
				Title: &chartjs.TitleOptions{
					Display: pointer.ToBool(true),
					Text:    chartjs.TitleText("Chart.js Line Chart - Cubic interpolation mode"),
				},
			},
		},
	}, nil
}

func Example2() (*chartjs.LineChartConfiguration, error) {
	return &chartjs.LineChartConfiguration{
		Data: &chartjs.LineChartData{
			Datasets: []*chartjs.LineChartDataset{
				{
					ControllerDatasetOptions: &chartjs.ControllerDatasetOptions{
						Label: "Series 1",
					},
					Data: []chartjs.Point{
						{
							X: float64(time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local).UnixMilli()),
							Y: pointer.ToFloat64(10),
						},
						{
							X: float64(time.Date(2024, 1, 1, 1, 0, 0, 0, time.Local).UnixMilli()),
							Y: pointer.ToFloat64(23.7),
						},
						{
							X: float64(time.Date(2024, 1, 1, 2, 0, 0, 0, time.Local).UnixMilli()),
							Y: pointer.ToFloat64(12.4),
						},
						{
							X: float64(time.Date(2024, 1, 1, 3, 0, 0, 0, time.Local).UnixMilli()),
							Y: pointer.ToFloat64(5.1),
						},
					},
				},
			},
		},
		Options: &chartjs.LineControllerChartOptions{
			CoreChartOptions: &chartjs.CoreChartOptions{
				MaintainAspectRatio: pointer.ToBool(false),
			},
			Scales: map[string]chartjs.ICartesianScaleType{
				"x": &chartjs.TimeScaleOptions{
					CartesianScaleOptions: &chartjs.CartesianScaleOptions{
						Title: &chartjs.CartesianScaleTitle{
							Display: pointer.ToBool(true),
							Text:    "Time",
						},
					},
					Time: &chartjs.TimeScaleTimeOptions{
						DisplayFormats: map[string]string{
							"month":  "yyyy-MM",
							"day":    "yyyy-MM-dd",
							"hour":   "yyyy-MM-dd HH:mm",
							"minute": "yyyy-MM-dd HH:mm",
							"second": "yyyy-MM-dd HH:mm:ss",
						},
						TooltipFormat: "yyy-MM-dd HH:mm",
						MinUnit:       chartjs.TimeUnitHour,
					},
				},
				"y": &chartjs.LinearScaleOptions{
					CartesianScaleOptions: &chartjs.CartesianScaleOptions{
						Title: &chartjs.CartesianScaleTitle{
							Display: pointer.ToBool(true),
							Text:    "Value",
						},
					},
					BeginAtZero: pointer.ToBool(true),
				},
			},
			Plugins: &chartjs.PluginOptions{
				Title: &chartjs.TitleOptions{
					Display: pointer.ToBool(true),
					Text:    chartjs.TitleText("Chart.js Line Chart - Time scale"),
				},
			},
		},
	}, nil
}
