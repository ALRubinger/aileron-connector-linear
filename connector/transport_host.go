// On non-wasip1 targets the host import ABI is unavailable; main()
// only exists on wasip1 (see main.go's build tag). This file ships
// the analogous non-wasip1 main(), so `go test` and friends still
// link without referencing the host-import functions.
//
// Running the resulting host-platform binary is not intended:
// connectors must execute inside Aileron's WASM sandbox to talk to
// any backing API.

//go:build !wasip1

package main

import (
	"errors"
	"os"
)

func main() {
	os.Exit(run(os.Stdin, os.Stdout, hostStubTransport))
}

// hostStubTransport always fails; the host build is for tests only,
// real GraphQL traffic flows through wasipTransport.
func hostStubTransport([]byte) ([]byte, int, error) {
	return nil, 0, errors.New("host platform: aileron_host transport unavailable; this connector must run inside the Aileron WASM sandbox")
}
