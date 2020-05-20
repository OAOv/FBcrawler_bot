package task

import (
	"FBcrawler/types"
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

type FacebookCrawler struct{}

var targetURL string

func (fbc *FacebookCrawler) Do(keyword string, records *[]*types.Record) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-popup-blocking", true),
	)

	allocator, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := context.WithTimeout(allocator, 60*time.Second)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	targetURL = "https://www.facebook.com/pg/" + keyword + "/posts/?ref=page_internal"

	var record [3]types.Record
	err := chromedp.Run(ctx, fbc.getTitle(keyword, &record))
	if err != nil {
		log.Println("run_error: " + err.Error())
	}

	for i := 0; i < 3; i++ {
		record[i].URL = "https://www.facebook.com" + record[i].URL
		*records = append(*records, &record[i])
	}
}

func (fbc *FacebookCrawler) getTitle(keyword string, record *[3]types.Record) chromedp.Tasks {
	var ok bool
	return chromedp.Tasks{
		chromedp.Navigate(targetURL),
		chromedp.Sleep(3 * time.Second),
		chromedp.Text(`//*[@id="pagelet_timeline_main_column"]/div/div[2]/div/div[1]/div/div/div/div/div[1]/div[3]/div[2]`, &record[0].Title),
		chromedp.AttributeValue(`//*[@id="pagelet_timeline_main_column"]/div/div[2]/div/div[1]/div/div/div/div/div[1]/div[3]/div[1]/div/div/div[2]/div/div/div[2]/div/span[1]/span/a`, "href", &record[0].URL, &ok),
		chromedp.Text(`//*[@id="pagelet_timeline_main_column"]/div/div[2]/div/div[2]/div/div/div/div/div[1]/div[3]/div[2]`, &record[1].Title),
		chromedp.AttributeValue(`//*[@id="pagelet_timeline_main_column"]/div/div[2]/div/div[2]/div/div/div/div/div[1]/div[3]/div[1]/div/div/div[2]/div/div/div[2]/div/span[1]/span/a`, "href", &record[1].URL, &ok),
		chromedp.Text(`//*[@id="pagelet_timeline_main_column"]/div/div[2]/div/div[3]/div/div/div/div/div[1]/div[3]/div[2]`, &record[2].Title),
		chromedp.AttributeValue(`//*[@id="pagelet_timeline_main_column"]/div/div[2]/div/div[3]/div/div/div/div/div[1]/div[3]/div[1]/div/div/div[2]/div/div/div[2]/div/span[1]/span/a`, "href", &record[2].URL, &ok),
	}
}
