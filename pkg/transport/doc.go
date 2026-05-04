// Package transport defines a text-message transport interface and provides an
// in-memory implementation for use in tests and local development.
//
// The Transport interface has three methods: Send, Receive, and Close.
// MemoryTransport is a fully in-process, goroutine-safe implementation backed
// by buffered channels.  Use NewMemoryPair to obtain a pair of transports
// where each side's Send delivers to the other side's Receive channel.
package transport
