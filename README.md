# suan

Suan( 算 ) is a CLI tool to calculate given mathematical expression.
Currently it supports addition, substraction, multiplication, division, and
exponent operations including any of their combinations with parenthesis.

## Install

```
go install github.com/OldPanda/suan
```

## Usage

```
» suan "1 + 1"
2.000000
» suan "(3 + 4) * 5 - 2 * (3 + 9)"
11.00000
» suan "3 / 4 + 3 * 2 - 8^2 * 2"
-121.250000
```
