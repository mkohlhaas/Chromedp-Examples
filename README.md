# About chromedp examples

<p align="center">
  <img src="https://raw.githubusercontent.com/chromedp/logo/main/chromedp.svg" alt="chromedp logo" width="160">
</p>

This repository holds 46 example programs for [`chromedp`][1], a Go package that
drives Chrome through the Chrome DevTools Protocol. The package documentation
has small examples that are self-contained. These programs are larger. Only the
program `fast` needs the internet. The program `extension` needs the files of
uBlock Origin Lite on disk. The program `remote` needs a Chrome with a debugging
port, unless you use its flag `-start`.

[![Unit Tests][examples-ci-status]][examples-ci]
[![Go Reference][goref-chromedp-status]][goref-chromedp]
[![Releases][release-status]][releases]
[![Discord Discussion][discord-status]][discord]

The programs use the typed API of `chromedp` v0.23.1 and `cdproto` v0.160.1.
The file `docs/API.md` in the `chromedp` repository shows the old code and the
new code side by side. The file `docs/MIGRATION.md` in the same folder lists
every name that changed. See
[the decision for the port](docs/decisions/2026-10-03-the-programs-use-the-new-typed-api.md).

## Build and run

The module needs Go 1.27. The packages `chromedp`, `cdproto` and `termcast` need
only Go 1.25 or later. The module requires `chromedp` v0.23.1, `cdproto` v0.160.1
and `termcast` v0.3.0, and Go downloads them when you build a program. The
programs use the package `chromedp/remote`, which is part of `chromedp`, for the
flag `-visible`, which keeps the browser open, and for the `remote` program.
They use `termcast` for the flag `-visible-on-terminal`.

The tags of this repository follow the tags of `chromedp`. The tag `v0.23.0` holds
programs that use `chromedp` v0.23, and a tag such as `v0.23.1` is a later change in
this repository.

`chromedp` starts the browser. It finds Chrome or Chromium on the `PATH`. If
Chrome has another name, link it to the name `google-chrome` in a folder on the
`PATH`. Run a program from the root of this repository:

```sh
# run the program <prog>, for example go run ./click
$ go run ./<prog>

# build the program <prog>
$ go build -o /tmp/<prog> ./<prog> && /tmp/<prog>
```

Every program accepts the flag `-v`. It prints the protocol messages between
the program and the browser. Without `-v`, a program prints only its own output.
For example, `go run ./cookie -v` shows the protocol messages. Run
`go run ./<prog> -h` to see all flags of a program. Put the flags of a program
after its folder, as in `go run ./multi -out out <url>`.

Every program except `remote` accepts the flag `-visible`. It shows the browser
window instead of a headless browser, and it leaves the browser open when the
program ends. The program prints the websocket address and the profile directory
of the browser to the standard error. The browser stays running until you close
its window, and the profile directory stays on disk until you delete it. To show
the window with no flag, set the variable `CHROMEDP_VISIBLEWINDOW=1`. Any value
other than an empty string, `0` and `false` turns it on. On Linux, a visible
window needs `DISPLAY` or `WAYLAND_DISPLAY`. The `remote` program has
no `-visible` flag, because it attaches to a browser that you started yourself.
The flag `-visible` has no effect on the remote browser of `forecast -remote`.
The `tabs` program works a little differently with `-visible`. It keeps every tab
open and waits until you close the browser, and it does not leave the browser
running. Its flag `-tabs-in-one-window` opens the tabs in one window, as real
tabs, and switches between them with the protocol command `Target.activateTarget`.
Without that flag, each tab opens in a window of its own. Try both with
`-visible` to see the difference.

Some programs need arguments:

- `forecast` needs the flag `-q`, for example `go run ./forecast -q Jakarta`.
- `geoip` takes IP addresses, for example `go run ./geoip 8.8.8.8`.
- `multi` takes URLs. See [multi/README.md](multi/README.md).
- `remote` needs a running browser, or the flag `-start`. See
  [remote/README.md](remote/README.md).
- `extension` needs the flag `-ext` with the files of uBlock Origin Lite.
- `upload` uploads its own source file, so you can run it from any folder.

The programs `download_file`, `download_image`, `emulate`, `har`, `pdf`,
`pdfstream` and `screenshot` write files into the current directory. The
programs `fast` and `forecast` write a file only when you give the flag `-out`,
and `multi` does so only with its flag `-out`. The programs `pdfoptions` and
`screencast` write their files into a new temporary directory, or into the
directory of their flag `-out`.

## The programs

