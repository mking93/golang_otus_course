package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})

	t.Run("Remove", func(t *testing.T) {
		l := NewList()

		for i := 1; i < 10; i++ {
			l.PushBack(i)
		} // [1, 2, 3, 4, 5, 6, 7, 8, 9]

		item1 := l.Back()  // 9
		item2 := l.Front() // 1
		l.Remove(item1)    // [1, 2, 3, 4, 5, 6, 7, 8]
		l.Remove(item2)    // [2, 3, 4, 5, 6, 7, 8]

		result := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			result = append(result, i.Value.(int))
		}

		require.Equal(t, []int{2, 3, 4, 5, 6, 7, 8}, result)
	})

	t.Run("MoveToFront", func(t *testing.T) {
		l := NewList()

		for i := 1; i < 10; i++ {
			l.PushBack(i)
		} // [1, 2, 3, 4, 5, 6, 7, 8, 9]

		item := l.Back().Prev // 8
		l.MoveToFront(item)   // [8, 1, 2, 3, 4, 5, 6, 7, 9]

		result := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			result = append(result, i.Value.(int))
		}

		require.Equal(t, []int{8, 1, 2, 3, 4, 5, 6, 7, 9}, result)
	})
}
