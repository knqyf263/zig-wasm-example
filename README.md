## Zig allocation example

This example shows how to pass strings in and out of a Wasm function defined
in Zig, built with `zig build`.

Ex.
```bash
$ go run greet.go wazero
wasm >> Hello, wazero!
go >> Hello, wazero!
```

Under the covers, [main.zig](testdata/src/main.zig) does a few things of interest:
* Uses `@ptrToInt` to change a Zig pointer to a numeric type.
* Uses `@intToPtr` to build back a string from a pointer, len pair.
* Relies on Zig not having garbage collector.

Zig code exports "malloc" and "free", which we use for that purpose.