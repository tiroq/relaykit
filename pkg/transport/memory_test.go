package transport_test

import (
	"context"
	"testing"

	"github.com/tiroq/relaykit/pkg/transport"
)

func TestNewMemoryPairSendReceive(t *testing.T) {
	a, b := transport.NewMemoryPair(transport.MemoryOptions{})
	defer a.Close()
	defer b.Close()

	ctx := context.Background()

	want := "hello relaykit"
	if err := a.Send(ctx, transport.Message{Text: want}); err != nil {
		t.Fatalf("Send: %v", err)
	}

	ch, err := b.Receive(ctx)
	if err != nil {
		t.Fatalf("Receive: %v", err)
	}

	msg := <-ch
	if msg.Text != want {
		t.Errorf("got %q want %q", msg.Text, want)
	}
}

func TestMemoryTransportClose(t *testing.T) {
	a, b := transport.NewMemoryPair(transport.MemoryOptions{})

	ctx := context.Background()
	ch, err := b.Receive(ctx)
	if err != nil {
		t.Fatalf("Receive: %v", err)
	}

	// Closing a should also close b (paired close).
	if err := a.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}

	// The receive channel on b should be closed once its closed signal fires.
	// Drain until closed.
	for range ch {
	}
}
