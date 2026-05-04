# Releasing relaykit

Follow this checklist every time a new version is tagged.

## Pre-release checks

Run all checks from the relaykit root:

```bash
# 1. All tests pass (including race detector)
task check

# 2. No unformatted files
gofmt -l .
# Expected: no output

# 3. No chunkbridge references in Go source
grep -R --include="*.go" "github.com/tiroq/chunkbridge" pkg/
# Expected: no output (exit 1 means references found)

# 4. No MAX/platform-specific identifiers in Go source
grep -R --include="*.go" \
  -e "MaxTransport" \
  -e "platform-api" \
  -e "peer_chat_id" \
  -e "/messages/poll" \
  pkg/
# Expected: no output
```

Alternatively, rely on the `guard-no-app-imports` Taskfile task (included in `task check`).

## Tag and push

```bash
cd /path/to/relaykit

git tag v0.1.0
git push origin main --tags
```

After the tag is pushed, the module is immediately available via the Go module proxy.

## Switch chunkbridge from local replace to tagged release

```bash
cd /path/to/chunkbridge

# 1. Remove the replace directive from go.mod:
#    replace github.com/tiroq/relaykit => ../relaykit
#    (edit manually or with sed)

# 2. Fetch the tagged release and update go.sum
go get github.com/tiroq/relaykit@v0.1.0
go mod tidy

# 3. Verify
task check
CHUNKBRIDGE_SHARED_KEY=testpassphrase go run ./cmd/chunkbridge selftest
```

## Verify the module is accessible

```bash
go list -m github.com/tiroq/relaykit@v0.1.0
# Expected: github.com/tiroq/relaykit v0.1.0
```

## After release

- Update `CHANGELOG.md`: move the `[Unreleased]` section content under the new version heading.
- Open a new `[Unreleased]` section at the top for the next release.
