# urlbyter

`urlbyter` is a CLI tool that takes a file with a list of URLs and returns a list of the sizes of the responses of each URL.

## Installation

```bash
$ go get -u github.com/thmhoag/urlbyter/cmd/urlbyter
```

## Getting started
For a list of commands, you can start with `urlbyter --help`. 

### Getting a list of URLs 

Create a text file with a list of URLs (including HTTP/HTTPS).

To use it, pass the path to the text file to `urlbyter`:
```bash
$ urlbyter ./path/to/myfile.txt
```