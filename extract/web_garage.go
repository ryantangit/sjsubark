package extract

import (
	"log"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type WebpageExtractor struct {
	webpageUrl string
}

func NewWebpageExtractor() WebpageExtractor {
	const ParkingStatusUrl = "https://sjsuparkingstatus.sjsu.edu/"
	return WebpageExtractor{webpageUrl: ParkingStatusUrl}
}

// The parking record is generated from the official SJSU parking status page.
// Parse through the HTML. The HTML of the Parking Status page is labelled with classes.
// The relevant information are located within a div element with the class "garage"
// Then it will be followed by an element with class "garage_name"
// Then an element with class "garage_text" will have an element "garage__address" and "garage__fullness"

func (e WebpageExtractor) FetchRecords() []GarageRecord {
	resp, err := http.Get(e.webpageUrl)
	if err != nil {
		log.Fatal("Fetching Request Page failed")
	}
	defer resp.Body.Close()
	timestamp := time.Now()
	rootNode, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	garageDiv := findGarageDiv(rootNode)
	if garageDiv == nil {
		log.Fatal("No garage element found in Doc")
		os.Exit(1)
	}

	GarageNameClassAttr := "garage__name"
	GarageAddrClassAttr := "garage__address"
	GarageFullClassAttr := "garage__fullness"
	GarageName := []string{}
	GarageAddr := []string{}
	GarageFull := []string{}
	for n := range garageDiv.Descendants() {
		if n.Type == html.ElementNode && hasClass(n, GarageNameClassAttr) {
			noWhiteSpace := strings.Join(strings.Fields(n.FirstChild.Data), "")
			GarageName = append(GarageName, noWhiteSpace)
		}
		if n.Type == html.ElementNode && hasClass(n, GarageAddrClassAttr) {
			GarageAddr = append(GarageAddr, n.FirstChild.Data)
		}
		if n.Type == html.ElementNode && hasClass(n, GarageFullClassAttr) {
			noWhiteSpace := strings.Join(strings.Fields(n.FirstChild.Data), "")
			if noWhiteSpace == "Full" {
				noWhiteSpace = "100%"
			}
			GarageFull = append(GarageFull, noWhiteSpace[:len(noWhiteSpace)-1])
		}
	}
	if !(len(GarageAddr) == len(GarageFull) && len(GarageFull) == len(GarageName)) {
		log.Fatal("The final results length do not match up.")
	}

	garages := []GarageRecord{}
	for idx := 0; idx < len(GarageAddr); idx++ {
		fullint, err := strconv.Atoi(GarageFull[idx])
		if err != nil {
			log.Fatalf("conversion of garage fullness from string to integer errored: %v", err)
		}
		garageStatus := GarageRecord{Name: GarageName[idx], Fullness: fullint, Addr: GarageAddr[idx], Timestamp: timestamp}
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
