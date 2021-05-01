package main

import (
	"fmt"
	"goLearning20200930/src/gorm_Project/dao/stockDao"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1781305192:AAEKbcA7kPAngmS_wkr_v_tlkuAklF60q34")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	var stock []stockDao.Stock
	var stockV2 []stockDao.StockV2
	var stockV2sv stockDao.StockV2sv
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		stock = stockDao.QueryStockByCode(update.Message.Text)

		stockV2 = stockDao.QueryStockV2ByCode(update.Message.Text)

		var str string
		if len(stock) == 0 && len(stockV2) == 0 {
			continue
		}
		if len(stock) != 0 {
			for _, i := range stock {
				str += fmt.Sprintf("%s %s %f %s %f\n", i.StockCode, i.StockName, i.ClosingPrice, i.DateStr, i.TurnoverRate)
			}
		} else if len(stockV2) != 0 {
			stockV2sv = stockDao.QueryStockV2svByCode(update.Message.Text)
			for _, i := range stockV2 {
				fmt.Println(i.SharesTraded)
				fmt.Println(stockV2sv.StockVolume)
				turnoverRate2 := (float64(i.SharesTraded) / float64(stockV2sv.StockVolume)) * 100
				str += fmt.Sprintf("%s %s %f %s %f\n", i.StockCode, i.StockName, i.ClosingPrice, i.DateStr, turnoverRate2)
			}
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, str)
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
