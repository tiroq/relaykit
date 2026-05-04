// Package compress provides thin wrappers around the standard library's
// compress/gzip package for compressing and decompressing byte slices.
//
// These helpers are used by the CB/1 encode/decode pipeline to shrink frame
// payloads before encryption.
package compress
