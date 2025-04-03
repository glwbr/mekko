# 🍹 Juice - Brazilian NFC-e Invoice Extraction Library

Juice is a Go library for retrieving and parsing Brazilian electronic invoices (NFC-e) from SEFAZ (State Treasury Department).
It provides a simple, reliable way to extract structured data from NFC-e invoices using their access keys. The library handles the complexities of interacting with the SEFAZ portal, including session management, captcha handling, form submission, and HTML parsing.

[![Go Reference](https://pkg.go.dev/badge/github.com/glwbr/juice.svg)](https://pkg.go.dev/github.com/glwbr/juice)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## 🧱 Architecture Overview

```mermaid
graph TD
    A[Juice CLI / Applications] -->|Request with Access Key| B[Juice Library]
    B --> C[HTTP Client]
    C -->|Initial Request| D[SEFAZ Portal]
    D -->|Response with Captcha| C
    C --> E[Captcha Handler]
    E -->|Store Captcha| F[captcha.png]
    F -->|User Views| G[Human]
    G -->|Solves Captcha| H[CLI Input]
    H -->|Captcha Solution| C
    C -->|Submit Form with Access Key & Captcha| D
    D -->|HTML Response| C
    C --> I[HTML Parser]
    I -->|Extract Data| J[Structured Invoice]
    J -->|JSON/CSV Output| A
```

## 🌟 Features

- **Complete Data Extraction**: Retrieve full invoice details including items, prices, merchant info, and tax data
- **Captcha Solutions**: Support for manual captcha solving, with extensibility for automated solutions
- **Structured Data**: Convert raw HTML responses into clean, well-structured Go objects for easy consumption
- **Error Handling**: Comprehensive error types for different failure scenarios
- **Session Management**: Automatic handling of cookies and session state
- **Dual Interface**: Use as both a Go library and CLI tool

## 🚀 Installation

```sh 
go get github.com/glwbr/juice
```
## 💻 Basic Usage
As a Library

```go
package main

import (
	"fmt"
	"github.com/glwbr/juice"
	"log"
)

func main() {
	client := juice.NewClient()

	// With automatic captcha handling
	invoice, err := client.GetInvoice("00000000000000000000000000000000000000000000")
	if err != nil {
		log.Fatalf("Failed to get invoice: %v", err)
	}

	fmt.Printf("Merchant: %s\nTotal: R$ %.2f\n",
		invoice.Merchant.Name,
		invoice.TotalValue)
}
```

## CLI Usage
```sh
# Get invoice data
juice fetch --key 00000000000000000000000000000000000000000000

# Get captcha image
juice captcha --output captcha.png
```

## 🛠 Error Handling
Juice provides typed errors for precise error handling:

```go
invoice, err := client.GetInvoice(accessKey)
if err != nil {
	switch err := err.(type) {
	case *juice.CaptchaError:
		// Handle captcha challenge
		solution := solveCaptcha(err.ImageURL)
		invoice, err = client.RetryWithCaptcha(solution)
	case *juice.NotFoundError:
		// Handle invalid access key
	case *juice.NetworkError:
		// Handle connection issues
	default:
		// Generic error handling
	}
}
```

## 🧑‍💻 Development

### With Nix (Recommended) ❄️

The project is managed as a Nix flake for reproducible development.

#### Prerequisites

- [Nix](https://nixos.org/download.html) installed with [flakes enabled](https://nixos.wiki/wiki/Flakes)

#### Development Environment

Enter the environment:

```sh
nix develop
```

**For direnv users** (auto-loads the environment):

```sh
echo 'use flake' > .envrc && direnv allow
```

#### Running Juice

```sh
nix run .
```

---

### Option 2: Without Nix 🛠️

For non-Nix users, manually install:

#### Prerequisites

- [Go](https://go.dev/dl/) (version specified in `go.mod`)
- Optional: `treefmt` for formatting (`https://github.com/numtide/treefmt`)

#### Build & Run

```sh
go build -o juice . && ./juice
```

---

## Code Formatting ✨

The project uses `treefmt` for consistent styling.

Format all code:

```sh
treefmt
```

Customize via `.treefmt.toml`. Example:

```toml
[formatter.alejandra]
command = "alejandra"
includes = ["*.nix"]
```

## 📌 To-Do

- [ ] Implement `juice.ParseInvoiceHTML(html)`
- [ ] Captcha solver (automated)
- [ ] Dockerization
- [ ] Add tests

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

When contributing:
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.
