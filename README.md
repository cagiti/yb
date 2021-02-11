[![Go Report Card](https://goreportcard.com/badge/github.com/cagiti/yb)](https://goreportcard.com/report/github.com/cagiti/yb)
[![Downloads](https://img.shields.io/github/downloads/cagiti/yb/total.svg)]()

# yb

A small CLI that helps manage and setup yubikeys for developers.

## Install

```
brew tap cagiti/tap
brew install yb
```

## Usage

List yubikeys
- `yb list`

**e.g.**

```
❯ yb list
+-----+------------------------------+
| REF | CARD                         |
+-----+------------------------------+
|   0 | Yubico YubiKey OTP+FIDO+CCID |
+-----+------------------------------+
```
