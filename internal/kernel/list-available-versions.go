package kernel

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

const (
	kernelFeed = "https://www.kernel.org/feeds/kdist.xml"
)

type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title string `xml:"title"`
	Items []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

// Output struct after fetching/listing kernels
type Kernel struct {
	Title         string
	ReleaseDate   string
	Version       string
	SourceTarball string
	PGPSignature  string
	Patch         string
	ChangeLog     string
}

func ListKernels() ([]Kernel, error) {
	kernels, err := fetchKernels()
	if err != nil {
		return nil, err
	}

	for i, kernel := range kernels {
		fmt.Printf("[%d]: %s\n", i, kernel.Title)
	}

	return kernels, nil
}

func fetchKernels() ([]Kernel, error) {
	resp, err := http.Get(kernelFeed)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rss RSS
	if err := xml.Unmarshal(body, &rss); err != nil {
		return nil, err
	}

	var kernels []Kernel
	for _, item := range rss.Channel.Items {
		version := extractLink(item.Description, "Version")
		source := extractLink(item.Description, "Source")
		pgpSignature := extractLink(item.Description, "PGP Signature")
		patch := extractLink(item.Description, "Patch")
		changeLog := extractLink(item.Description, "ChangeLog")

		if source == "" {
			continue
		}

		kernels = append(kernels, Kernel{
			Title:         item.Title,
			ReleaseDate:   item.PubDate,
			Version:       version,
			SourceTarball: source,
			PGPSignature:  pgpSignature,
			Patch:         patch,
			ChangeLog:     changeLog,
		})
	}

	return kernels, nil
}

func extractLink(description string, label string) string {
	re := regexp.MustCompile(`<tr><th align="right">` +
		regexp.QuoteMeta(label) + `:</th><td><a href="([^"]+)">`)
	match := re.FindStringSubmatch(description)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}
