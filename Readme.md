# Domain Scan

Simple tool to look for domain names against a given pattern, filling in with a wordlist.

## Installation

```shell
go get github.com/caius/domainscan/cli
```

## Usage

```shell
$ domainscan --help

  --help, -h                You\'re looking at it
  --words-file, -w FILE     File containing replacements to iterate. One per line.
  --pattern, -p PATTERN     Pattern to insert words into. % gets replaced.
  --resolver, -r RESOLVER   Resolver to make DNS requests to
  --concurrency, -c         Concurrent requests to make. Defaults to 1
```

```shell
$ echo "theory" > names.txt
$ domainscan --words-file names.txt --pattern caius%.com --resolver 8.8.8.8
caiustheory.com 185.199.110.153
caiustheory.com 185.199.109.153
caiustheory.com 185.199.108.153
caiustheory.com 185.199.111.153
```
