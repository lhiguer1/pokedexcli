package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Minute * 5)
	if cache.cache == nil {
		t.Error("cache is nil")
	}

}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Minute * 5)
	cases := []struct {
		inputKey string
		inputval []byte
	}{
		{
			inputKey: "k1",
			inputval: []byte("v1"),
		},
		{
			inputKey: "k2",
			inputval: []byte("v2"),
		},
		{
			inputKey: "k3",
			inputval: []byte("v3"),
		},
	}
	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputval)

		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("%s not found", cas.inputKey)
			continue
		}

		if string(actual) != string(cas.inputval) {
			t.Errorf("%s does not match %s", string(actual), string(cas.inputKey))
			continue
		}

	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	k1 := "k1"
	v1 := "v1"
	cache.Add(k1, []byte(v1))

	time.Sleep(interval + time.Millisecond)
	if _, ok := cache.Get(k1); ok {
		t.Errorf("%s should have been reaped", k1)
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	k1 := "k1"
	v1 := "v1"
	cache.Add(k1, []byte(v1))

	time.Sleep(interval / 2)
	if _, ok := cache.Get(k1); !ok {
		t.Errorf("%s should not have been reaped", k1)
	}
}
