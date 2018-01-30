package bot

import (
	"fmt"
	"strconv"
	"time"
)

func (b Bot) SetAlarm(m *botmessage) {
	if len(m.Args) != 1 {
		b.SendMessage(m.Msg, "/set_alarm: Set a alarm from the point.\n Only need a argument, unit is minute, like this:\n/set_alarm 30", true)
	} else {
		min, _ := strconv.Atoi(m.Args[0])
		b.GoSafely(func() {
			a := time.NewTimer(time.Minute * time.Duration(min))
			<-a.C
			b.SendMessage(m.Msg, "Alarm! Alarm! Alarm!", true)
		})
		b.SendMessage(m.Msg, fmt.Sprintf("/set_alarm: Set a alarm successfully,will notice you in %dmin", min), true)
	}
}
