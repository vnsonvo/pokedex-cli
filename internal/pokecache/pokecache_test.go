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

func TestAddCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("value1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("value2"),
		},
		{inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	for _, testCase := range cases {
		cache.Add(testCase.inputKey, testCase.inputVal)

		actual, ok := cache.Get(testCase.inputKey)

		if !ok {
			t.Errorf("%s not found", testCase.inputKey)
			continue
		}

		if string(actual) != string(testCase.inputVal) {
			t.Errorf("%s doesn't match %s", string(testCase.inputKey), string(actual))
			continue
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("value1"))

	time.Sleep(interval - time.Millisecond)

	_, ok := cache.Get(keyOne)
	if !ok {
		t.Errorf("%s should have been reaped", keyOne)
	}
}
