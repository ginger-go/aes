# aes

This package provides AES encryption and decryption functionality for encrypting and decrypting text using a secret key. The package provides functions for encrypting and decrypting text, as well as encoding and decoding the encrypted text to and from base64 strings. The encryption algorithm used is the Advanced Encryption Standard (AES) with a 128-bit block size and a 128, 192, or 256-bit key size. The key size can be customized by providing a key of the desired length. Note that the package does not provide any key management or key exchange functionality, so the key must be securely managed and exchanged between the parties that will use it.

## Installation

```bash
go get -u github.com/ginger-go/aes
```

## Documentation

Please refer to [https://ginger-go.gitbook.io/ase/](https://ginger-go.gitbook.io/ase/) for the documentation.

## Perform tests

Run ginkgo tests with the following command:
```bash
ginkgo -r -v -p --cover --coverpkg=github.com/ginger-go/aes
```

Read the coverage report with the following command:
```bash
go tool cover -html=coverprofile.out
```
