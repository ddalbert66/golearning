package main

import (
	"goLearning20200930/src/gorm_Project/dao"
	"goLearning20200930/src/gorm_Project/dao/stockDao"
	"goLearning20200930/src/gorm_Project/stockModels"

	"github.com/jinzhu/gorm"
	"gonum.org/v1/plot"

	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"

	"gonum.org/v1/plot/vg"
)

type Stock stockModels.Stock

func db() *gorm.DB {
	return dao.GetDB()
}

func main() {

	stockArray := stockDao.QueryStockByName("全宇昕")

	p, _ := plot.New()

	p.Title.Text = "全宇昕"
	p.X.Label.Text = "date"
	p.Y.Label.Text = "rate"

	var plotterXys plotter.XYs
	for i, st := range stockArray {
		x := plotter.XY{float64(i), st.TurnoverRate}
		plotterXys = append(plotterXys, x)
	}

	plotutil.AddLinePoints(p, plotterXys)

	p.Save(4*vg.Inch, 4*vg.Inch, "price.png")
}
