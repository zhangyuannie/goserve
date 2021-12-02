# goserve

`goserve` is a simple and secure command-line HTTP server. It can be used as an alternative to `python3`'s `http.server`. It leverages go's standard library to do all the heavy lifting.

## Installation

```
$ go install github.com/zhangyuannie/goserve@latest
```

## Usage

```
goserve [options]
```

By default, the current working directory is served.

### Options

```
-cert string
      path to the TLS certificate file
-dir string
      alternate directory to serve
-host string
      address to listen on
-key string
      path to the TLS private key file
-password string
      password for basic authentication
-port int
      port number (default 8000)
-username string
      username for basic authentication
-version
      print goserve version
```

### Examples

To serve the current directory on port `8001`:

```
$ goserve --port 8001
```

To serve `/var/www/` over TLS:

```
$ goserve --dir /var/www/ --key key.pem --cert cert.pem
```
