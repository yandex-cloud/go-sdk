package operation

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func marshalAny(t *testing.T, msg proto.Message) *anypb.Any {
	any, err := anypb.New(msg)
	require.NoError(t, err)
	return any
}

//go:generate mockery -name=mockClient -out-name=MockClient -testonly

// TODO(skipor): get rid of, when mockery will understand type aliases
//
//nolint:deadcode,megacheck
type mockClient interface {
	Client
}

type mockTimer struct {
	id   int
	mock mock.Mock
}

func (t *mockTimer) Stop(id int) {
	t.mock.Called(id)
}

func (t *mockTimer) Start(id int, d time.Duration) {
	t.mock.Called(id, d)
}

func (t *mockTimer) Read(id int) {
	t.mock.Called(id)
}

func (t *mockTimer) New(d time.Duration) (func() <-chan time.Time, func() bool) {
	t.id++
	id := t.id
	t.Start(id, d)
	return func() <-chan time.Time {
			t.Read(id)
			ch := make(chan time.Time)
			close(ch)
			return ch
		}, func() bool {
			t.Stop(id)
			return true
		}
}
