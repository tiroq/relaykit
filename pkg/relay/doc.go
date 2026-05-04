// Package relay provides request/response session management over a message
// transport.
//
// A Session correlates outbound requests with inbound responses using
// request IDs embedded in CB/1 frames.  Callers use SendRequest to send an
// encoded frame and wait for the matching response; the remote end calls
// SendResponse to return data.
//
// Session is transport-agnostic: it works with any value that satisfies
// [github.com/tiroq/relaykit/pkg/transport.Transport].
package relay
