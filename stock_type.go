package stocks

type StockType struct {
	List List
}

type List struct {
	Meta      Meta        `json:"meta"`
	Resources []Resources `json:"resources"`
}

type Meta struct {
	Type  string `json:"type"`
	Start uint   `json:"start"`
	Count uint   `json:"count"`
}

type Resources struct {
	Resource Resource `json:"resource"`
}

type Resource struct {
	Classname string `json:"classname"`
	Fields    Fields `json:"fields"`
}

type Fields struct {
	Name    string `json:"name"`
	Price   string `json:"price"`
	Symbol  string `json:"symbol"`
	Ts      string `json:"ts"`
	Type    string `json:"type"`
	UTCTime string `json:"utctime"`
	Volume  string `json:"volume"`
}
