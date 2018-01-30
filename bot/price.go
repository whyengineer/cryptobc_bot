package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	endpoint = "https://api.huobi.pro/market/trade?symbol=%s"
)

type huobidata struct {
	Amount float64 `json:"amount"`
	Ts     float64 `json:"ts"`
	Id     float64 `json:"id"`
	Price  float64 `json:"price"`
	Dir    string  `json:"direction"`
}
type huobitick struct {
	Id   float64     `json:"id"`
	Data []huobidata `json:"data"`
}
type huobires struct {
	Status string    `json:"status"`
	Ch     string    `json:"ch"`
	Ts     float64   `json:"ts"`
	Tick   huobitick `json:"tick"`
}

func (b *Bot) GetPrice(m *botmessage) {
	if len(m.Args) == 0 {
		//b.SendMessage(m.Msg,)
		//todo print the user subscribe's coin,use the database
	} else {
		for _, pair := range m.Args {
			url := fmt.Sprintf(endpoint, pair)
			resp, err := http.Get(url)
			if err != nil {
				b.log.Println(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			var res huobires
			err = json.Unmarshal(body, &res)
			if err != nil {
				b.log.Println(err)
			}
			b.log.Println(res.Tick.Data[0].Price)
		}
	}

}
