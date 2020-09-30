package stack

import (
	"strconv" //此包提供了基本数据类型转化为字符串，或者由字符串转化为基本数据类型
)

type Stack struct {
	i    int
	data [20]int
}

func (s *Stack) Push(k int) {
	s.data[s.i] = k
	s.i++
}

func (s *Stack) Pop(ret int) {
	s.i--
	ret = s.data[s.i]
}

func (s *Stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
