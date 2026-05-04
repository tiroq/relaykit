// Package crypto provides symmetric encryption and key derivation primitives
// used by the CB/1 relay protocol.
//
// Encryption uses XChaCha20-Poly1305 (via golang.org/x/crypto/chacha20poly1305)
// with a 24-byte random nonce prepended to the ciphertext.
//
// Key derivation uses Argon2id (via golang.org/x/crypto/argon2).  Call
// DeriveKey with a passphrase, a 16-byte salt, and DeriveParams to obtain a
// 32-byte key suitable for use with Encrypt/Decrypt.  GenerateSalt produces a
// cryptographically random 16-byte salt.
package crypto
