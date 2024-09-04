package test

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"os"
	"testing"
)

func TestRssParser(t *testing.T) {
	file, _ := os.Open(`C:\Users\ms\Desktop\Bangumi.xml`)
	defer file.Close()
	fp := gofeed.NewParser()
	feed, _ := fp.Parse(file)
	fmt.Println(feed.Title)
}
