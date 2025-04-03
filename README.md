# Mekko

> market ekkonomy, who knows?

## Overview

Mekko is a project with an unknown purpose, but it’s probably useful.

## Setup

### Option 1: Using Nix ❄️

Mekko is managed as a Nix flake for reproducible development.

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

#### Running Mekko

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
go build -o mekko . && ./mekko
```

---

## Code Formatting ✨

Mekko uses `treefmt` for consistent styling.

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

---
