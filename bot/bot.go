package bot

import (
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/telegram-bot-api.v4"
)

// Wrapper struct for a message
type botmessage struct {
	Cmd  string
	Args []string
	Msg  *tgbotapi.Message
}

// ResponseFunc is a handler for a bot command.
type ResponseFunc func(m *botmessage)

// Bot is the main strcut.
type Bot struct {
	telegramApi string
	Bot         *tgbotapi.BotAPI
	fmap        map[string]ResponseFunc
	db          *gorm.DB
	log         *log.Logger
}

// NewBot create the
func NewBot(token string, lg *log.Logger) *Bot {
	cbc := new(Bot)
	cbc.telegramApi = token
	var err error
	if lg == nil {
		cbc.log = log.New(os.Stdout, "[bot] ", 0)
	}
	cbc.Bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		cbc.log.Fatal(err)
	}
	//todo db

	//function map
	cbc.fmap = cbc.loadDefaultFunc()

	return cbc
}
func (b *Bot) parseMessage(msg *tgbotapi.Message) *botmessage {
	cmd := ""
	args := []string{}
	if msg.ReplyToMessage != nil {

	} else if msg.Text != "" {
		msgTokens := strings.Fields(msg.Text)
		cmd, args = strings.ToLower(msgTokens[0]), msgTokens[1:]
		//inline mode
		if strings.Contains(cmd, "@") {
			c := strings.Split(cmd, "@")
			cmd = c[0]
		}
	}
	return &botmessage{Cmd: cmd, Args: args, Msg: msg}
}
func (b *Bot) Router(msg tgbotapi.Message) {
	//don't handle forward message
	if msg.ForwardFrom != nil || msg.ForwardFromChat != nil {
		return
	}
	bmsg := b.parseMessage(&msg)
	if bmsg.Cmd != "" {
		b.log.Println(bmsg.Cmd, bmsg.Args)
	}
	execFn := b.fmap[bmsg.Cmd]
	if execFn != nil {
		b.GoSafely(func() { execFn(bmsg) })
	}

}
func (b *Bot) loadDefaultFunc() map[string]ResponseFunc {
	return map[string]ResponseFunc{
		"/start":     b.Start,
		"/price":     b.GetPrice,
		"/set_alarm": b.SetAlarm,
	}
}

// GoSafely is a utility wrapper to recover and log panics in goroutines.
// If we use naked goroutines, a panic in any one of them crashes
// the whole program. Using GoSafely prevents this.
// It's a good habit
func (b *Bot) GoSafely(fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				stack := make([]byte, 1024*8)
				stack = stack[:runtime.Stack(stack, false)]

				b.log.Printf("PANIC: %s\n%s", err, stack)
			}
		}()

		fn()
	}()
}
