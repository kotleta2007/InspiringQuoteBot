package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mxschmitt/playwright-go"
	tb "gopkg.in/tucnak/telebot.v2"
)

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func main() {
	b, err := tb.NewBot(tb.Settings{
		URL: "",

		Token:  "",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatalf("could not create bot: %v", err)
		return
	}

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

	if _, err = page.Goto("https://inspirobot.me/"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	assertErrorToNilf("could not click: %v", page.Click("div.btn-text"))

	fmt.Println("Bot started.")

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, "Say /generate")
	})

	b.Handle("/generate", func(m *tb.Message) {
		b.Send(m.Sender, "generating image")

		assertErrorToNilf("could not click: %v", page.Click("div.btn-text"))

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

	b.Start()
}
