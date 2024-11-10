package main

import (
	"context"
	"fmt"
	"math"
	"os"
	"runtime"
	"sort"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	"web_lab/cmd/benchmark/test_containers"
	"web_lab/internal/models"
)

type movieStore interface {
	CreateMovie(ctx context.Context, movie models.Movie) error
	GetMovie(ctx context.Context, id int64) (models.Movie, error)
}

var movie = models.Movie{
	Title:       "title",
	Description: "description",
}

func main() {
	ctx := context.Background()
	fileStd, err := os.Create("cmd/benchmark/result1")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create file")
	}

	defer fileStd.Close()

	filePgx, err := os.Create("cmd/benchmark/result2")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create file")
	}

	defer filePgx.Close()

	valuesCreateMovieStd := make([]float64, 0, 100)
	valuesCreateMoviePgx := make([]float64, 0, 100)

	valuesCreateMovieBytesStd := make([]float64, 0, 100)
	valuesCreateMovieBytesPgx := make([]float64, 0, 100)

	valuesGetMovieStd := make([]float64, 0, 100)
	valuesGetMoviePgx := make([]float64, 0, 100)

	valuesGetMovieBytesStd := make([]float64, 0, 100)
	valuesGetMovieBytesPgx := make([]float64, 0, 100)

	rps := 10

	for j := 0; j < 100; j++ {
		if j%10 == 0 {
			fmt.Printf("Итерация № %d\n", j)
		}

		var (
			nsSumStd, memSumStd, nsSumPgx, memSumPgx                                 float64
			nsSumGetMovieStd, memSumGetMovieStd, nsSumGetMoviePgx, memSumGetMoviePgx float64
		)

		for i := 0; i < 100; i++ {
			containerStd, storeStd, err := test_containers.SetupTestDatabaseStdLib()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to setup container")
			}

			ns, mem := benchCreateMovie(storeStd, rps)
			nsSumStd += ns
			memSumStd += mem

			ns, mem = benchGetMovie(storeStd, rps)
			nsSumGetMovieStd += ns
			memSumGetMovieStd += mem

			err = containerStd.Terminate(ctx)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to terminate container")
			}

			containerPgx, storePgx, err := test_containers.SetupTestDatabasePgx()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to setup container")
			}

			ns, mem = benchCreateMovie(storePgx, rps)
			nsSumPgx += ns
			memSumPgx += mem

			ns, mem = benchGetMovie(storePgx, rps)
			nsSumGetMoviePgx += ns
			memSumGetMoviePgx += mem

			err = containerPgx.Terminate(ctx)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to terminate container")
			}
		}

		valuesCreateMovieStd = append(valuesCreateMovieStd, nsSumStd/100)
		valuesCreateMovieBytesStd = append(valuesCreateMovieBytesStd, memSumStd/100)

		valuesCreateMoviePgx = append(valuesCreateMoviePgx, nsSumPgx/100)
		valuesCreateMovieBytesPgx = append(valuesCreateMovieBytesPgx, memSumPgx/100)

		valuesGetMovieStd = append(valuesGetMovieStd, nsSumGetMovieStd/100)
		valuesGetMovieBytesStd = append(valuesGetMovieBytesStd, memSumGetMovieStd/100)

		valuesGetMoviePgx = append(valuesGetMoviePgx, nsSumGetMoviePgx/100)
		valuesGetMovieBytesPgx = append(valuesGetMovieBytesPgx, memSumGetMoviePgx/100)

		rps += 100
	}

	createGraphic(valuesCreateMovieStd, valuesCreateMoviePgx, "CreateMovieNsPerOp", "CreateMovieNsPerOp.png", "rps", "ns/op", min(valuesCreateMovieStd), max(valuesCreateMovieStd))
	createGraphic(valuesCreateMovieBytesStd, valuesCreateMovieBytesPgx, "CreateMovieBytesPerOp", "CreateMovieBytesPerOp.png", "rps", "bytes/op", min(valuesCreateMovieBytesStd), max(valuesCreateMovieBytesStd))

	createHistogram(valuesCreateMovieStd, "create_movie_hist.png")
	createHistogram(valuesCreateMoviePgx, "create_movie_pgx_hist.png")

	createGraphic(valuesGetMovieStd, valuesGetMoviePgx, "GetMovieNsPerOp", "GetMovieNsPerOp.png", "rps", "ns/op", min(valuesGetMovieStd), max(valuesGetMovieStd))
	createGraphic(valuesGetMovieBytesStd, valuesGetMovieBytesPgx, "GetMovieBytesPerOp", "GetMovieBytesPerOp.png", "rps", "bytes/op", min(valuesGetMovieBytesStd), max(valuesGetMovieBytesStd))

	createHistogram(valuesGetMovieStd, "get_movie_hist.png")
	createHistogram(valuesGetMoviePgx, "get_movie_pgx_hist.png")
}

func max(arr []float64) float64 {
	var maxEl float64

	for _, v := range arr {
		if v > maxEl {
			maxEl = v
		}
	}

	return maxEl
}

func min(arr []float64) float64 {
	var maxEl = arr[0]

	for _, v := range arr {
		if v < maxEl {
			maxEl = v
		}
	}

	return maxEl
}

