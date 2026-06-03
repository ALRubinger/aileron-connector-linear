// Package main is the WASM source for the aileron-connector-linear
// reference connector. It targets Go's native WASI Preview 1
// (`GOOS=wasip1 GOARCH=wasm`) and reads JSON-shaped op dispatches from
// stdin, writing JSON results to stdout.
//
// This file is wasip1-tagged because it imports `aileron_host`, the
// runtime's host-import ABI for outbound HTTP. The dispatch core and
// per-op handlers live in dispatch.go / handlers.go / graphql.go,
// which build on every Go target so `go test` exercises them on the
// host platform with a mock Transport. The non-wasip1 main() lives
// in transport_host.go alongside the always-fail host-side transport
// stub.
//
// Build:
//
//	cd connector && GOOS=wasip1 GOARCH=wasm go build -trimpath \
//	  -ldflags="-s -w" -o ../connector.wasm .
//
// Or via Taskfile from the repo root:
//
//	task build

//go:build wasip1

package main

import (
	"fmt"
	"os"
	"unsafe"
)

//go:wasmimport aileron_host log
//go:noescape
func hostLog(levelPtr unsafe.Pointer, levelLen uint32, msgPtr unsafe.Pointer, msgLen uint32)

//go:wasmimport aileron_host http_request
//go:noescape
func hostHTTPRequest(reqPtr unsafe.Pointer, reqLen uint32) int32

//go:wasmimport aileron_host http_response_size
//go:noescape
func hostHTTPResponseSize() int32

//go:wasmimport aileron_host http_response_status
//go:noescape
func hostHTTPResponseStatus() int32

//go:wasmimport aileron_host http_response_read
//go:noescape
func hostHTTPResponseRead(dstPtr unsafe.Pointer, dstLen uint32) int32

// _emptyPtrSentinel keeps the address of an empty byte slice valid; Go
// can't take the address of an empty slice's first element directly.
var _emptyPtrSentinel = [1]byte{}

func ptr(b []byte) unsafe.Pointer {
	if len(b) == 0 {
		return unsafe.Pointer(&_emptyPtrSentinel[0])
	}
	return unsafe.Pointer(&b[0])
}

func main() {
	os.Exit(run(os.Stdin, os.Stdout, wasipTransport))
}

// wasipTransport sends the request envelope through Aileron's host
// import ABI and reads the response back. Implements Transport.
func wasipTransport(req []byte) ([]byte, int, error) {
	rc := hostHTTPRequest(ptr(req), uint32(len(req)))
	if rc != 0 {
		// The host has stuck a structured *Error on per-call state and
		// will surface it to the caller; bailing here avoids
		// double-wrapping. rc != 0 covers capability_denied and
		// transport-level failure both.
		return nil, 0, fmt.Errorf("http_request denied or failed (rc=%d)", rc)
	}
	size := hostHTTPResponseSize()
	if size < 0 {
		return nil, 0, fmt.Errorf("http_response_size returned %d", size)
	}
	respBody := make([]byte, size)
	if size > 0 {
		n := hostHTTPResponseRead(ptr(respBody), uint32(size))
		if n < 0 {
			return nil, 0, fmt.Errorf("http_response_read returned %d", n)
		}
		respBody = respBody[:n]
	}
	return respBody, int(hostHTTPResponseStatus()), nil
}

// hostLog is currently unused but kept exported-to-host so the
// runtime can call into us if we ever need to surface progress
// information mid-op. It is referenced by `_ = hostLog` to satisfy
// staticcheck without actually invoking the host call.
var _ = hostLog
