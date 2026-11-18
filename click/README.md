# click

The program shows how to click an element with a selector. A selector is a
pattern that finds an element on a page. The program reads the reference page of
the Go package `time` from the local test site, so it needs no internet. It
opens three code examples and prints the first line of the function `main` of
each one.

Run the program from the root of the repository:

```sh
$ go run ./click
```

The program uses the CSS selector `#example-After summary` to click the summary
of the first example. The click opens the example and shows its code. The
program waits until the `textarea` of the example is visible, then reads its
value. It repeats these steps for the examples `After`, `Sleep` and `NewTimer`.
It prints one line for each example and prints the full code of `After` as
well.

The flag `-url` sets the full URL of the page to read. When the flag is empty,
the program starts the local test site and reads `/docs/time`. The selectors of
the program are written for that page. A live page can have a different layout,
so the same selectors can fail there.

```sh
$ go run ./click -url http://127.0.0.1:41233/docs/time
```

The flag `-v` prints the protocol messages between the program and the browser.
The flag `-visible` shows the browser window and leaves it open when the
program ends. The flag `-visible-on-terminal` draws the page in the terminal
with terminal graphics. Use `-terminal-fps` to change the number of frames that
the program draws each second. The flag `-visible-on-terminal` does not work
with `-v`. Read the [root README](../README.md) for more about the flags.
