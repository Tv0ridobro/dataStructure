package list

import (
	"math/rand"
	"testing"

	"github.com/Tv0ridobro/data-structure/slices"
)

func TestList_PopBack(t *testing.T) {
	t.Parallel()
	list := New[int]()
	for i := 0; i < 20; i++ {
		list.PushBack(i)
	}
	for i := 19; i >= 0; i-- {
		if v := list.PopBack(); v != i {
			t.Errorf("Wrong value got %d expected %d", v, i)
		}
	}
}

func TestList_PopBackEmpty(t *testing.T) {
	t.Parallel()
	list := New[int]()
	var a int
	for i := 0; i < 20; i++ {
		v := list.PopBack()
		if v != a {
			t.Errorf("Wrong value got %d expected %d", v, a)
		}
	}
}

func TestList_PopFrontEmpty(t *testing.T) {
	t.Parallel()
	list := New[int]()
	var a int
	for i := 0; i < 20; i++ {
		v := list.PopFront()
		if v != a {
			t.Errorf("Wrong value got %d expected %d", v, a)
		}
	}
}

func TestList_NodeEmpty(t *testing.T) {
	t.Parallel()
	list := New[int]()
	for i := 0; i < 20; i++ {
		v := list.Node(i)
		if v != nil {
			t.Errorf("Wrong value got %v expected %v", v, nil)
		}
	}
}

func TestList_Back(t *testing.T) {
	t.Parallel()
	list := New[int]()
	var a int
	if v := list.Back(); v != a {
		t.Errorf("Wrong value got %d expected %d", v, a)
	}
	list.PushBack(2)
	if v := list.Back(); v != 2 {
		t.Errorf("Wrong value got %d expected %d", v, 2)
	}
	list.PushBack(3)
	if v := list.Back(); v != 3 {
		t.Errorf("Wrong value got %d expected %d", v, 3)
	}
}

func TestList_Clear(t *testing.T) {
	t.Parallel()
	list := New[int]()
	list.PushBack(10)
	list.PushBack(2)
	list.PushBack(27)
	list.PushFront(5)
	list.PushBack(8)
	list.PushBack(9)
	list.Clear()
	var a int
	size := list.Len()
	if size != 0 {
		t.Errorf("Wrong value got %d expected %d", size, 0)
	}
	for i := 0; i < 100; i++ {
		v := list.PopBack()
		if v != a {
			t.Errorf("Wrong value got %d expected %d", v, a)
		}
	}
}

func TestList_Reverse(t *testing.T) {
	t.Parallel()
	list := New[int]()
	list.Reverse()
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i := 0; i < len(values); i++ {
		list.PushBack(values[len(values)-i-1])
	}
	list.Reverse()
	if v := list.All(); !slices.Equal(v, values) {
		t.Errorf("Wrong value got %v expected %v", v, values)
	}
}

func TestList_PopFront(t *testing.T) {
	t.Parallel()
	list := New[int]()
	for i := 0; i < 20; i++ {
		list.PushFront(i)
	}
	for i := 19; i >= 0; i-- {
		if v := list.PopFront(); v != i {
			t.Errorf("Wrong value got %d expected %d", v, i)
		}
	}
}

func TestList_FrontAndBack(t *testing.T) {
	t.Parallel()
	list := New[int]()
	value := 10
	list.PushBack(value)
	v1 := list.Back()
	v2 := list.Front()
	if v1 != value {
		t.Errorf("Wrong value got %d expected %d", v1, value)
	}
	if v2 != value {
		t.Errorf("Wrong value got %d expected %d", v2, value)
	}
}

func TestList_FrontAndBackEmpty(t *testing.T) {
	t.Parallel()
	list := New[int]()
	var value int
	v1 := list.Back()
	v2 := list.Front()
	if v1 != value {
		t.Errorf("Wrong value got %d expected %d", v1, value)
	}
	if v2 != value {
		t.Errorf("Wrong value got %d expected %d", v2, value)
	}
}

func TestList_Len(t *testing.T) {
	t.Parallel()
	list := New[int]()
	if list.Len() != 0 {
		t.Error("wrong length size")
	}
	list.PushFront(7)
	if list.Len() != 1 {
		t.Error("wrong length size")
	}
	list.PopBack()
	if list.Len() != 0 {
		t.Error("wrong length size")
	}
	for i := 0; i < 15; i++ {
		list.PushBack(0)
	}
	if list.Len() != 15 {
		t.Error("wrong length size")
	}
}

