package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("Clear", func(t *testing.T) {
		c := NewCache(10)

		for i := 0; i < 10; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}

		c.Clear()
		_, ok := c.Get("1")
		require.False(t, ok)
	})

	t.Run("pushing up first key", func(t *testing.T) {
		c := NewCache(5)

		c.Set("aaa", 1) // aaa: 1
		c.Set("bbb", 2) // aaa: 1, bbb: 2
		c.Set("ccc", 3) // aaa: 1, bbb: 2, ccc: 3
		c.Set("ddd", 4) // aaa: 1, bbb: 2, ccc: 3, ddd: 4
		c.Set("eee", 5) // aaa: 1, bbb: 2, ccc: 3, ddd: 4, eee: 5
		c.Set("fff", 6) // bbb: 2, ccc: 3, ddd: 4, eee: 5, fff: 6

		val, ok := c.Get("aaa")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("pushing up long used elements", func(t *testing.T) {
		c := NewCache(3)

		c.Set("aaa", 1) // aaa: 1
		c.Set("bbb", 2) // aaa: 1, bbb: 2
		c.Set("ccc", 3) // aaa: 1, bbb: 2, ccc: 3

		c.Get("aaa")
		c.Get("bbb")
		c.Get("aaa")
		c.Get("bbb")

		c.Set("ddd", 4) // aaa: 1, bbb: 2, ccc: 3, ddd: 4

		val, ok := c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
