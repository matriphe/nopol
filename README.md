# Nopol

[![Build Status](https://travis-ci.org/matriphe/nopol.svg?branch=master)](https://travis-ci.org/matriphe/nopol)

Nopol is Go package to check and format Indonesian police number (vehicle's license plate number) of vehicles.

It will format police number format like this `{XX} {DDDD} {YYY}`, for example `AB 1234 XYZ` or `RI 1`.

## Installation

```shell
go get github.com/matriphe/nopol
```

## Usage

```go
// Check validation
val := nopol.IsValid("AB 1234 XYZ") // return true
val := nopi.IsValid("ABC 12345") // return false

// Format
f, err := nopol.Format("AB1234-XYZ") // return AB 1234 XYZ, nil
f, err := nopol.Format("ABC12345") // return "", Not a valid police number
```