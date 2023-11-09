# Brazilian Phone Number Generator

This Go package provides a simple way to generate random Brazilian phone numbers, both landline and mobile, including the ability to format them according to the international E.164 standard.

## Features

- Generate random landline phone numbers with valid area codes.
- Generate random mobile phone numbers, prefixed with '9' as per Brazilian standards.
- Format generated phone numbers in the E.164 international format.
- Generate a specified number of phone numbers at once.

## How It Works

The package contains a `PhoneGen` struct that provides methods to generate random phone numbers. It uses predefined area codes and number patterns to ensure that the generated numbers are plausible for Brazilian phone numbers.

### Generating Phone Numbers

To generate phone numbers, create an instance of `PhoneGen` and call one of the following methods:

- `Random(limit int)`: Generates a mix of random landline and mobile numbers.
- `RandomMobile(limit int)`: Generates random mobile phone numbers.
- `RandomLandline(limit int)`: Generates random landline phone numbers.
- `RandomE164(limit int, countryCode string)`: Generates random phone numbers formatted in E.164 format.

### Example Usage

```go
package main

import (
    "fmt"
    "github.com/thiagozs/go-phonegen"
)

func main() {
    gen := phonegen.New()

    // Generate 5 random mobile numbers
    mobileNumbers := gen.RandomMobile(5)
    fmt.Println("Mobile Numbers:", mobileNumbers)

    // Generate 5 random landline numbers
    landlineNumbers := gen.RandomLandline(5)
    fmt.Println("Landline Numbers:", landlineNumbers)

    // Generate 5 random E.164 formatted numbers
    e164Numbers, err := gen.RandomE164(5, "55") // 55 is the country code for Brazil
    if err != nil {
        fmt.Println("Error generating E.164 numbers:", err)
        return
    }
    fmt.Println("E.164 Formatted Numbers:", e164Numbers)
}
```

### Testing

The package includes unit tests to verify the correctness of the generated phone numbers. Run the tests using the standard Go testing tools:

```bash
go test ./...
```

### Contributions

Contributions to this package are welcome. Please ensure that any pull requests include corresponding unit tests to verify the changes.

### License

This package is licensed under the MIT License. See the LICENSE file for details.

This `README.md` provides a brief overview of the package, its features, how it works, and how to use it. It also includes a section on testing and contributions to encourage community involvement.

2023, thiagozs
