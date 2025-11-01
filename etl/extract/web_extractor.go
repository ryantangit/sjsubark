package extract

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/ryantangit/sjsubark/etl/config"
	"golang.org/x/net/html"
)

type WebpageExtractor struct {
	webpageUrl string
	webpageDir string
}

func NewWebpageExtractor(webpageUrl string, webpageDir string) WebpageExtractor {
	return WebpageExtractor{webpageUrl: webpageUrl, webpageDir: webpageDir}
}

// The parking record is generated from the official SJSU parking status page.
// Parse through the HTML. The HTML of the Parking Status page is labelled with classes.
// The relevant information are located within a div element with the class "garage"
// Then it will be followed by an element with class "garage_name"
// Then an element with class "garage_text" will have an element "garage__address" and "garage__fullness"

func (e WebpageExtractor) FetchRecords() []GarageRecord {

	// Ignore TLS
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	// Retry with Exponential Backoff if encounter unsucessful GET request
	max_retries := 3
	retry := 0
	var resp *http.Response
	var err error
	for retry < max_retries {
		resp, err = http.Get(e.webpageUrl)
		if err != nil {
			log.Fatal("Fetching Request Page failed", err)
		}
		if (resp.StatusCode == 200) {
			break
		}
		backoff := 1<<retry
		log.Printf("Backoff for %d seconds", backoff)
		time.Sleep(time.Duration(backoff) * time.Second)
		retry += 1
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatal("Could not HTTP GET garage data with a 200 within retry window")
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	timestamp := time.Now()
	err = os.WriteFile(filepath.Join(config.WebpageDir(), webpageFilename(timestamp)), respBody, 0644)
	if err != nil {
		log.Fatal(err)
	}

	respReader := bytes.NewReader(respBody)
	rootNode, err := html.Parse(respReader)
	if err != nil {
		log.Fatal(err)
	}
	garageDiv := findGarageDiv(rootNode)
	if garageDiv == nil {
		log.Fatal("No garage element found in Doc")
		os.Exit(1)
	}

	GarageNameClassAttr := "garage__name"
	GarageFullClassAttr := "garage__fullness"
	GarageName := []string{}
	GarageFull := []string{}
	for n := range garageDiv.Descendants() {
		if n.Type == html.ElementNode && hasClass(n, GarageNameClassAttr) {
			noWhiteSpace := strings.Join(strings.Fields(n.FirstChild.Data), "")
			GarageName = append(GarageName, noWhiteSpace)
		}
		if n.Type == html.ElementNode && hasClass(n, GarageFullClassAttr) {
			noWhiteSpace := strings.Join(strings.Fields(n.FirstChild.Data), "")
			if noWhiteSpace == "Full" {
				noWhiteSpace = "100%"
			}
			GarageFull = append(GarageFull, noWhiteSpace[:len(noWhiteSpace)-1])
		}
	}
	if len(GarageFull) != len(GarageName) {
		log.Fatal("The final results length do not match up.")
	}

	garages := []GarageRecord{}
	for idx := 0; idx < len(GarageFull); idx++ {
		fullint, err := strconv.Atoi(GarageFull[idx])
		if err != nil {
			log.Fatalf("conversion of garage fullness from string to integer errored: %v", err)
		}
		garageStatus := GarageRecord{Name: GarageName[idx], Fullness: fullint, Timestamp: timestamp}
		garages = append(garages, garageStatus)
	}

	return garages
}

func findGarageDiv(doc *html.Node) *html.Node {
	GarageClassAttr := "garage"
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && hasClass(n, GarageClassAttr) {
			return n
		}
	}
	return nil
}

func hasClass(n *html.Node, targetClass string) bool {
	for _, attr := range n.Attr {
		if attr.Key == "class" {
			classes := strings.Fields(attr.Val)
			return slices.Contains(classes, targetClass)
		}
	}
	return false
}

func webpageFilename(timestamp time.Time) string {
	year, month, day := timestamp.Date()
	hour := timestamp.Hour()
	min := timestamp.Minute()
	return fmt.Sprintf("%d_%d_%d__%d::%d_.html", month, day, year, hour, min)
}
