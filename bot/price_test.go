package bot

import (
	"testing"
)

func Test_price(t *testing.T) {
	cbc := NewBot("497425065:AAFDaMLuxdghQblsf6QG-ByH7YB4FvETlBs", nil)
	cbc.GetPrice(&botmessage{
		Args: []string{"eosusdt", "ethusdt"},
	})
}
