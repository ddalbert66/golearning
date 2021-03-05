package stockService

import (
	"fmt"
	"goLearning20200930/src/gorm_Project/dao"
	"goLearning20200930/src/gorm_Project/dao/stockDao"

	"github.com/jinzhu/gorm"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

//type Stock stockModels.Stock

func db() *gorm.DB {
	return dao.GetDB()
}

func FuncgetGoodStock(code string) {
	stockArray := stockDao.QueryStockByCode(code)

	p, _ := plot.New()

	p.Title.Text = code
	p.X.Label.Text = "date"
	p.Y.Label.Text = "rate"

	plotutil.AddLinePoints(p, "TurnoverRate*1000", randomPoints(stockArray),
		"StockVolume 1000", randomPoints2(stockArray))

	p.Save(4*vg.Inch, 4*vg.Inch, "price"+code+".png")
}

func randomPoints(stockArray []stockDao.Stock) plotter.XYs {
	points := make(plotter.XYs, len(stockArray))

	for i := range points {
		fmt.Println(stockArray[i].TurnoverRate)
		points[i].X = float64(i)
		points[i].Y = (stockArray[i].TurnoverRate * 1000)
	}

	return points
}

func randomPoints2(stockArray []stockDao.Stock) plotter.XYs {
	points := make(plotter.XYs, len(stockArray))

	for i := range points {
		fmt.Println(float64(stockArray[i].StockVolume / 1000))
		points[i].X = float64(i)
		points[i].Y = float64(stockArray[i].StockVolume / 1000)
	}

	return points
}
