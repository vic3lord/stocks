package stocks

// Stock structure
type Stock struct {
	List struct {
		Meta struct {
			Type  string `json:"type"`
			Start int    `json:"start"`
			Count int    `json:"count"`
		} `json:"meta"`
		Resources []struct {
			Resource struct {
				Classname string `json:"classname"`
				Fields    struct {
					Change         string `json:"change"`
					ChgPercent     string `json:"chg_percent"`
					DayHigh        string `json:"day_high"`
					DayLow         string `json:"day_low"`
					IssuerName     string `json:"issuer_name"`
					IssuerNameLang string `json:"issuer_name_lang"`
					Name           string `json:"name"`
					Price          string `json:"price"`
					Symbol         string `json:"symbol"`
					Ts             string `json:"ts"`
					Type           string `json:"type"`
					Utctime        string `json:"utctime"`
					Volume         string `json:"volume"`
					YearHigh       string `json:"year_high"`
					YearLow        string `json:"year_low"`
				} `json:"fields"`
			} `json:"resource"`
		} `json:"resources"`
	} `json:"list"`
}

// {
//    "list":{
//       "meta":{
//          "type":"resource-list",
//          "start":0,
//          "count":1
//       },
//       "resources":[
//          {
//             "resource":{
//                "classname":"Quote",
//                "fields":{
//                   "change":"-0.045000",
//                   "chg_percent":"-0.144834",
//                   "day_high":"31.070000",
//                   "day_low":"30.820000",
//                   "issuer_name":"General Electric Company",
//                   "issuer_name_lang":"General Electric Company",
//                   "name":"General Electric Company Common",
//                   "price":"31.025000",
//                   "symbol":"GE",
//                   "ts":"1458847026",
//                   "type":"equity",
//                   "utctime":"2016-03-24T19:17:06+0000",
//                   "volume":"18048790",
//                   "year_high":"31.490000",
//                   "year_low":"19.370000"
//                }
//             }
//          }
//       ]
//    }
// }
