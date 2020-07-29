package chart

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/SoundRequest/backend/structure"
	"github.com/gin-gonic/gin"
)

func Melon(c *gin.Context) {
	// Request the HTML page.
	res, err := http.Get("https://www.melon.com/chart/day/index.htm")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var list []structure.MelonChart

	// Find the review items
	doc.Find(".lst50").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".rank01 a").Text()
		artist := s.Find(".rank02 a").Text()
		album := s.Find(".rank03 a").Text()
		list = append(list, structure.MelonChart{Count: i, Title: title, Artist: artist, Album: album})
	})
	c.JSON(http.StatusOK, gin.H{"list": list})
}
