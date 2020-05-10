package main

import (
	"log"
	"os"
	"time"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		port      = os.Getenv("PORT")
		publicURL = os.Getenv("PUBLIC_URL") 
		token     = os.Getenv("TOKEN")      
	)
	
// Webhook	
	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
	
//for inline keybord buttons, Defining buttons
	
	inlineBtn1 := tb.InlineButton{
    Unique: "play",
    Text:   "Play ▶",
	}
	
	inlineBtn2 := tb.InlineButton{
    Unique: "about",
    Text:   "About ℹ",
	}
	
	inlineBtn3 := tb.InlineButton{
    Unique: "stop",
    Text:   "Stop ⛔",
	}
	
	
// Need to defne a keyboard for buttons
	
	
	inlineKeys := [][]tb.InlineButton{
    []tb.InlineButton{inlineBtn1, inlineBtn2, inlineBtn3},
	}
	
	
	//----Handlers for Buttons
	
	
		//----Btn1
	
	b.Handle(&inlineBtn1, func(c *tb.Callback) {
        // Required for proper work
    b.Respond(c, &tb.CallbackResponse{
        ShowAlert: false,
    })
        // Send messages here	play()
    b.Send(c.Sender, "Play---- STAY tunned!")
})


		//----Btn2
		
	b.Handle(&inlineBtn2, func(c *tb.Callback) {
        // Required for proper work
    b.Respond(c, &tb.CallbackResponse{
        ShowAlert: false,
    })
        // Send messages here	about()
    b.Send(c.Sender, "About---- From a time-pass Engineer🥺 @TympazEngineer")
})

		//----Btn3		
		
	b.Handle(&inlineBtn3, func(c *tb.Callback) {
        // Required for proper work
    b.Respond(c, &tb.CallbackResponse{
        ShowAlert: false,
    })
        // Send messages here	stop()
    b.Send(c.Sender, "👋 Play again later after a while🤓!!!")
		
		time.Sleep(2 * time.Second)
		b.Stop()
})
	

	b.Handle("/start", func(m *tb.Message) {
		
		//b.Send(m.Sender, "Hello! To start playing you have to click on Play..!")	})
		
		b.Send(m.Sender,
				"Hello! To start playing you have to click on Play..!",
				&tb.ReplyMarkup{InlineKeyboard: inlineKeys})	})

		
	//b.Handle("/auth", func(m *tb.Message) {
	//	Auth()
	//	b.Send(m.Sender, "Authentificating...")
	//})
	
	

	b.Start()
}
