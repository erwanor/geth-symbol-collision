Providing an example of quirky linker behavior.

The `celo-blockchain` (fork) and `go-ethereum` (orig) both wrap the
`secp256k1` library this creates an issue at link time: the same
symbols are defined multiple time.

No big deal, you can simply inform the linker that you're in control
by passing `--allow-multiple-definitions` like a grown up: `go build
--ldflags '-extldflags "-Wl,--allow-multiple-definition"'`

For some reason however, it seems that the Darwin linker has
deprecated its very own `-zmuldefs` flag.

Now, this would be solved by _not_ having colliding namespaces in
the first place. That is to say, CGO should take care of this. One
can always dream of course.
