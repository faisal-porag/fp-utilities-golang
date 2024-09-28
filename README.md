# fp-utilities-golang - Utility Functions

## Overview

`fp-utilities-golang` is a comprehensive Go package that offers a collection of utility functions, 
featuring secure encoding and decoding mechanisms utilizing AES encryption. This package is designed
to encrypt and decrypt sensitive data, such as customer profile information, ensuring confidentiality and security.
By consolidating essential utility functions into a single package, `fp-utilities-golang` streamlines development and
enhances the robustness of applications that require secure data handling.

## Unit Testing

This package includes unit tests to ensure the functionality and reliability of the encoding and decoding methods.
The tests verify that the data can be encoded and decoded correctly, maintaining data integrity.

### Testing Framework

The unit tests are built using the Go testing framework, which is included in the Go standard library.

### Test Files

- `encode_decode_utils.go`: Contains the encoding and decoding utility functions.
- `encode_decode_utils_test.go`: Contains the unit tests for the utility functions.

### Running Tests

To run the unit tests for this package, follow these steps:

1. **Open a terminal** and navigate to the package directory:

   ```bash
   cd $GOPATH/src/github.com/faisal-porag/fp-utilities-golang
   ```
   
2. **Run the test using:**
   
   ```bash
   cd fp-utilities-golang
   go test
   ```
   
3. **Running Specific Tests**
   To run a specific test, you can use the -run flag followed by a regular expression that matches the test name:

    ```bash
   cd fp-utilities-golang
   go test -run TestSecureEncodeDecodeUtilsFunc
   ```

### How to use this package 

```bash
package main

import (
    "fmt"
    "log"
    fp-utils "github.com/faisal-porag/fp-utilities-golang"
)

func main() {
    key := "e2f8c35b9cd447b2e29f3738b54c4f2e"
    text := "example664.ex@example.com"

    // Use the SecureEncode function
    encoded, err := fp-utils.SecureEncodeUtilsFunc(text, key)
    if err != nil {
        log.Fatal("Encoding failed:", err)
    }
    fmt.Println("Encoded Text:", encoded)

    // Use the SecureDecode function
    decoded, err := fp-utils.SecureDecodeUtilsFunc(encoded, key)
    if err != nil {
        log.Fatal("Decoding failed:", err)
    }
    fmt.Println("Decoded Text:", decoded)
}
```