func benchCreateMovie(movieCreator movieStore, rps int) (float64, float64) {
	wg := sync.WaitGroup{}

	sumMem := uint64(0)
	sumNs := int64(0)
	mu := &sync.Mutex{}

	for i := 0; i < rps; i++ {
		wg.Add(1)
		go func() {
			runtime.GC()
			var memStats runtime.MemStats
			runtime.ReadMemStats(&memStats)
			startMem := memStats.Alloc
			start := time.Now()
			err := movieCreator.CreateMovie(context.Background(), movie)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to create movie")
			}
			since := time.Since(start).Nanoseconds()
			var memStatsAfter runtime.MemStats
			runtime.ReadMemStats(&memStatsAfter)
			endAlloc := memStatsAfter.Alloc

			mu.Lock()
			sumNs += since
			sumMem += endAlloc - startMem
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	return float64(sumNs / int64(rps)), float64(sumMem / uint64(rps))
}

func benchGetMovie(movieCreator movieStore, rps int) (float64, float64) {
	wg := sync.WaitGroup{}

	sumMem := uint64(0)
	sumNs := int64(0)
	mu := &sync.Mutex{}

	for i := 0; i < rps; i++ {
		wg.Add(1)
		go func() {
			runtime.GC()
			var memStats runtime.MemStats
			runtime.ReadMemStats(&memStats)
			startMem := memStats.Alloc
			start := time.Now()
			_, err := movieCreator.GetMovie(context.Background(), 1)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to get movie")
			}
			since := time.Since(start).Nanoseconds()
			var memStatsAfter runtime.MemStats
			runtime.ReadMemStats(&memStatsAfter)
			endAlloc := memStatsAfter.Alloc

			mu.Lock()
			sumNs += since
			sumMem += endAlloc - startMem
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	return float64(sumNs / int64(rps)), float64(sumMem / uint64(rps))
}

func createHistogram(values []float64, histName string) {
	percentiles := []float64{50, 75, 90, 95, 99}

	percentilesValues := make([]float64, 0, len(percentiles))

	for _, p := range percentiles {
		percentilesValues = append(percentilesValues, percentile(values, p))
	}

	histData := make(plotter.Values, len(values))
	for i, value := range values {
		for j, percentileValue := range percentilesValues {
			if value <= percentileValue {
				histData[i] = percentiles[j]
				break
			}
		}
	}

	// Создание гистограммы
	hist, err := plotter.NewHist(histData, 30)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create new hist")
	}

	hist.LineStyle.Width = vg.Length(2)
	hist.LineStyle.Color = plotutil.Color(1)
	hist.FillColor = plotutil.Color(1)

	// Создание нового графика
	p := plot.New()

	// Настройка заголовка и меток
	p.Title.Text = "Гистограмма по перцентилям времени на одну операцию"
	p.X.Label.Text = "Перцентили"
	p.Y.Label.Text = "Количество значений"

	// Добавление гистограммы на график
	p.Add(hist)

	p.X.Min = 50
	p.X.Max = 100

	p.X.Tick.Marker = plot.ConstantTicks([]plot.Tick{
		{50, fmt.Sprintf("50%%\n%.0f", percentilesValues[0])},
		{75, fmt.Sprintf("75%%\n%.0f", percentilesValues[1])},
		{90, fmt.Sprintf("90%%\n%.0f", percentilesValues[2])},
		{95, fmt.Sprintf("95%%\n%.0f", percentilesValues[3])},
		{99, fmt.Sprintf("99%%\n%.0f", percentilesValues[4])},
	})

	// Сохранение графика в файл
	if err := p.Save(8*vg.Inch, 4*vg.Inch, fmt.Sprintf("cmd/benchmark_no_const/graphics/%s", histName)); err != nil {
		log.Fatal().Err(err).Msg("failed to save histogram")
	}
}

func createGraphic(valuesFirst, valuesSecond []float64, graphicName, fileName, x, y string, ymin, ymax float64) {
	p := plot.New()

	// Настройка заголовка и меток
	p.Title.Text = graphicName
	p.X.Label.Text = x
	p.Y.Label.Text = y
	p.Y.Min = ymin
	p.Y.Max = ymax

	// Создание данных для графика
	points := make(plotter.XYs, len(valuesFirst))
	for i := range points {
		points[i].X = float64(100*i + 10)
		points[i].Y = valuesFirst[i]
	}

	// Добавление данных на график
	line, err := plotter.NewLine(points)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create graphic")
	}

	line.LineStyle.Color = plotutil.Color(2)

	p.Add(line)
	p.Legend.Add("std lib", line)

	for i := range points {
		points[i].X = float64(100*i + 10)
		points[i].Y = valuesSecond[i]
	}

	// Добавление данных на график
	line, err = plotter.NewLine(points)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create graphic")
	}

	line.LineStyle.Color = plotutil.Color(3)

	p.Add(line)
	p.Legend.Add("pgx", line)

	// Сохранение графика в файл
	if err := p.Save(8*vg.Inch, 4*vg.Inch, fmt.Sprintf("cmd/benchmark_no_const/graphics/%s", fileName)); err != nil {
		log.Fatal().Err(err).Msg("failed to save graphic")
	}
}

func percentile(data []float64, percentile float64) float64 {
	// Сортируем данные
	sort.Float64s(data)

	// Вычисляем индекс перцентиля
	index := (float64(len(data)) - 1) * (percentile / 100.0)

	// Если индекс является целым числом, возвращаем значение по этому индексу
	if math.Floor(index) == index {
		return data[int(index)]
	}

	// Если индекс не является целым числом, используем линейную интерполяцию
	lowerIndex := int(math.Floor(index))
	upperIndex := int(math.Ceil(index))
	lowerValue := data[lowerIndex]
	upperValue := data[upperIndex]

	return lowerValue + (upperValue-lowerValue)*(index-float64(lowerIndex))
}
