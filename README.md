QR code generator for text terminals
======================================================================

Inspired by https://github.com/fumiyas/qrc, just outputs a smaller qrcode.

What's this?
---------------------------------------------------------------------

This program generates QR codes in text format for
text terminals.

Use case
---------------------------------------------------------------------

You can transfer data to smartphones with a QR code reader application
from your terminal.

Usage
---------------------------------------------------------------------

`qrcode` program takes a text from command-line argument or standard
input (if no command-line argument) and encodes it to a QR code.

```console
$ quote --help
Usage: quote [OPTIONS] [TEXT]

Options:
  -h, --help
    Show this help message
```

Download
---------------------------------------------------------------------

Binary files are here:

  * https://github.com/spazzymoto/qrcode/releases

Build from source codes
---------------------------------------------------------------------

If you have Go language environment, try the following:

```console
$ go get github.com/spazzymoto/qrcode
```