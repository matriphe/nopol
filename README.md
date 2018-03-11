# Nopol

[![Build Status](https://travis-ci.org/matriphe/nopol.svg?branch=master)](https://travis-ci.org/matriphe/nopol)

**Nopol** is Go package to check and format Indonesian vehicle registration number (police number).

It will format vehicle registration number format like this `{XX} {DDDD} {YYY}`, for example `AD 6742 DZ`, `CD 129` or `RI 1`.

For more info about Indonesian vehicle registration number, please refer to the [Wikipedia](https://en.wikipedia.org/wiki/Vehicle_registration_plates_of_Indonesia).

## Installation

```shell
go get -u github.com/matriphe/nopol
```

## Usage

```go
// Check validation
val := nopol.IsValid("AB 1234 XYZ") // return true
val := nopol.IsValid("ABC 12345") // return false

// Format
f, err := nopol.Format("AB1234-XYZ") // return AB 1234 XYZ, nil
f, err := nopol.Format("ABC12345") // return "", Not a valid police number
```

For more info, refer to [this documentation](https://godoc.org/github.com/matriphe/nopol).

## To Do

* Get information from the vehicle registration number
* Check validity of the vehicle registration number

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.