<!-- the following section is updated by running `go run gen.go` -->
<!-- START EXAMPLES -->
| Example                           | Description                                                                                                                              |
|-----------------------------------|------------------------------------------------------------------------------------------------------------------------------------------|
| [click](/click)                   | use a selector to click on an element                                                                                                    |
| [console](/console)               | read the console of a page and its uncaught exceptions with chromedp                                                                     |
| [contexts](/contexts)             | see the JavaScript contexts that a tab already has, with chromedp                                                                        |
| [cookie](/cookie)                 | set an HTTP cookie on requests                                                                                                           |
| [dialogs](/dialogs)               | handle the JavaScript dialogs of a page                                                                                                  |
| [distrust](/distrust)             | check that a browser rejects the certificates of a compromised certificate authority                                                     |
| [download_file](/download_file)   | do headless file downloads                                                                                                               |
| [download_image](/download_image) | do headless image downloads                                                                                                              |
| [dragdrop](/dragdrop)             | drag and drop with the actions DragAndDrop and DragAndDropXY                                                                             |
| [emulate](/emulate)               | emulate a specific device such as an iPhone                                                                                              |
| [eval](/eval)                     | evaluate JavaScript and retrieve the result                                                                                              |
| [eventsiter](/eventsiter)         | listen to the events of a page with iterators                                                                                            |
| [exposefunc](/exposefunc)         | call Go functions from a page with chromedp                                                                                              |
| [extension](/extension)           | install a browser extension, uBlock Origin Lite, into the browser that chromedp starts, and how to show that it blocks the ads of a page |
| [fast](/fast)                     | measure the speed of the internet connection and show the result in the terminal                                                         |
| [forecast](/forecast)             | render the weather forecast of a city in the terminal                                                                                    |
| [frames](/frames)                 | reach elements inside an iframe and inside a shadow root                                                                                 |
| [geoip](/geoip)                   | look up the location of an IP address and show its map in the terminal                                                                   |
| [har](/har)                       | generate a HAR file from the network events of a page                                                                                    |
| [headers](/headers)               | add extra HTTP headers to browser requests                                                                                               |
| [intercept](/intercept)           | block, mock and change the requests of a page with the Fetch domain                                                                      |
| [keys](/keys)                     | send key events to an element                                                                                                            |
| [latlon](/latlon)                 | retrieve the latitude and the longitude of a map from the URL of the page with the navigation events of the page                         |
| [logic](/logic)                   | combine actions and Go code in a function that reads a list from a page                                                                  |
| [multi](/multi)                   | use headless-shell and a container (Docker, Podman, other)                                                                               |
| [pdf](/pdf)                       | capture a PDF of a page                                                                                                                  |
| [pdfoptions](/pdfoptions)         | print a page to PDF files with different options of chromedp                                                                             |
| [pdfstream](/pdfstream)           | print a page to a PDF file with a stream                                                                                                 |
| [popups](/popups)                 | work with the popups of a page and with several targets                                                                                  |
| [proxy](/proxy)                   | authenticate to a proxy server that requires authentication                                                                              |
| [rawcall](/rawcall)               | send protocol commands that chromedp has no action for                                                                                   |
| [remote](/remote)                 | connect to an existing Chrome DevTools instance using a remote WebSocket URL                                                             |
| [screencast](/screencast)         | record the screen of a page as a series of JPEG images                                                                                   |
| [screenshot](/screenshot)         | take a screenshot of a specific element and of the entire browser viewport                                                               |
| [selectors](/selectors)           | choose the elements of a page with the typed selectors                                                                                   |
| [session](/session)               | save the state of a session and restore it in another browser                                                                            |
| [structeval](/structeval)         | evaluate JavaScript into typed Go values                                                                                                 |
| [submit](/submit)                 | fill out and submit a form                                                                                                               |
| [subtree](/subtree)               | populate and travel a subtree of the DOM                                                                                                 |
| [tabs](/tabs)                     | use several tabs of one browser                                                                                                          |
| [termcast](/termcast)             | stream the screen of the browser to the terminal with terminal graphics                                                                  |
| [text](/text)                     | extract text from a specific element                                                                                                     |
| [timeline](/timeline)             | see the events of a page of every type in the order in which Chrome sent them                                                            |
| [upload](/upload)                 | upload a file on a form                                                                                                                  |
| [visible](/visible)               | wait until an element is visible                                                                                                         |
| [workers](/workers)               | run many jobs at the same time in one browser with a pool of goroutines                                                                  |
<!-- END EXAMPLES -->

## A reading order for learning chromedp

The examples do not need to be read in alphabetic order. Start with `click/main.go`. It shows the setup that most programs repeat. The setup has flags, a local site, a browser context, a timeout, `Navigate`, `WaitVisible` and actions that read or change the page. A browser context holds the browser state for one program. A timeout stops a program that waits too long. Skip `termstream` and `-visible` on the first pass. That code draws the page in the terminal or leaves the browser open.

Read `text`, `eval` and `structeval` next. They show how to read page content and how JavaScript results become Go values. Then read `selectors` as a catalog of typed selectors. A typed selector names a way to find an element on the page. These five programs teach the main loop of `chromedp`.

After that, read the programs in this table. Run each program as you read it. Add `-v` to see the protocol messages. The expected output of the offline programs is in `AGENTS.md`.

