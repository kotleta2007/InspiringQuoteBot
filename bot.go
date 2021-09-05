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
	//fmt.Println()
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		//URL: "http://195.129.111.17:8012",
		URL: "",

		Token:  "",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
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

	fmt.Println("bot started")

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hello World!")
	})

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, "Say /generate")
	})

	b.Handle("/generate", func(m *tb.Message) {
		b.Send(m.Sender, "generating image")

		time.Sleep(1 * time.Second)

		//url, err := page.EvalOnSelector("img.generated-image", "el => el.src")
		url, _ := page.EvalOnSelector("img.generated-image", "el => el.src")

		/*
			if _, err = page.Screenshot(playwright.PageScreenshotOptions{
				Path: playwright.String("foo.png"),
			}); err != nil {
				log.Fatalf("could not create screenshot: %v", err)
			}
		*/

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
