# Nopol

[![Build Status](https://travis-ci.org/matriphe/nopol.svg?branch=master)](https://travis-ci.org/matriphe/nopol)

Nopol is Go package to check and format Indonesian police number (vehicle’s registration plate number).

It will format vehicle registration number format like this `{XX} {DDDD} {YYY}`, for example `AB 1234 XYZ` or `RI 1`.

For more info about this, please refer to the [Wikipedia](https://en.wikipedia.org/wiki/Vehicle_registration_plates_of_Indonesia).

## Installation

```shell
go get -u github.com/matriphe/nopol
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