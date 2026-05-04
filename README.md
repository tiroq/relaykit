# relaykit

[![Go](https://img.shields.io/badge/go-1.25+-00ADD8?logo=go)](https://go.dev)

relaykit is a Go library of generic building blocks for encrypted, chunked message relay over text-only transports. It implements the CB/1 relay protocol: frame encoding, encryption, chunking, reassembly, session management, and rate limiting.

[chunkbridge](https://github.com/tiroq/chunkbridge) is one example consumer. relaykit itself contains no application-specific code.

## Module path

```
github.com/tiroq/relaykit
```

## Install

```bash
go get github.com/tiroq/relaykit@v0.1.0
```

## Packages

| Package | Description |
|---------|-------------|
| `pkg/compress` | Thin gzip compress/decompress wrappers |
| `pkg/crypto` | XChaCha20-Poly1305 encryption and Argon2id key derivation |
| `pkg/protocol` | CB/1 frame encoding, chunking, reassembly, and wire-format encode/decode |
| `pkg/ratelimit` | Token-bucket and adaptive rate limiter |
| `pkg/relay` | Session: request/response correlation over any Transport |
| `pkg/transport` | Transport interface and in-process MemoryTransport (for tests and local development) |

## Quick start

### Key derivation + frame encode/decode

```go
import (
    "fmt"
    "log"

    "github.com/tiroq/relaykit/pkg/crypto"
    "github.com/tiroq/relaykit/pkg/protocol"
)

// Derive a shared key from a passphrase.
salt, err := crypto.GenerateSalt()
if err != nil {
    log.Fatal(err)
}
key, err := crypto.DeriveKey([]byte("my-passphrase"), salt, crypto.DefaultDeriveParams)
if err != nil {
    log.Fatal(err)
}

// Encode a frame to a CB/1 wire string.
frame := &protocol.Frame{
    Version:     1,
    Type:        protocol.FrameDATA,
    SessionID:   "session-abc",
    RequestID:   "req-001",
    SeqNum:      1,
    TotalChunks: 1,
    ChunkIndex:  0,
    Payload:     []byte(`{"hello":"world"}`),
}
wireText, err := protocol.EncodeMessage(frame, key)
if err != nil {
    log.Fatal(err)
}

// Decode it back.
got, err := protocol.DecodeMessage(wireText, key)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(got.Payload)) // {"hello":"world"}
```

### Relay session (request/response correlation)

```go
import (
    "context"
    "time"

    "github.com/tiroq/relaykit/pkg/protocol"
    "github.com/tiroq/relaykit/pkg/relay"
    "github.com/tiroq/relaykit/pkg/transport"
)

ctx := context.Background()

// In-memory transport pair — useful for tests and local development.
clientT, _ := transport.NewMemoryPair(transport.MemoryOptions{})

sess := relay.NewSession("client-1", clientT, key)
go sess.Start(ctx) //nolint:errcheck

resp, err := sess.SendRequest(ctx, &protocol.Frame{
    Type:    protocol.FrameDATA,
    Payload: []byte(`{"hello":"world"}`),
}, 5*time.Second)
```

## Wire format

The CB/1 wire format is implemented in `pkg/protocol`:

```
CB1|D|<session_id>|<seq_num>|<base64(nonce||ciphertext)>
```

Encode pipeline: `Frame → json.Marshal → gzip → XChaCha20-Poly1305(AAD=sessionID|seqNum) → prepend 24-byte nonce → base64`

`DecodeMessage` is the exact reverse.

## What relaykit does NOT include

- No MAX transport or any platform-specific adapter
- No HTTP proxy or exit executor
- No YAML application config
- No domain allowlist or policy enforcement
- No observability / structured logging
- No browser cache logic

## Development

```bash
task fmt         # go fmt ./...  (mutating)
task fmt-check   # fail if any file is not gofmt-clean
task lint        # go vet ./...
task test        # go test ./... -timeout 120s
task test-race   # go test -race ./... -timeout 120s
task check       # fmt-check + lint + test + test-race + guard-no-app-imports
```

## License

MIT
