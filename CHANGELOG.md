# Changelog

All notable changes to relaykit are documented here.

Format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).
relaykit uses [Semantic Versioning](https://semver.org/).

---

## [Unreleased]

## [v0.1.0] — unreleased

Initial extraction from [chunkbridge](https://github.com/tiroq/chunkbridge).

### Added

- `pkg/compress`: thin gzip compress/decompress wrappers.
- `pkg/crypto`: XChaCha20-Poly1305 `Encrypt`/`Decrypt`; Argon2id `DeriveKey`/`GenerateSalt`/`DefaultDeriveParams`.
- `pkg/protocol`: CB/1 wire format — `Frame`, `EncodeMessage`, `DecodeMessage`, `Chunk`, `Reassembler`, `NewACKFrame`, `IsACK`, error code constants, `MarshalErrorPayload`/`UnmarshalErrorPayload`.
- `pkg/ratelimit`: `TokenBucket` rate limiter; `AdaptiveRateLimiter` with `On429` backoff.
- `pkg/relay`: `Session` with `SendRequest`, `SendResponse`, `WithRateLimiter`, `WithMaxPendingRequests`.
- `pkg/transport`: `Transport` interface; `MemoryTransport` / `NewMemoryPair` for in-process testing.
- Package-level `doc.go` files for all six packages.
- `Taskfile.yml` with `fmt`, `fmt-check`, `lint`, `test`, `test-race`, `check`, `guard-no-app-imports` tasks.
- `docs/release.md` with step-by-step publish checklist.
