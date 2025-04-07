package line

import (
	"github.com/ngyewch/go-chartjs"
	"go.octolab.org/pointer"
	"math/rand/v2"
	"time"
)

func Example1() (chartjs.IChartConfiguration, error) {
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
				Animation: &chartjs.AnimationSpec{
					Duration: pointer.ToFloat64(0),
				},
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
					Text:    chartjs.String("Line Chart / Cubic interpolation mode"),
				},
			},
		},
	}, nil
}

func Example2() (chartjs.IChartConfiguration, error) {
	const minY = -100
	const maxY = 100
	var data []chartjs.Point
	t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
	for i := 0; i < 8; i++ {
		data = append(data, chartjs.Point{
			X: float64(t.UnixMilli()),
			Y: pointer.ToFloat64(randFloat64(minY, maxY)),
		})
		t = t.Add(1 * time.Hour)
	}
	return &chartjs.LineChartConfiguration{
		Data: &chartjs.LineChartData{
			Datasets: []*chartjs.LineChartDataset{
				{
					ControllerDatasetOptions: &chartjs.ControllerDatasetOptions{
						Label: "Series 1",
					},
					Data: data,
				},
			},
		},
		Options: &chartjs.LineControllerChartOptions{
			CoreChartOptions: &chartjs.CoreChartOptions{
				Animation: &chartjs.AnimationSpec{
					Duration: pointer.ToFloat64(0),
				},
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
					Text:    chartjs.String("Line Chart / Time scale"),
				},
			},
		},
	}, nil
}

func Example3() (chartjs.IChartConfiguration, error) {
	const minY = -100
	const maxY = 100
	var data1 []chartjs.Point
	var data2 []chartjs.Point
	for i := 0; i < 8; i++ {
		data1 = append(data1, chartjs.Point{
			X: float64(i),
			Y: pointer.ToFloat64(randFloat64(minY, maxY)),
		})
		data2 = append(data2, chartjs.Point{
			X: float64(i),
			Y: pointer.ToFloat64(randFloat64(minY, maxY)),
		})
	}
	return &chartjs.LineChartConfiguration{
		Data: &chartjs.LineChartData{
			Datasets: []*chartjs.LineChartDataset{
				{
					Data:    data1,
					YAxisID: "y",
					ControllerDatasetOptions: &chartjs.ControllerDatasetOptions{
						Label: "Dataset 1",
					},
					LineOptions: &chartjs.LineOptions{
						CommonElementOptions: &chartjs.CommonElementOptions{
							BorderColor: "red",
						},
					},
				},
				{
					Data:    data2,
					YAxisID: "y1",
					ControllerDatasetOptions: &chartjs.ControllerDatasetOptions{
						Label: "Dataset 2",
					},
					LineOptions: &chartjs.LineOptions{
						CommonElementOptions: &chartjs.CommonElementOptions{
							BorderColor: "blue",
						},
					},
				},
			},
		},
		Options: &chartjs.LineControllerChartOptions{
			CoreChartOptions: &chartjs.CoreChartOptions{
				Animation: &chartjs.AnimationSpec{
					Duration: pointer.ToFloat64(0),
				},
				MaintainAspectRatio: pointer.ToBool(false),
			},
			Scales: map[string]chartjs.ICartesianScaleType{
				"x": &chartjs.LinearScaleOptions{},
				"y": &chartjs.LinearScaleOptions{},
				"y1": &chartjs.LinearScaleOptions{
					CartesianScaleOptions: &chartjs.CartesianScaleOptions{
						Position: chartjs.CartesianAxisPositionRight,
						Grid: &chartjs.GridLineOptions{
							DrawOnChartArea: pointer.ToBool(false),
						},
					},
				},
			},
			Plugins: &chartjs.PluginOptions{
				Title: &chartjs.TitleOptions{
					Display: pointer.ToBool(true),
					Text:    chartjs.String("Line Chart / Multi Axis"),
				},
			},
		},
	}, nil
}

func Example4() (chartjs.IChartConfiguration, error) {
	const minY = -100
	const maxY = 100
	var data []chartjs.Point
	for i := 0; i < 8; i++ {
		data = append(data, chartjs.Point{
			X: float64(i),
			Y: pointer.ToFloat64(randFloat64(minY, maxY)),
		})
	}
	return &chartjs.LineChartConfiguration{
		Data: &chartjs.LineChartData{
			Datasets: []*chartjs.LineChartDataset{
				{
					Data: data,
					ControllerDatasetOptions: &chartjs.ControllerDatasetOptions{
						Label: "Dataset",
					},
					LineOptions: &chartjs.LineOptions{
						CommonElementOptions: &chartjs.CommonElementOptions{
							BorderColor:     "red",
							BackgroundColor: "rgba(255, 0, 0, 0.5)",
						},
					},
					PointPrefixedOptions: &chartjs.PointPrefixedOptions{
						PointRadius: pointer.ToFloat64(10),
						PointStyle:  chartjs.PointStyleCircle,
					},
					PointPrefixedHoverOptions: &chartjs.PointPrefixedHoverOptions{
						PointHoverRadius: pointer.ToFloat64(15),
					},
				},
			},
		},
		Options: &chartjs.LineControllerChartOptions{
			CoreChartOptions: &chartjs.CoreChartOptions{
				Animation: &chartjs.AnimationSpec{
					Duration: pointer.ToFloat64(0),
				},
				MaintainAspectRatio: pointer.ToBool(false),
			},
			Scales: map[string]chartjs.ICartesianScaleType{
				"x": &chartjs.LinearScaleOptions{},
				"y": &chartjs.LinearScaleOptions{},
			},
			Plugins: &chartjs.PluginOptions{
				Title: &chartjs.TitleOptions{
					Display: pointer.ToBool(true),
					Text:    chartjs.String("Line Chart / Point Styling"),
				},
			},
		},
	}, nil
}

