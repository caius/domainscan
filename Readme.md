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
```

We will output all valid domains on standard out as the domain name, and all missing domains on standard error with a message. If you pipe the output to a file you'll get valid domains in the file, and missing domains in the terminal.

```shell
$ echo -e "theory\n2" > names.txt

$ domainscan --words-file names.txt --pattern caius%.com --resolver 8.8.8.8 > valid.txt
caius2.com: not registered

$ cat valid.txt
caiustheory.com
```
