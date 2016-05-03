package stocks

import "time"

type Stock struct {
	Query struct {
		Count   int       `json:"count"`
		Created time.Time `json:"created"`
		Lang    string    `json:"lang"`
		Results struct {
			Quote struct {
				Symbol                                      string `json:"symbol"`
				Ask                                         string `json:"Ask"`
				AverageDailyVolume                          string `json:"AverageDailyVolume"`
				Bid                                         string `json:"Bid"`
				BookValue                                   string `json:"BookValue"`
				ChangePercentChange                         string `json:"Change_PercentChange"`
				Change                                      string `json:"Change"`
				Currency                                    string `json:"Currency"`
				LastTradeDate                               string `json:"LastTradeDate"`
				EarningsShare                               string `json:"EarningsShare"`
				EPSEstimateCurrentYear                      string `json:"EPSEstimateCurrentYear"`
				EPSEstimateNextYear                         string `json:"EPSEstimateNextYear"`
				EPSEstimateNextQuarter                      string `json:"EPSEstimateNextQuarter"`
				DaysLow                                     string `json:"DaysLow"`
				DaysHigh                                    string `json:"DaysHigh"`
				YearLow                                     string `json:"YearLow"`
				YearHigh                                    string `json:"YearHigh"`
				MarketCapitalization                        string `json:"MarketCapitalization"`
				EBITDA                                      string `json:"EBITDA"`
				ChangeFromYearLow                           string `json:"ChangeFromYearLow"`
				PercentChangeFromYearLow                    string `json:"PercentChangeFromYearLow"`
				ChangeFromYearHigh                          string `json:"ChangeFromYearHigh"`
				PercebtChangeFromYearHigh                   string `json:"PercebtChangeFromYearHigh"`
				LastTradeWithTime                           string `json:"LastTradeWithTime"`
				LastTradePriceOnly                          string `json:"LastTradePriceOnly"`
				DaysRange                                   string `json:"DaysRange"`
				FiftydayMovingAverage                       string `json:"FiftydayMovingAverage"`
				TwoHundreddayMovingAverage                  string `json:"TwoHundreddayMovingAverage"`
				ChangeFromTwoHundreddayMovingAverage        string `json:"ChangeFromTwoHundreddayMovingAverage"`
				PercentChangeFromTwoHundreddayMovingAverage string `json:"PercentChangeFromTwoHundreddayMovingAverage"`
				ChangeFromFiftydayMovingAverage             string `json:"ChangeFromFiftydayMovingAverage"`
				PercentChangeFromFiftydayMovingAverage      string `json:"PercentChangeFromFiftydayMovingAverage"`
				Name                                        string `json:"Name"`
				Open                                        string `json:"Open"`
				PreviousClose                               string `json:"PreviousClose"`
				ChangeinPercent                             string `json:"ChangeinPercent"`
				PriceSales                                  string `json:"PriceSales"`
				PriceBook                                   string `json:"PriceBook"`
				PEGRatio                                    string `json:"PEGRatio"`
				PriceEPSEstimateCurrentYear                 string `json:"PriceEPSEstimateCurrentYear"`
				PriceEPSEstimateNextYear                    string `json:"PriceEPSEstimateNextYear"`
				ShortRatio                                  string `json:"ShortRatio"`
				LastTradeTime                               string `json:"LastTradeTime"`
				OneyrTargetPrice                            string `json:"OneyrTargetPrice"`
				Volume                                      string `json:"Volume"`
				YearRange                                   string `json:"YearRange"`
				StockExchange                               string `json:"StockExchange"`
				PercentChange                               string `json:"PercentChange"`
			} `json:"quote"`
		} `json:"results"`
	} `json:"query"`
}
