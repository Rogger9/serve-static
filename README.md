# Serve

A simple command-line http server for local development.

\*_Required: [Go](https://go.dev/)_

## Installation

```bash
git clone ...
cd serve-static

go build
go install
```

## Usage

```bash
serve -d path/directory
```

In the directory where static files are located

```bash
serve
```

### Options

| Flag |               Description               | Default |
| :--: | :-------------------------------------: | :-----: |
| `-d` | Path. Directory where files are located |  `./`   |
| `-p` |               Port to use               | `8080`  |
