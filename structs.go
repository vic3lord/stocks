package stocks

type Stock struct {
	List struct {
		Resources []Resource
	}
}

type Resource struct {
	Resource struct {
		Fields Fields
	}
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
