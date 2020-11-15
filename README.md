# goserve

`goserve` is a simple command-line http server. It can be used as an alternative to `python3`'s `http.server`. It leverages go's standard library to do all the heavy lifting.

## Installation

```
$ go get -u github.com/zhangyuannie/goserve
```
## Usage
```
goserve [options]
```

Options:
```
-d  the directory to serve (default to the current directory)
-p  the port number (default 8000)
```

To serve the current directory on port `8001`:

```
$ goserve -p 8001
```