| Step | Read | What it teaches | First command |
| --- | --- | --- | --- |
| 1 | `README.md`, then `click/main.go` | The shared setup: flags, local site, browser context, timeout, `Navigate`, `WaitVisible` and actions | `go run ./click` |
| 2 | `text`, `eval`, `structeval` | Text extraction, JavaScript evaluation and typed Go results | `go run ./text` |
| 3 | `selectors` | CSS, ID, XPath, `JSPath`, `NodeIDs`, `ByFunc` and node options | `go run ./selectors` |
| 4 | `visible`, `keys`, `dialogs`, `dragdrop`, `submit`, `upload`, `logic` | Waiting, input, dialogs, drag and drop, forms, file upload and Go loops around browser actions | `go run ./logic` |
| 5 | `frames`, `subtree`, `contexts`, `workers` | Iframes, shadow roots, JavaScript contexts and concurrent jobs | `go run ./frames` |
| 6 | `console`, `eventsiter`, `timeline`, `latlon` | Browser events, their order and navigation events | `go run ./console` |
| 7 | `headers`, `cookie`, `intercept`, `har`, `rawcall` | Requests, responses and raw protocol commands | `go run ./intercept` |
| 8 | `emulate`, `session`, `popups`, `tabs` | Device state, saved state and several targets | `go run ./session` |
| 9 | `screenshot`, `pdf`, `pdfoptions`, `pdfstream`, `screencast`, `download_file`, `download_image` | Images, PDF files and downloads | `go run ./screenshot` |
| 10 | `exposefunc`, `proxy`, `remote`, `extension`, `distrust`, `fast`, `forecast`, `geoip`, `termcast`, `multi` | Special setup. Read these only when the task needs them | `go run ./exposefunc` |

`multi` needs a container tool. `fast`, `forecast`, `geoip` and `termcast` need a terminal that shows images. `distrust` needs Podman or Docker and reads a live test page.

Plan on about 30 hours to read and run all 46 examples with care. You can reach a useful understanding of `chromedp` after about 12 to 15 hours, because the first 15 programs teach most of the library.

The programs hold about 10,800 lines of Go. Much of that is repeated setup for flags, the local test site, the browser context, the timeout and terminal graphics. After you know that block in `click/main.go`, you can skim it in the other programs. The real learning is in the action calls, selectors, events, network commands and target handling.

The first 10 examples take about 4 to 6 hours. Run each one and compare the output with `AGENTS.md`. The next 15 take about 6 to 8 hours. They cover frames, events, network commands and browser state. The last 21 take about 12 to 16 hours. Many of them teach extra setup for images, files, containers, remote browsers and extensions. At 90 minutes a day, the full path takes about three weeks.

