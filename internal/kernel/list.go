package kernel

// TODO: refactor: this looks super ugly but it works since the beginning :P

import (
	"encoding/xml"
	"fmt"
	"io"
	"kman/pkg/logger"
	"kman/pkg/utils"
	"net/http"
	"regexp"
)

type ListCommand struct {
	logger *logger.Logger
	ctx    *KernelContext
}

// Ensure struct implements interface
var _ ICommand = (*ListCommand)(nil)

func (c *ListCommand) String() string {
	return "List"
}

func (c *ListCommand) Execute() error {
	kernels, err := fetchKernels()
	if err != nil {
		return err
	}

	for i, kernel := range kernels {
		fmt.Printf("[%d]: %s\n", i, kernel.Title)
	}

	n_kernels := len(kernels) - 1
	var selectedIndex int
	for {
		fmt.Printf("Please select a kernel to download (0-%d): ", n_kernels)
		_, err := fmt.Scanf("%d", &selectedIndex)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		if selectedIndex >= 0 && selectedIndex <= n_kernels {
			break
		} else {
			fmt.Printf("Please enter a valid number between 0 and %d.\n", n_kernels)
		}
	}

	selectedKernel := kernels[selectedIndex]

	err = validateKernel(selectedKernel)
	if err != nil {
		return err
	}

	if selectedKernel.PGPSignature != "" {
		c.ctx.signatureURL = selectedKernel.PGPSignature
	}

	c.ctx.tarballURL = selectedKernel.SourceTarball

	return nil
}

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

// Output struct after scraping kernels
type Kernel struct {
	Title         string
	ReleaseDate   string
	Version       string
	SourceTarball string
	PGPSignature  string
	Patch         string
	ChangeLog     string
}

func validateKernel(k Kernel) error {
	if k.PGPSignature == "" {
		if !utils.ConfirmAction("This kernel tarball does not have a PGP signature. Are you okay with this?") {
			return fmt.Errorf("user declined to install without PGP signature.")
		}
	}

	if k.SourceTarball == "" {
		return fmt.Errorf("source tarball not found for this kernel: %+v\n", k)
	}

	return nil
}

func fetchKernels() ([]Kernel, error) {
	resp, err := http.Get("https://www.kernel.org/feeds/kdist.xml")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch kernel feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var rss RSS
	if err := xml.Unmarshal(body, &rss); err != nil {
		return nil, fmt.Errorf("failed to unmarshal XML: %w", err)
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
