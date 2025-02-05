package netmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func newFilter(name string, k, v string, op Operation, fs ...Filter) (f Filter) {
	f.SetName(name)
	f.SetKey(k)
	f.SetOperation(op)
	f.SetValue(v)
	f.SetInnerFilters(fs...)
	return f
}

func newSelector(name string, attr string, c Clause, count uint32, filter string) (s Selector) {
	s.SetName(name)
	s.SetAttribute(attr)
	s.SetCount(count)
	s.SetClause(c)
	s.SetFilter(filter)
	return s
}

func newPlacementPolicy(bf uint32, rs []Replica, ss []Selector, fs []Filter) *PlacementPolicy {
	p := NewPlacementPolicy()
	p.SetContainerBackupFactor(bf)
	p.SetReplicas(rs...)
	p.SetSelectors(ss...)
	p.SetFilters(fs...)
	return p
}

func newReplica(c uint32, s string) (r Replica) {
	r.SetCount(c)
	r.SetSelector(s)
	return r
}

func nodeInfoFromAttributes(props ...string) NodeInfo {
	attrs := make([]NodeAttribute, len(props)/2)
	for i := range attrs {
		attrs[i].SetKey(props[i*2])
		attrs[i].SetValue(props[i*2+1])
	}
	n := NewNodeInfo()
	n.SetAttributes(attrs...)
	return *n
}

type enumIface interface {
	FromString(string) bool
	String() string
}

type enumStringItem struct {
	val enumIface
	str string
}

func testEnumStrings(t *testing.T, e enumIface, items []enumStringItem) {
	for _, item := range items {
		require.Equal(t, item.str, item.val.String())

		s := item.val.String()

		require.True(t, e.FromString(s), s)

		require.EqualValues(t, item.val, e, item.val)
	}

	// incorrect strings
	for _, str := range []string{
		"some string",
		"undefined",
	} {
		require.False(t, e.FromString(str))
	}
}
