# relaykit

relaykit is a Go toolkit for encrypted chunked message relay over text/message transports.

It provides the generic building blocks used by [chunkbridge](https://github.com/tiroq/chunkbridge)
and can be embedded in any Go application that needs to tunnel arbitrary data over a text-only
messaging channel.

## Module path

```
github.com/tiroq/relaykit
```

## Packages

| Package | Description |
|---|---|
| `pkg/compress` | Thin gzip compress/decompress wrappers |
| `pkg/crypto` | XChaCha20-Poly1305 encryption and Argon2id key derivation |
| `pkg/protocol` | CB/1 frame encoding, chunking, reassembly, and wire-format encode/decode |
| `pkg/ratelimit` | Token-bucket and adaptive rate limiter |
| `pkg/relay` | Session: request/response correlation over any Transport |
| `pkg/transport` | Transport interface and in-process MemoryTransport (for testing) |

## What is NOT included

- No HTTP proxy or exit executor
- No application config (YAML loaders, environment variables)
- No MAX-specific transport (`MaxTransport` lives in chunkbridge)
- No policy / domain allow-list enforcement
- No observability / structured logging

## Development

```bash
task test        # go test ./... -timeout 120s
task test-race   # go test -race ./... -timeout 120s
task lint        # go vet ./...
task fmt         # go fmt ./...
task check       # fmt-check + lint + test + test-race
```

## Wire format

The CB/1 wire format is defined in `pkg/protocol`:

```
CB1|D|<session_id>|<seq_num>|<base64(nonce||ciphertext)>
```

Encode pipeline: `Frame → json.Marshal → gzip → XChaCha20-Poly1305 → prepend 24-byte nonce → base64`

## License

MIT
