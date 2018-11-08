// Copyright (c) 2018 Yandex LLC. All rights reserved.
// Author: Dmitry Novikov <novikoff@yandex-team.ru>

package sdkresolvers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	compute "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
)

func TestBaseResolverFindName(t *testing.T) {
	base := func() *BaseResolver {
		x := &BaseResolver{Name: "name1"}
		x.opts = &resolveOptions{out: &x.id}
		return x
	}
	t.Run("only one correct name", func(t *testing.T) {
		x := base()
		err := x.findName("test", []*compute.Instance{
			{Id: "id1", Name: "name1"},
		}, nil)
		require.NoError(t, err)
		assert.Equal(t, x.id, "id1")
	})
	t.Run("two records with same name", func(t *testing.T) {
		x := base()
		err := x.findName("test", []*compute.Instance{
			{Id: "id1", Name: "name1"},
			{Id: "id2", Name: "name1"},
		}, nil)
		require.Error(t, err)
		assert.Equal(t, "multiple test items with name \"name1\" found", err.Error())
	})
	t.Run("two records with same name but not found", func(t *testing.T) {
		x := base()
		err := x.findName("test", []*compute.Instance{
			{Id: "id1", Name: "name2"},
			{Id: "id2", Name: "name2"},
		}, nil)
		require.Error(t, err)
		assert.Equal(t, &ErrNotFound{Caption: "test", Name: "name1"}, err)
	})
	t.Run("resolve error", func(t *testing.T) {
		x := base()
		err := x.findName("test", nil, errors.New("forward this"))
		require.Error(t, err)
		assert.Equal(t, "failed to find test with name \"name1\": forward this", err.Error())
	})
	t.Run("multiple items returned 1", func(t *testing.T) {
		x := base()
		err := x.findName("test", []*compute.Instance{
			{Id: "id1", Name: "name1"},
			{Id: "id2", Name: "name2"},
		}, nil)
		require.NoError(t, err)
		assert.Equal(t, x.id, "id1")
	})
	t.Run("multiple items returned 2", func(t *testing.T) {
		x := base()
		err := x.findName("test", []*compute.Instance{
			{Id: "id2", Name: "name2"},
			{Id: "id1", Name: "name1"},
		}, nil)
		require.NoError(t, err)
		assert.Equal(t, x.id, "id1")
	})
}
