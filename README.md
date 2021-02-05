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
-d string
      directory to serve (default current directory)
-p int
      port number (default 8000)
-password string
      password for basic authentication (default none)
-username string
      username for basic authentication (default none)
-version
      print goserve version
```

To serve the current directory on port `8001`:

```
$ goserve -p 8001
```
