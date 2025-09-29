package sdkerrors

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestWithMessage_StatusErr(t *testing.T) {
	_, ok := status.FromError(WithMessage(errors.New("err"), "msg"))
	assert.False(t, ok)

	expectedStatus := status.New(codes.InvalidArgument, "invalid argument")
	statusWithMessage := WithMessage(expectedStatus.Err(), "msg")

	st, ok := status.FromError(statusWithMessage)
	assert.True(t, ok)
	assert.Equal(t, expectedStatus, st)

	st, ok = status.FromError(WithMessage(statusWithMessage, "extra msg"))
	assert.True(t, ok)
	assert.Equal(t, expectedStatus, st)
}

func TestWithMessage_ErrorsIs(t *testing.T) {
	assert.True(t, errors.Is(WithMessage(context.DeadlineExceeded, "wait operation"), context.DeadlineExceeded))
}
