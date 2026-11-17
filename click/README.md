# click

## About

The program opens three code examples on the page `/docs/time` of the local test site. It clicks on each example to show its code, and it prints the first line of the function `main` of each example. The examples are `After`, `Sleep`, and `NewTimer`. The program needs no internet.

A local test site is a web server that runs on your own computer. The program starts the server when it starts.

## Running

Run the program from the root of the repository:

```sh
$ go run ./click
```

The program starts the local test site and reads the page `/docs/time`. It prints the line count of each example and the first line of `main`. It prints the whole code of the example `After`.

## Flags

The flag `-url` sets the full URL of another page to read, for example a live site. The program then does not start the local test site. The selectors are written for the local site, so a live site can differ.

The flag `-v` prints the protocol messages.

The flag `-visible` shows the browser window and leaves it open.

The flag `-visible-on-terminal` draws the page in the terminal with terminal graphics.

Use `-v` to print the protocol messages and `-visible` to show the browser window and leave it open.
