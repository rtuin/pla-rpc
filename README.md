# Pla-RPC

Run Pla commands from anywhere!

## Installation
### Manual

```bash
$ go get ./...
```

## Usage

```bash
$ go run cmd/pla-rpc.go
```

## Configuration

By default Pla-RPC binds to `localhost:7777`. You can change this value by passing the `-bind=0.0.0.0:8888` parameter when starting Pla-RPC, like so:

```bash
$ go run cmd/pla-rpc.go -bind=127.0.0.1:9090
```

## Dependencies

* [Pla library](https://github.com/rtuin/go-plalib.git)

## Change log

Please see [CHANGELOG](CHANGELOG.md) for more information what has changed recently.

## Credits

- [Richard Tuin](http://github.com/rtuin)
- [All Contributors](https://github.com/rtuin/pla-rpc/contributors)