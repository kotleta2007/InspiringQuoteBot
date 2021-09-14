package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mxschmitt/playwright-go"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	// Creating bot
	b, err := tb.NewBot(tb.Settings{
		URL: "",

		Token:  "",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatalf("could not create bot: %v", err)
		return
	}

	// FROM HERE

	// Opening browser, navigating to the page, loading first image
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	_, err = page.Goto("https://inspirobot.me/")
	if err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	err = page.Click("div.btn-text")
	if err != nil {
		log.Fatalf("could not click: %v", err)
	}

	// TO HERE
	// move to corresponding goroutine

	// check if message channel is full, then start

	fmt.Println("Bot started.")

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, "Say /generate")
	})

	b.Handle("/generate", func(m *tb.Message) {
		b.Send(m.Sender, "generating image")

		err = page.Click("div.btn-text")
		if err != nil {
			log.Fatalf("could not click: %v", err)
		}

		time.Sleep(1 * time.Second)

		//url, err := page.EvalOnSelector("img.generated-image", "el => el.src")
		url, _ := page.EvalOnSelector("img.generated-image", "el => el.src")

		/*
			if err = browser.Close(); err != nil {
				log.Fatalf("could not close browser: %v", err)
			}
			if err = pw.Stop(); err != nil {
				log.Fatalf("could not stop Playwright: %v", err)
			}
		*/

		fmt.Println(url)

		urlString := fmt.Sprintf("%s", url)

		p := &tb.Photo{File: tb.FromURL(urlString)}

		b.Send(m.Sender, p)

	})

	go b.Start()

	// TODO: replace with signal handling
	fmt.Scanln()
}