The programs `fast`, `forecast`, `geoip` and `remote` draw an image in the
terminal. Run them in a terminal that can show images. Most other programs have
the flag `-visible-on-terminal`, which the section [Draw the page in the
terminal](#draw-the-page-in-the-terminal) describes.

## The calls in the click program

The program `click` drives one Chrome tab. It opens a documentation page,
clicks three closed code examples and reads the code of each one. These are
the `chromedp` calls that it uses, in the order the program meets them.

An option of type `chromedp.ContextOption` changes how the browser starts.

`chromedp.ContextOption` is the type `func(*Context)`. It is a small setter
that receives the chromedp context before Chrome starts and changes one field,
such as whether a real window opens, whether protocol messages print, or
whether the browser stays alive after the program ends. `chromedp.NewContext`
applies every option it receives when it builds the context.

`chromedp.NewContext` prepares the context and, if no allocator exists yet,
sets one up. It does not start Chrome or open a tab yet. That happens on the
first `Run`, which starts Chrome and creates the first tab. `NewContext`
returns a context plus a cancel function. The cancel function closes the tab
when the program ends. `chromedp.WithDebugf(log.Printf)`
adds the option that prints each protocol message, so that a person can watch
the talk between the program and Chrome. The messages go to the terminal where
you run the program, on the standard error. They are not the messages that the
page prints to its own browser console. `chromedp.WithVisibleWindow()` adds
the option for a real window. `remote.WithKeepOpen()` comes from the
`chromedp/remote` package. It adds the option that leaves the browser alive
after the program exits, and it connects over a websocket because the pipe dies
with the Go process. `chromedp.KeptOpen(ctx)` then reports the websocket address
and the profile directory of that kept browser, and the program prints both to
the standard error.

`chromedp.Navigate(urlstr)` opens the page in the tab. `chromedp.WaitVisible`
waits until the element `body > footer` is on the page and visible, and the
program treats that as the sign that the page finished loading.
`chromedp.Do` runs a list of actions in order and stops at the first error.

`chromedp.Click(selector, chromedp.NodeVisible)` sends a real mouse click to
the first element that matches the selector. A selector is a CSS pattern, so
`#example-After summary` means the summary element inside the details element
whose id is `example-After`.

The backticks in `` `#example-` `` make it a Go raw string literal. It is just
the plain string `#example-`. Double quotes work the same here. The
program concatenates it with the example name to build a CSS selector, so
`#example-After summary` means "the summary element inside the element whose id
is `example-After`". The `#` is a CSS ID selector.

`chromedp.NodeVisible` is a query option that
makes the click wait until that element exists and is visible. The second
`chromedp.WaitVisible` call waits until the click opened the block and the
textarea with the code is visible.

`chromedp.Value(selector)` reads the value of a form field, and the example code
sits in the value of the textarea. `chromedp.Run` runs one action and returns
its result, so the program gets the code as a string. The typed API carries the
result type in the action. `Value` is an action that returns a string, and
`Navigate`, `Click` and `WaitVisible` return nothing and go through `Do`.

`chromedp.Do` runs several actions in order and returns only an error.
`chromedp.Run` runs one action and returns its result. Use `Do` when you do not
need a value back. Use `Run` when you need the result. `Do` stops at the first
error.

The remaining calls support the run but are not `chromedp`.
`termstream.Register` adds the flags `-visible-on-terminal` and
`-terminal-fps`, and `tc.Start` starts a stream that draws frames of the page
in the terminal. `s.LogWriter` sends the log lines through that stream.
`termstream.Fatal` stops the stream before it calls `log.Fatal`, because
`log.Fatal` skips deferred calls. `testsite.New` starts the local test site,
and `context.WithTimeout` gives the whole run 15 seconds.

## What does chromedp.WaitVisible do ?

`chromedp.WaitVisible` waits until the element matching a selector is visible in the browser.

It is an element query action. It calls `Query` with the `NodeVisible` option added. `NodeVisible` waits until the browser has sent all elements that match the selector. For each matched element, it asks Chrome for the box model. If Chrome cannot compute a box model, the element is treated as not visible.

It returns `Action[Void]`. It produces no value, it just waits. If the element never becomes visible within the context timeout, the action returns an error.

In `click/main.go`, `WaitVisible` is used twice:

- `chromedp.WaitVisible(`body > footer`)` waits until the page footer is visible. The program treats that as the sign that the page finished loading.
- `chromedp.WaitVisible(details+` textarea`)` waits until a click on an example's summary opened the block and the textarea holding the code is visible.

There are sibling actions: `WaitReady` waits until an element is loaded, `WaitNotVisible` waits until an element is hidden, `WaitEnabled` waits until an element is not disabled, and `WaitSelected` waits until an element is selected.

## What does chromedp.NodeVisible do ?

`chromedp.NodeVisible` is a query option, not an action by itself. It waits until
the browser has sent all elements matching the selector and they are visible.

It performs two checks on each matched element:

1. It asks Chrome for the box model (`dom.GetBoxModel`). If Chrome cannot compute
   a box model, the element is treated as not visible.
2. It runs a JavaScript snippet against the element to check CSS visibility. If
   the element is hidden by CSS, the option returns an error.

If every element passes both checks, the query succeeds. If any element never
becomes visible within the context timeout, the action returns an error.

Query options configure what an element query checks before it succeeds. Similar
options include:

| Option | What it waits for |
| --- | --- |
| `NodeReady` | Browser has sent all queried element nodes |
| `NodeNotVisible` | Browser has sent all nodes and they are not visible |
| `NodeEnabled` | Nodes exist and have no `disabled` attribute |
| `NodeSelected` | Nodes exist and have a `selected` attribute |
| `NodeNotPresent` | No element matches the query |
| `AtLeast(n)` | At least `n` elements match |
| `RetryInterval(interval)` | How often to retry when selection fails |
| `After(f)` | Runs a function once nodes match and the condition is true |

Matching actions use these options internally:

| Action | Uses |
| --- | --- |
| `WaitVisible` | `NodeVisible` |
| `WaitReady` | `NodeReady` |
| `WaitNotVisible` | `NodeNotVisible` |
| `WaitEnabled` | `NodeEnabled` |
| `WaitSelected` | `NodeSelected` |
| `WaitNotPresent` | `NodeNotPresent` |

In `click/main.go`, `chromedp.NodeVisible` is passed to `chromedp.Click`, so the
click waits until the summary element exists and is visible before clicking it.
And `chromedp.WaitVisible(details+` textarea`)` uses `NodeVisible` implicitly to
wait until the opened example's code textarea is visible.

## What does chromedp.Value do ?

`chromedp.Value(selector)` reads the value of a form element. The selector only
needs to find the element. The element itself must support the built-in `value`
property.

You do not have to add anything to the HTML. The `value` property exists
automatically on form elements such as `textarea`, `input`, `select`, `button`
and `option`. The browser sets and maintains it by the HTML standard.

A CSS class named `.value` is just a class like `.button` or `.highlight`. It
styles an element. It does not give the element a value or make it readable. If
you add `class="value"` to a plain `div`, the div gets a class named
`value`. It does not gain a readable value.

In `click/main.go`, `chromedp.Value(details+` textarea`)` reads the code from
the textarea of the opened example.

## Options for starting a browser

`chromedp.ContextOption` is the type `func(*Context)`. It is a small setter
that receives the chromedp context before Chrome starts and changes one field,
such as whether a real window opens, whether protocol messages print, or
whether the browser stays alive after the program ends. `chromedp.NewContext`
applies every option it receives when it builds the context.

These functions from the `chromedp` package return a `chromedp.ContextOption`:

| Function | What it does |
| --- | --- |
| `chromedp.WithVisibleWindow()` | Opens a real browser window instead of running headless. |
| `chromedp.WithAllocatorOptions(opts ...)` | Passes options to the default allocator that starts Chrome. |
| `chromedp.WithNewWindow(bool)` | Chooses whether a new tab opens in its own window or as a real tab in the browser window. |
| `chromedp.WithDetachOnCancel()` | Leaves the tab open when the context is canceled instead of closing it. |
| `chromedp.WithTargetID(id)` | Attaches to an existing target, such as a tab or an iframe, instead of creating a new tab. |
| `chromedp.WithListener(fn)` | Calls `fn` for every event of the target from the first message. |
| `chromedp.WithNewBrowserContext(...)` | Makes a context create a new BrowserContext and a new target in it. |
| `chromedp.WithExistingBrowserContext(id)` | Creates a new target in an existing BrowserContext. |
| `chromedp.WithLogf(f)` | Sets a logger for general log messages. |
| `chromedp.WithErrorf(f)` | Sets a logger for error messages. |
| `chromedp.WithDebugf(f)` | Sets a logger for protocol messages between the program and Chrome. |
| `chromedp.WithBrowserOption(opts ...)` | Passes browser options to the allocator. |

This function from the `chromedp/remote` package also returns a
`chromedp.ContextOption`:

| Function | What it does |
| --- | --- |
| `remote.WithKeepOpen()` | Leaves the browser alive after the program exits. It connects over a websocket. |

All of these are passed to `chromedp.NewContext(context.Background(), opts...)`.
Each one is just a `func(*Context)` that sets one field on the context before
Chrome starts. An option that tells chromedp to build the allocator one way has
no effect if you already built your own allocator or if the context inherits a
browser from a parent context.

## What does chromedp.KeptOpen do ?

`chromedp.KeptOpen(ctx)` reports the connection information of a browser that
was started with `remote.WithKeepOpen()`. It returns two strings:

1. The websocket address of the browser. Another program can attach to that
   address using the allocator of `chromedp/remote`.
2. The user data directory of the browser, that is, Chrome's profile directory
   on disk. You can delete it after you close the browser.

`KeptOpen` returns two empty strings when the context holds no browser, or when
the browser was not started with `remote.WithKeepOpen()`. It does not start or
stop anything. It only reads the fields stored in the context.

The browser starts at the first `Run`, not at `NewContext`, so call
`KeptOpen` after the first `Run`. The program `click` does that in a deferred
call and prints both values to the standard error:

```go
wsURL, dir := chromedp.KeptOpen(ctx)
fmt.Fprintf(os.Stderr, "browser kept open at %s with profile directory %s\n", wsURL, dir)
```

This tells you where the still-running browser is, so you can keep inspecting
the page after the program ends. To close a kept browser, close its window or
kill its process.

`KeptOpen` does not collide with `testsite.Site.Close()`. They deal with
different things:

- `site.Close()` stops the local test HTTP server that serves pages such as
  `/docs/time`.
- `chromedp.KeptOpen(ctx)` reports the address and profile directory of the
  Chrome browser started by chromedp.

You can close the test site while the browser stays open, or close the browser
while the test site keeps running. `click` uses both together: `defer
site.Close()` shuts down the page server, and the deferred `KeptOpen` call
reports the kept browser.

## The local test site

The package `internal/testsite` serves a small website from your own machine.
It runs an HTTP server on a free port of `127.0.0.1`, and it needs no internet.
The example programs read this site instead of a live site, because a live site
can change, can vanish and can need a network.

The site holds long pages that look real, so a program has something real to
read, click, fill, screenshot or print. The pages are built in Go from fixed
seeds, so every run serves the same text and the same pictures. That is why the
offline programs have a fixed output that you can compare.

`testsite.New()` starts the server and returns a `*Site`. The field `URL` holds
the base address. The field `OtherURL` holds the same server under the name
`localhost`, which a browser treats as another origin. `Close` stops the
server. A program that wants its own listener uses `testsite.Handler`. A
program that reads the site takes a flag `-url`, and when that flag is empty
the program starts this site.

You can run the site for yourself with `go run ./internal/testsite/serve`. It
prints the two addresses and keeps the site running until you press Ctrl-C. A
browser or another program can then read the pages. The file
[internal/testsite/README.md](internal/testsite/README.md) lists every route.

The routes cover the pages that the programs need. The page `/docs/time` looks
like package documentation and holds fourteen closed code examples, and the
program `click` opens three of them. There are encyclopedia pages with a search form,
a repository list with a download, a gallery of images, forecasts for eight
cities, an IP lookup and a map with tiles. Other pages show the user agent and
the viewport, a report for printing and an animated SVG of a harbour. The news
pages are the one exception to the rule that a page loads nothing from another
host, because the program `extension` needs real ad scripts to block.

The package has its own tests, and they need no browser. They load each page
with an HTTP client and check the status, the landmarks that the programs use,
the sizes and the unique ids. They also check that no page refers to another
host, except the news pages. The binary files come from
`internal/testsite/gen/main.go`, and the repository commits them. The file
[internal/testsite/README.md](internal/testsite/README.md) lists every route,
every selector and the size of each page.

## Draw the page in the terminal

Most programs take the flag `-visible-on-terminal`. It draws the page of the
program in the terminal while the program runs, with terminal graphics (Kitty,
iTerm2 or Sixel). The flag also works with a headless browser over `ssh`. The
package [`termcast`][termcast] does this work. It uses the screencast of the
Chrome DevTools Protocol and 4 frames each second by default. Use the flag
`-terminal-fps` to change the rate.

```
go run github.com/chromedp/examples/click@latest -visible-on-terminal
```

In kitty and Ghostty, the stream changes the picture with no flicker. In other
terminals, it clears the terminal at each redraw. It holds the log lines and
the results of the program, and prints them after the final frame. The flag does not
work with `-v`, and it stops with an error when the terminal has no graphics.
The programs `tabs`, `popups`, `workers`, `multi`, `session`, `screencast`,
`pdfstream`, `har` and `rawcall` do not have the flag. The program `termcast`
has the flag `-fps` and not the two flags. It shows how to use the package in
your own code, with no helper of this repository. The other programs define the
two flags with the package `internal/termstream`. It plays an animated SVG and streams it to the terminal.

`internal/termstream` is a small package in this repository, and it is not part
of `chromedp`. It exists so that all programs offer the same two flags and make
the same calls around the stream. The package `termcast` has no flags of its
own, so `Register` adds `-visible-on-terminal`, which turns the stream on, and
`-terminal-fps`, which sets the frame rate. `Start` returns a nil stream and no
error when the flag is off, so a program calls it without a condition. A nil
stream is safe to use, and every method on it does nothing. `Start` also refuses
to run when `-v` is on, because the protocol messages and the frames both write
to the terminal and the messages draw over the picture.

`Stop` ends the stream and draws a final frame. `LogWriter` gives the program a
writer for the log lines that the stream holds until the final frame, so a log
line never smears the picture. `Fatal` stops the stream and then calls
`log.Fatal`, because `log.Fatal` exits at once and skips deferred calls. Without
that order the terminal keeps the final frame and loses the log lines.

`Draw` writes one still image in the terminal. It takes the terminal that the
stream found, or it detects the terminal itself when the stream is off. A
program that makes a picture rather than shows the page live uses it, for
example `forecast`.

## termstream is a helper for this repository

`termstream` is helpful in this repository, not a general chromedp debugging
tool. It is an internal package of this module, so programs outside this
repository cannot import it.

It exists so that every example program offers the same two flags,
`-visible-on-terminal` and `-terminal-fps`, and makes the same calls around a
terminal stream of the page. The actual drawing work is done by the external
package `termcast`, which uses the Chrome DevTools Protocol screencast to send
pictures of the page to a terminal that supports Kitty, iTerm2 or Sixel
graphics.

For debugging, `termstream` is valuable when you want to see what the browser
is showing while a program runs, especially over `ssh` where a real browser
window is not available. But it is not a replacement for chromedp's own
debugging tools. The flag `-v` prints the protocol messages between the
program and Chrome, which is the better tool for understanding what chromedp
sends and receives. The flag `-visible` shows a real browser window, which
works on any machine with a display. `termstream` only works on terminals that
can render graphics, and it does not work with `-v`, because the protocol
messages and the frames both write to the terminal.

So `termstream` is a convenient visual aid for the examples in this
repository, but chromedp's `-v` and `-visible` are the broader debugging
tools.

## What does log.SetOutput(s.LogWriter()) do ?

`log.SetOutput(s.LogWriter())` redirects Go's standard logger so that every
call to `log.Printf` writes to the stream's log writer instead of directly to
standard error.

While the stream runs, the writer holds log lines in memory. They are not
printed immediately, so they do not smear over the page frames drawn in the
terminal. When `s.Stop` runs, the held lines print after the final frame.
After `Stop`, each line passes through at once to the original log
destination. If the stream is nil, `LogWriter` returns `os.Stderr`, so logs
print normally.

In `click/main.go`, the program calls `log.SetOutput(s.LogWriter())` right
after `tc.Start`. Every `log.Printf` in the program is buffered and printed
cleanly when the stream stops. This includes the lines that report each
example's code. The lines are not interleaved with the terminal graphics.

## The most used chromedp functions

Across the 46 example programs, these are the `chromedp` functions that appear
most often:

| Group | Function | What it does |
| --- | --- | --- |
| Start | `chromedp.NewContext` | Creates the browser context and starts Chrome. Every program calls it. |
| Start | `chromedp.WithDebugf` | Prints the protocol messages between the program and Chrome. |
| Start | `chromedp.WithVisibleWindow` | Opens a real browser window instead of running headless. |
| Run | `chromedp.Run` | Runs one action and returns its result. The most called function. |
| Run | `chromedp.Do` | Runs several actions in order and stops at the first error. |
| Page | `chromedp.Navigate` | Opens a page in the tab. |
| Page | `chromedp.WaitVisible` | Waits until an element is visible. |
| Page | `chromedp.WaitReady` | Waits until an element is loaded. |
| Click | `chromedp.Click` | Sends a mouse click to an element. |
| Read | `chromedp.Value` | Reads the value of a form field. |
| Read | `chromedp.Text` | Reads the text of an element. |
| Read | `chromedp.Nodes` | Reads multiple matching elements. |
| Find | `chromedp.CSS` | Finds elements by CSS selector. |
| Find | `chromedp.ID` | Finds an element by its HTML id. |
| JavaScript | `chromedp.Evaluate` | Runs JavaScript in the page and returns the result. |
| JavaScript | `chromedp.Call` | Sends a raw Chrome DevTools Protocol command. |
| Events | `chromedp.Listen` | Listens for browser events. |
| State | `chromedp.Target` | Gets the browser target of a context. |
| State | `chromedp.WithTargetID` | Attaches to an existing tab instead of creating a new one. |

`NewContext`, `Run`, `Do`, `Navigate` and `WaitVisible` appear in nearly every
program. They are the core of `chromedp`.

## chromedp is easy to use

chromedp is straightforward once you know the core loop. You start a
context. You run actions such as Navigate, Click and WaitVisible. You read
results with Run or Value. The typed API carries the result type in each action.
You do not have to guess what a call returns.

The examples repeat the same setup block. After you know that block, the real
work is in the selectors, actions and events.

## Verification

The table says which programs run with the current API, and what each program
needs. "Offline" means that the program needs no internet and no service. A
program that needs a live site reads it from the internet. A program that needs
a terminal draws an image with `rasterm`. Every program runs in a headless
Chrome unless you set the flag `-visible`, so no program in the table needs a
display.

The column "Checked" says how the result was found. "Run on 2026-10-04" means that
the offline programs ran with `go run`, with Chrome 154 on Linux, and finished
with exit code 0. "Earlier test" means a run of the same code with Chrome 154 on
Linux before that date. These programs were not run again on 2026-10-04,
because they need a live site that can change. No program was run with the flag
`-visible`.

| Example         | Needs                            | Result with the current API                       | Checked           |
|-----------------|----------------------------------|---------------------------------------------------|-------------------|
| click           | offline                          | works                                             | run on 2026-10-04 |
| console         | offline                          | works                                             | run on 2026-10-04 |
| contexts        | offline                          | works                                             | run on 2026-11-06 |
| cookie          | offline                          | works                                             | run on 2026-10-04 |
| dialogs         | offline                          | works                                             | run on 2026-10-04 |
| distrust        | internet, Podman                 | works, the browser connected to 7 of 14 hosts     | run on 2026-10-26 |
| download_file   | offline                          | works                                             | run on 2026-10-04 |
| download_image  | offline                          | works                                             | run on 2026-10-04 |
| dragdrop        | offline                          | works                                             | run on 2026-10-04 |
| emulate         | offline                          | works                                             | run on 2026-10-04 |
| eval            | offline                          | works                                             | run on 2026-10-04 |
| eventsiter      | offline                          | works                                             | run on 2026-10-04 |
| exposefunc      | offline                          | works                                             | run on 2026-10-04 |
| extension       | offline, the uBlock Origin Lite files on disk | works                          | run on 2026-10-04 |
| fast            | internet (fast.com), terminal    | needs a terminal that shows images                | earlier test |
| forecast        | terminal                         | works in a terminal, stops without one            | run on 2026-10-04 |
| frames          | offline                          | works                                             | run on 2026-10-04 |
| geoip           | terminal                         | works in a terminal, stops without one            | run on 2026-10-04 |
| har             | offline                          | works, now with chromedp.Listen                   | run on 2026-11-04 |
| headers         | offline                          | works                                             | run on 2026-10-04 |
| intercept       | offline                          | works                                             | run on 2026-10-04 |
| keys            | offline                          | works                                             | run on 2026-10-04 |
| latlon          | offline                          | works                                             | run on 2026-10-04 |
| logic           | offline                          | works                                             | run on 2026-10-04 |
| multi           | offline                          | works                                             | run on 2026-10-04 |
| pdf             | offline                          | works                                             | run on 2026-10-04 |
| pdfoptions      | offline                          | works                                             | run on 2026-10-04 |
| pdfstream       | offline                          | works                                             | run on 2026-10-04 |
| popups          | offline                          | works                                             | run on 2026-10-04 |
| proxy           | offline                          | works                                             | run on 2026-10-04 |
| rawcall         | offline                          | works                                             | run on 2026-10-04 |
| remote          | a Chrome with a debugging port, or -start | works up to the terminal image                    | earlier test      |
| screencast      | offline                          | works                                             | run on 2026-10-04 |
| screenshot      | offline                          | works                                             | run on 2026-10-04 |
| selectors       | offline                          | works                                             | run on 2026-10-04 |
| session         | offline                          | works                                             | run on 2026-10-04 |
| structeval      | offline                          | works                                             | run on 2026-10-04 |
| submit          | offline                          | works                                             | run on 2026-10-04 |
| subtree         | offline                          | works                                             | run on 2026-10-04 |
| tabs            | offline                          | works                                             | run on 2026-10-04 |
| termcast        | offline, a terminal              | works in a terminal, stops without one            | run on 2026-10-04 |
| text            | offline                          | works                                             | run on 2026-10-04 |
| timeline        | offline                          | works                                             | run on 2026-11-04 |
| upload          | offline                          | works                                             | run on 2026-10-04 |
| visible         | offline                          | works                                             | earlier test      |
| workers         | offline                          | works                                             | run on 2026-10-04 |

Notes:

1. The programs `forecast` and `geoip` draw an image with `rasterm`. Without a
   terminal that shows images, they stop with the error `term graphics not
   available`. `fast` does the same, and it reads the live site `fast.com`.
2. The programs `geoip` and `forecast` embed their data files, `GeoLite2-City.mmdb`
   and `hl.json`, so they run from any folder.
3. `remote` needs a Chrome that runs with a debugging port, for example
   `chrome --headless --remote-debugging-port=9222`. Use the flag `-start` to
   let `remote` start that Chrome.
4. Every other program except `distrust` reads local content. These programs read the local test
   site in `internal/testsite`, and they take the flag `-url` to read another
   site: `click`, `download_file`, `download_image`, `emulate`, `eval`,
   `extension`, `forecast`, `geoip`, `latlon`, `logic`, `pdf`, `screenshot`,
   `submit`, `termcast` and `text`. The program `remote` takes the flag `-nav`.
   The other programs serve their own page from a local server.
5. `extension` needs the files of uBlock Origin Lite on disk. Use the flag
   `-ext`.
6. `distrust` starts the image `headless-shell` in a container with Podman,
   and reads the test page `emudhra.we-have-a-problem.com` and its test hosts.
   Use the flag `-podman` to name another command, such as `docker`. On
   2026-10-26 the image `latest` was Chrome 157.0.8103.15. The browser connected
   to the 7 hosts that chain to the root of eMudhra, so it still trusted the
   compromised intermediate certificates, and the program ended with the
   status 1. Chrome blocks these certificates with an update of its CRLSet,
   and the image does not receive that update.

## Live sites can change

`fast` reads the live site `fast.com`, and `distrust` reads the test page
`emudhra.we-have-a-problem.com`. Their results can be different on another day.

## Questions, bugs and changes

Ask questions in the [GitHub Discussions of `chromedp`][2]. Report a bug as an
[issue of `chromedp`][3]. Do not open an issue for a question. Pull requests
are welcome, and `CONTRIBUTING.md` says how to prepare one. A new program must
keep to the rules in `AGENTS.md`.

| Document | Holds |
| --- | --- |
| [AGENTS.md](AGENTS.md) | the rules, the layout and the commands for a coding agent and for a person |
| [CONTRIBUTING.md](CONTRIBUTING.md) | how to prepare a change |
| [docs/decisions/README.md](docs/decisions/README.md) | the index of the decisions |

[1]: https://github.com/chromedp/chromedp
[termcast]: https://github.com/chromedp/termcast
[2]: https://github.com/chromedp/chromedp/discussions
[3]: https://github.com/chromedp/chromedp/issues
[examples-ci]: https://github.com/chromedp/examples/actions/workflows/test.yml (Test CI)
[examples-ci-status]: https://github.com/chromedp/examples/actions/workflows/test.yml/badge.svg (Test CI)
[goref-chromedp]: https://pkg.go.dev/github.com/chromedp/chromedp
[goref-chromedp-status]: https://pkg.go.dev/badge/github.com/chromedp/chromedp.svg
[release-status]: https://img.shields.io/github/v/release/chromedp/examples?display_name=tag&sort=semver (Latest Release)
[releases]: https://github.com/chromedp/examples/releases (Releases)
[discord]: https://discord.gg/WDWAgXwJqN "Discord Discussion"
[discord-status]: https://img.shields.io/discord/829150509658013727.svg?label=Discord&logo=Discord&colorB=7289da&style=flat-square "Discord Discussion"
