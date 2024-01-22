package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGet(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inpKey string
		inpVal []byte
	}{
		{
			inpKey: "key1",
			inpVal: []byte("val1"),
		},
		{
			inpKey: "key2",
			inpVal: []byte("val2"),
		},
		{
			inpKey: "",
			inpVal: []byte("val3"),
		},
	}
	for _, cas := range cases {

		cache.Add(cas.inpKey, cas.inpVal)
		actual, ok := cache.Get(cas.inpKey)
		if !ok {
			t.Errorf("expected to find %s", cas.inpKey)
			continue
		}
		if string(actual) != string(cas.inpVal) {
			t.Errorf("%s does not match %s", string(actual), string(cas.inpVal))
			continue
		}
	}

}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))
	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("expected %s to be reaped", keyOne)
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))
	time.Sleep(interval / 2)

	_, ok := cache.Get(keyOne)
	if !ok {
		t.Errorf("expected %s to not be reaped", keyOne)
	}
}
