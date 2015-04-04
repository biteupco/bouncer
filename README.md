# bouncer
[![Build Status](https://travis-ci.org/gobbl/bouncer.svg?branch=master)](https://travis-ci.org/gobbl/bouncer)

Toughest guy in the hood.

## Setup

Please ensure you have Go installed.

You can check if you have Go installed properly by typing:
```
$ go help
```

Once done, pull down Bouncer via the command:
```
$ go get github.com/gobbl/bouncer
```

To install all dependencies of bouncer, simply:
```
$ go get ./...
```

Build the code, and run the server by:
```
$ go build
$ go run app.go
```

### Editing code

We need to avoid importing unnecessary libraries, to which Golang would complain :)
To help us with that, you can install [`goimports`](https://github.com/bradfitz/goimports) which cleans up imports as well as formatting code.

- [goimports + GoSublime with Sublime Text](http://michaelwhatcott.com/gosublime-goimports/)
