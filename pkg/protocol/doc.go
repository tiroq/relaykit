// Package protocol implements the CB/1 relay wire format: frame encoding,
// encryption-aware message serialisation, chunking, reassembly, ACK helpers,
// and structured error payloads.
//
// Wire format:
//
//	CB1|D|<session_id>|<seq_num>|<base64(nonce||ciphertext)>
//
// Encode pipeline: Frame → json.Marshal → gzip → XChaCha20-Poly1305 → prepend
// 24-byte nonce → base64-encode.  DecodeMessage is the exact reverse.
//
// Use Chunk to split a large encoded frame across multiple transport messages,
// and Reassembler to collect the pieces back into a single frame.
package protocol
