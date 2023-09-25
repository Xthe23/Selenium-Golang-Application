package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tebeka/selenium"
)

func main() {
	// Set the path to the Chrome binary.
	chromeBinaryPath := os.Getenv("CHROME_BINARY_PATH") // This is the path to the Chrome executable, which I set in my ~/.zshrc file.
	caps := selenium.Capabilities{
		"browserName": "chrome",
		"goog:chromeOptions": map[string]interface{}{
			"binary": chromeBinaryPath, // Specify the path to the Chrome binary here.
		},
	}

	// Start a Selenium WebDriver server instance.
	opts := []selenium.ServiceOption{}
	service, err := selenium.NewSeleniumService("chromedriver", 4444, opts...) // This is the path to the chromedriver executable, which I have set in my ~/.zshrc file.
	if err != nil {
		log.Fatalf("Error starting the Selenium service: %v", err)
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 4444))
	if err != nil {
		log.Fatalf("Error connecting to the WebDriver: %v", err)
	}
	defer wd.Quit()

	// Navigate to the URL.
	err = wd.Get("https://www.ebay.com/b/PCGS-Certified-Medieval-Coins/18466/bn_26719429?rt=nc&LH_BO=1&mag=1")
	if err != nil {
		log.Fatalf("Error navigating to the URL: %v", err)
	}
}
