package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mxschmitt/playwright-go"
)

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func main() {
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
	fmt.Println("here")
	assertErrorToNilf("could not click: %v", page.Click("div.btn-text"))
	fmt.Println("here")

	time.Sleep(3 * time.Second)

	url, err := page.EvalOnSelector("img.generated-image", "el => el.src")

	fmt.Println(url)

	//fmt.Println(page.PageGetAttributeOptions("img.generated-image"))

	if _, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("foo.png"),
	}); err != nil {
		log.Fatalf("could not create screenshot: %v", err)
	}

	assertErrorToNilf("could not click: %v", page.Click("div.btn-text"))

	time.Sleep(3 * time.Second)

	if _, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("foo2.png"),
	}); err != nil {
		log.Fatalf("could not create screenshot: %v", err)
	}
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}

	// img.generated-image
	/*
		entries, err := page.QuerySelectorAll(".athing")
		if err != nil {
			log.Fatalf("could not get entries: %v", err)
		}
		for i, entry := range entries {
			titleElement, err := entry.QuerySelector("td.title > a")
			if err != nil {
				log.Fatalf("could not get title element: %v", err)
			}
			title, err := titleElement.TextContent()
			if err != nil {
				log.Fatalf("could not get text content: %v", err)
			}
			fmt.Printf("%d: %s\n", i+1, title)
		}
		if err = browser.Close(); err != nil {
			log.Fatalf("could not close browser: %v", err)
		}
		if err = pw.Stop(); err != nil {
			log.Fatalf("could not stop Playwright: %v", err)
		}
	*/
}