func Example5() (chartjs.IChartConfiguration, error) {
	const minY = -100
	const maxY = 100
	var data []chartjs.Point
	for i := 0; i < 8; i++ {
		data = append(data, chartjs.Point{
			X: float64(i),
			Y: pointer.ToFloat64(randFloat64(minY, maxY)),
		})
	}
	return &chartjs.LineChartConfiguration{
		Data: &chartjs.LineChartData{
			Datasets: []*chartjs.LineChartDataset{
				{
					Data: data,
					ControllerDatasetOptions: &chartjs.ControllerDatasetOptions{
						Label: "Dataset",
					},
					LineOptions: &chartjs.LineOptions{
						CommonElementOptions: &chartjs.CommonElementOptions{
							BorderColor:     "red",
							BackgroundColor: "rgba(255, 0, 0, 0.5)",
						},
						Stepped: chartjs.Bool(true),
					},
				},
			},
		},
		Options: &chartjs.LineControllerChartOptions{
			CoreChartOptions: &chartjs.CoreChartOptions{
				Animation: &chartjs.AnimationSpec{
					Duration: pointer.ToFloat64(0),
				},
				MaintainAspectRatio: pointer.ToBool(false),
				Interaction: &chartjs.CoreInteractionOptions{
					Intersect: pointer.ToBool(false),
					Axis:      "x",
				},
			},
			Scales: map[string]chartjs.ICartesianScaleType{
				"x": &chartjs.LinearScaleOptions{},
				"y": &chartjs.LinearScaleOptions{},
			},
			Plugins: &chartjs.PluginOptions{
				Title: &chartjs.TitleOptions{
					Display: pointer.ToBool(true),
					Text:    chartjs.String("Line Chart / Stepped"),
				},
			},
		},
	}, nil
}

func Example6() (chartjs.IChartConfiguration, error) {
	const minY = -100
	const maxY = 100
	var data1 []chartjs.Point
	var data2 []chartjs.Point
	var data3 []chartjs.Point
	for i := 0; i < 8; i++ {
		data1 = append(data1, chartjs.Point{
			X: float64(i),
			Y: pointer.ToFloat64(randFloat64(minY, maxY)),
		})
		data2 = append(data2, chartjs.Point{
			X: float64(i),
			Y: pointer.ToFloat64(randFloat64(minY, maxY)),
		})
		data3 = append(data3, chartjs.Point{
			X: float64(i),
			Y: pointer.ToFloat64(randFloat64(minY, maxY)),
		})
	}
	return &chartjs.LineChartConfiguration{
		Data: &chartjs.LineChartData{
			Datasets: []*chartjs.LineChartDataset{
				{
					Data: data1,
					ControllerDatasetOptions: &chartjs.ControllerDatasetOptions{
						Label: "Unfilled",
					},
					LineOptions: &chartjs.LineOptions{
						CommonElementOptions: &chartjs.CommonElementOptions{
							BackgroundColor: "blue",
							BorderColor:     "blue",
						},
					},
				},
				{
					Data: data2,
					ControllerDatasetOptions: &chartjs.ControllerDatasetOptions{
						Label: "Dashed",
					},
					LineOptions: &chartjs.LineOptions{
						CommonElementOptions: &chartjs.CommonElementOptions{
							BackgroundColor: "green",
							BorderColor:     "green",
						},
						BorderDash: []float64{5, 5},
					},
				},
				{
					Data: data3,
					ControllerDatasetOptions: &chartjs.ControllerDatasetOptions{
						Label: "Filled",
					},
					LineOptions: &chartjs.LineOptions{
						CommonElementOptions: &chartjs.CommonElementOptions{
							BackgroundColor: "red",
							BorderColor:     "red",
						},
						Fill: chartjs.Bool(true),
					},
				},
			},
		},
		Options: &chartjs.LineControllerChartOptions{
			CoreChartOptions: &chartjs.CoreChartOptions{
				Animation: &chartjs.AnimationSpec{
					Duration: pointer.ToFloat64(0),
				},
				MaintainAspectRatio: pointer.ToBool(false),
			},
			Scales: map[string]chartjs.ICartesianScaleType{
				"x": &chartjs.LinearScaleOptions{},
				"y": &chartjs.LinearScaleOptions{},
			},
			Plugins: &chartjs.PluginOptions{
				Title: &chartjs.TitleOptions{
					Display: pointer.ToBool(true),
					Text:    chartjs.String("Line Chart / Line Styling"),
				},
			},
		},
	}, nil
}

func randFloat64(min float64, max float64) float64 {
	return (rand.Float64() * (max - min)) + min
}
