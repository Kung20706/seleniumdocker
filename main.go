package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/tebeka/selenium"
)

const (
	port            = 8087
	rmtport         = 4444
	seleniumhost    = "chrome"
	maxAttempts     = 5
	retryInterval   = 11
	containerdbhost = "db"        // docker runtime
	localdbhost     = "127.0.0.1" // local runtime
	sqlport         = ":3306"
)

func main() {

	// Set up Remote WebDriver
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, "http://chrome:4444/wd/hub")
	if err != nil {
		log.Print("Failed to open session:", err)
	}
	defer func() {
		if wd != nil {
			wd.Quit()
		}
	}()
	// Test Google page title
	err = wd.Get("http://www.google.com")
	if err != nil {
		log.Print("Failed to load page:", err)
	}
	title, err := wd.Title()
	if err != nil {
		log.Print("Failed to get page title:", err)
	}
	if title != "Google" {
		log.Print("Unexpected page title: got %q, want %q", title, "Google")
	}
	log.Print(wd.Title())
	MCGfantasyoflotteryusa(wd)

}
func ParseDate(input string) (string, error) {
	// Define the layout pattern of the input date string
	// Note: In Go, the reference time is "Mon Jan 2 15:04:05 MST 2006"
	const inputLayout = "Monday, Jan 2, 2006"

	// Define the desired output format
	const outputLayout = "20060102"

	// Parse the input string into time.Time
	parsedTime, err := time.Parse(inputLayout, input)
	if err != nil {
		return "", err
	}

	// Format the time into the desired output format
	return parsedTime.Format(outputLayout), nil
}
func MCGfantasyoflotteryusa(wd selenium.WebDriver) {
	// 取得 第一個分頁的遊戲表(包括跨境遊戲)

	soruceurl := "https://www.lotteryusa.com/michigan/fantasy-5/"
	if err := wd.Get(soruceurl); err != nil {
		log.Fatalf("Error opening the website: %v", err)
	}

	Source, err := wd.PageSource()
	if err != nil {
		log.Fatalf("Failed to get page source: %v", err)
	}

	// ballsoruce := "":nth-child(%d)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(Source))
	if err != nil {
		log.Fatal(err)
	}

	roundsoruce := fmt.Sprintf(" tr.c-result-card.c-result-card--squeeze ")
	doc.Find(roundsoruce).Each(func(i int, s *goquery.Selection) {
		// 获取日期
		date := s.Find("time.c-result-card__title").Text()

		// 获取每个球的号码
		var resultBuilder strings.Builder
		s.Find("span.c-ball__label").Each(func(j int, span *goquery.Selection) {
			if j > 0 {
				resultBuilder.WriteString(",")
			}
			resultBuilder.WriteString(span.Text())
		})
		fmt.Println(date, resultBuilder.String())

	})

	fmt.Println("MCG")
}
