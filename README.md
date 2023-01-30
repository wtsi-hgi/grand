# Grand

[![GoDoc](https://godoc.org/github.com/wtsi-hgi/grand?status.svg)](http://godoc.org/github.com/wtsi-hgi/grand)

Grand is a Go random string generator.

It is not cryptographically secure. But it is fast and memory efficient. 

## Installation

`go get github.com/wtsi-hgi/grand`

## Usage

1. Generate your random string, given a length parameter `n`:

    ```go
    grand.String(32)
    // returns "qzrWbaoLTVpQoottZyPFfNOoMioXHRuF"
    ```

1. Generate random string from other character sets:

    ```go
    gen := grand.New(grand.CharSetBase62)
	gen.String(20)
    // returns "q3rWba2LTVpQ4ottZyPv"
    ```

## Concurrency

The return value from NewGenerator is not safe for concurrent use. Create a new
one in each goroutine.

## Credits

I claim no credit for the generation logic. It's originally from user icza in
https://stackoverflow.com/a/31832326/1161743.

This is a fork of github.com/ernsheong/grand that ensures the seed is unique to
us, is faster, and changes the API to have shorter function names.
