# suan

[![Go Reference](https://pkg.go.dev/badge/github.com/OldPanda/suan.svg)](https://pkg.go.dev/github.com/OldPanda/suan)
[![Go Report Card](https://goreportcard.com/badge/github.com/OldPanda/suan)](https://goreportcard.com/report/github.com/OldPanda/suan)

Suan( 算 ) is a CLI tool to calculate given mathematical expression.
Currently it supports addition, substraction, multiplication, division, and
exponent operations including any of their combinations with parenthesis.

## Install

```
go install github.com/OldPanda/suan
```

## Usage

```
» suan -h
Suan( 算 ) is a CLI tool to calculate given mathematical expression.
Currently it supports addtion, substraction, multiplication, division, and
exponent operations and any of their combinations with parenthesis.

Usage:
  suan [flags]

Flags:
  -h, --help      help for suan
  -v, --version   Version information
```

### Examples

```
» suan "1 + 1"
2.000000
» suan "(3 + 4) * 5 - 2 * (3 + 9)"
11.00000
» suan "3 / 4 + 3 * 2 - 8^2 * 2"
-121.250000
```
