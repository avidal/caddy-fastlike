# caddy-fastlike

A [Caddy v2](https://caddyserver.com) plugin that can run WebAssembly code via [fastlike](https://fastlike.dev)

## try it out

Extremely beta. But you can try it out:

```
# Build the example wasm
$ cd testdata && cargo build

# Get xcaddy and use it to build this plugin
$ go get -u github.com/caddyserver/xcaddy/cmd/xcaddy
$ env CGO_ENABLED=1 xcaddy build --with fastlike.dev/caddy-fastlike=$(pwd)

# Run caddy with the example config
$ ./caddy run --config example-config.json
```

Now all requests to `localhost:2105` will run through the wasm module via fastlike. Try
http://localhost:2105/teapot for a static response.