func TestList_GetAllAndPeek(t *testing.T) {
	t.Parallel()
	length := 100
	data := make([]int, 0)
	list := New[int]()
	for i := 0; i < length; i++ {
		value := rand.Int()
		command := rand.Intn(2)
		switch command {
		case 0:
			data = append(data, value)
			list.PushBack(value)
		case 1:
			data = append([]int{value}, data...)
			list.PushFront(value)
		}
	}
	for i := 0; i < length; i++ {
		if v := list.Peek(i); data[i] != v {
			t.Errorf("data[i] != list.Peek(i) %d %d", data[i], v)
		}
	}
	if !slices.Equal(data, list.All()) {
		t.Error(data, " != ", list.All())
	}
}

func TestList_ChangeAt(t *testing.T) {
	t.Parallel()
	list := New[int]()
	for i := 0; i < 20; i++ {
		list.PushBack(i)
	}
	for i := 0; i < 20; i++ {
		list.ChangeAt(i, i+30)
		if value := list.Peek(i); value != i+30 {
			t.Errorf("value != list.Peek(i) %d %d", i+30, value)
		}
	}
}

func TestList_PeekEmpty(t *testing.T) {
	t.Parallel()
	list := New[int]()
	v := list.Peek(10)
	var a int
	if v != a {
		t.Errorf("Wrong value got %d expected %d", v, a)
	}
}

func TestList_Cut(t *testing.T) {
	t.Parallel()
	list := New[int]()
	for i := 0; i < 20; i++ {
		list.PushBack(i)
	}
	l, r := list.Cut(5)
	if v := l.Len(); v != 6 {
		t.Errorf("l len %d != %d", v, 6)
	}
	if v := r.Len(); v != 14 {
		t.Errorf("l len %d != %d", v, 14)
	}
	if v := l.All(); !slices.Equal(v, []int{0, 1, 2, 3, 4, 5}) {
		t.Errorf("l All is wrong %v", v)
	}
	if v := r.All(); !slices.Equal(v, []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}) {
		t.Errorf("l All is wrong %v", v)
	}
}

func TestList_Cut2(t *testing.T) {
	t.Parallel()
	list := New[int]()
	list.PushBack(1)
	l, r := list.Cut(0)
	if v := l.All(); !slices.Equal(v, []int{1}) {
		t.Errorf("Wrong value got %v expected %v", v, []int{1})
	}
	if v := r.All(); !slices.Equal(v, []int{}) {
		t.Errorf("Wrong value got %v expected %v", v, []int{})
	}
}

func TestList_Cut3(t *testing.T) {
	t.Parallel()
	list := New[int]()
	list.PushBack(1)
	l, r := list.Cut(1)
	if v := l.All(); !slices.Equal(v, []int{1}) {
		t.Errorf("Wrong value got %v expected %v", v, []int{1})
	}
	if v := r.All(); !slices.Equal(v, []int{}) {
		t.Errorf("Wrong value got %v expected %v", v, []int{})
	}
}

func TestList_Merge(t *testing.T) {
	t.Parallel()
	tests := []struct {
		values1 []int
		values2 []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}},
		{[]int{}, []int{6, 7, 8, 9, 10}},
		{[]int{1, 2, 3, 4, 5}, []int{}},
		{[]int{1}, []int{6, 7, 8, 9, 10}},
		{[]int{1, 2, 3, 4, 5}, []int{6}},
	}
	for _, test := range tests {
		list1 := New[int]()
		list2 := New[int]()
		for i := 0; i < len(test.values1); i++ {
			list1.PushBack(i)
		}
		for i := 0; i < len(test.values2); i++ {
			list2.PushBack(i)
		}
		list1.Merge(list2)
		if v := list1.Len(); v != len(test.values1)+len(test.values2) {
			t.Errorf("Wrong value got %d expected %d", v, len(test.values1)+len(test.values2))
		}
		if v := list1.All(); slices.Equal(v, append(test.values1, test.values2...)) {
			t.Errorf("Wrong value got %v expected %v", v, append(test.values1, test.values2...))
		}
	}
}
