// Command serve runs the local test site until you press Ctrl-C. It prints the
// two addresses of the site, so a browser or a program can read them. It is a
// development tool and not an example program for chromedp.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/chromedp/examples/internal/testsite"
)

func main() {
	site := testsite.New()
	defer site.Close()

	fmt.Println(site.URL)
	fmt.Println(site.OtherURL)
	fmt.Println("press Ctrl-C to stop the site")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
}
