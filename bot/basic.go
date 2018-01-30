package bot

func (b *Bot) Start(m *botmessage) {
	b.SendMessage(m.Msg, `There is a cryptobc.info robot with the following things:

/img - gets an image
/gif - gets a gif
/google - does a Google search
/xchg - does an exchange rate conversion
/youtube - does a Youtube search
/clear - clears your NSFW images for you
/psi - returns the current PSI numbers
/echo - parrots stuff back at you
/urbandict - does an Urban Dictionary search

Give these commands a try!
		`, false)
}
