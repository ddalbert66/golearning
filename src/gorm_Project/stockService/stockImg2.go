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
	//plotutil.AddBoxPlots(plt *plot.Plot, width vg.Length, vs ...interface{})
	plotutil.AddLinePoints(p, "TurnoverRate", randomPoints(stockArray))

	p.Save(4*vg.Inch, 4*vg.Inch, "price"+code+".png")

	p2, _ := plot.New()

	p2.Title.Text = code
	p2.X.Label.Text = "date"
	p2.Y.Label.Text = "rate"
	//plotutil.AddBoxPlots(plt *plot.Plot, width vg.Length, vs ...interface{})
	plotutil.AddLinePoints(p2, "ClosingPrice", randomPoints2(stockArray))

	p2.Save(4*vg.Inch, 4*vg.Inch, "price"+code+"_2.png")
}

func randomPoints(stockArray []stockDao.Stock) plotter.XYs {
	points := make(plotter.XYs, len(stockArray))

	for i := range points {
		//fmt.Println(stockArray[i].TurnoverRate)
		points[i].X = float64(i)
		points[i].Y = (stockArray[i].TurnoverRate)
	}

	return points
}

func randomPoints2(stockArray []stockDao.Stock) plotter.XYs {
	points := make(plotter.XYs, len(stockArray))

	for i := range points {
		fmt.Print(stockArray[i].ClosingPrice)
		fmt.Print(" ")
		points[i].X = float64(i)
		points[i].Y = stockArray[i].ClosingPrice
	}
	fmt.Println()
	return points
}
