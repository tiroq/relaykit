// Package ratelimit provides rate-limiting primitives for controlling message
// throughput over a transport.
//
// TokenBucket is a classic token-bucket limiter: it allows bursts up to a
// configured maximum and refills at a steady rate.
//
// AdaptiveRateLimiter wraps three TokenBuckets (global, data, control) and
// reduces the data rate on receipt of a back-pressure signal (e.g. HTTP 429).
// Call On429 when the downstream transport reports rate-limiting; the limiter
// halves the data RPS and sets a backoff window automatically.
package ratelimit
