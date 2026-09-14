# kryvora-cli

[![Go Version](https://img.shields.io/badge/go-%3E%3D%201.22-30363d.svg?style=flat-square&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-Apache--2.0-30363d.svg?style=flat-square)](LICENSE)
[![Release](https://img.shields.io/github/v/release/kryvora-network/kryvora-cli?style=flat-square&color=30363d&label=release)](https://github.com/kryvora-network/kryvora-cli/releases)
[![CI Status](https://img.shields.io/github/actions/workflow/status/kryvora-network/kryvora-cli/ci.yml?branch=main&style=flat-square&label=ci&color=30363d)](https://github.com/kryvora-network/kryvora-cli/actions)

Command line interface and diagnostic operator tool for Kryvora Network nodes.

## Overview

The `kryvora` CLI provides node operators with real time diagnostics, availability probing, sync verification, and identity inspection. It connects to a local or remote `kryvora-node` daemon over HTTP or Unix domain socket.

## Installation

### From Source

Requires Go 1.22 or higher.

```bash
git clone https://github.com/kryvora-network/kryvora-cli.git
cd kryvora-cli
go build -o kryvora .
sudo mv kryvora /usr/local/bin/
```

### Prebuilt Binaries

Prebuilt binaries for Linux (amd64, arm64) and macOS (Apple Silicon, Intel) are available on the [Releases](https://github.com/kryvora-network/kryvora-cli/releases) page.

## Commands

### Node Status

Query operational telemetry and peer synchronization status:

```bash
kryvora status
```

Example output:

```text
Kryvora Node Status
Status:     active
Version:    0.2.0
Uptime:     342000 seconds
Peers:      8 connected
Sync State: synced
```

Machine readable JSON format:

```bash
kryvora status --json
```

### Health Check

Verify daemon process availability:

```bash
kryvora health
```

### Node Identity

Inspect the local node identifier and ed25519 public key:

```bash
kryvora identity
```

## Options

| Flag | Default | Description |
|---|---|---|
| `--endpoint` | `http://127.0.0.1:4177` | Target daemon HTTP API endpoint |
| `--json` | `false` | Format command output as raw JSON |

## License

Apache License 2.0. See LICENSE for details.
