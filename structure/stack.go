//Package structure
//@program: gutils
//@author: LiuShuo
//@create: 2021-10-12 11:47
package structure

import (
	"errors"
	"github.com/gogf/gf/frame/g"
	"sync"
)

type Stack struct {
	maxNum int           //栈最大元素
	top    int           //栈顶当前下标
	arr    []interface{} //模拟栈
	lock   sync.Mutex    //锁
}

// NewStack 创建并初始化一个新的栈并返回指针
//  - length 栈大小
func NewStack(length int) *Stack {
	arr := make([]interface{}, length)
	return &Stack{maxNum: length, top: -1, arr: arr, lock: sync.Mutex{}}
}

// Push 入栈
func (s *Stack) Push(val interface{}) (err error) {
	if s.IsFull() {
		err = errors.New("Stack Full. ")
		return
	}
	s.top++            //先向上走一步
	s.arr[s.top] = val //再赋值
	return
}

// Pop 出栈
func (s *Stack) Pop() (val interface{}, err error) {
	s.lock.Lock()
	if s.IsEmpty() {
		err = errors.New("Stack Empty. ")
		return
	}
	val = s.arr[s.top]
	s.top--
	s.lock.Unlock()
	return
}

// List 遍历栈
func (s *Stack) List() (err error) {
	if s.IsEmpty() {
		err = errors.New("Can't List Stack Because Empty. ")
		return
	}
	g.Log().Info(">>>>> List Stack <<<<<")
	for i := s.top; i >= 0; i-- {
		g.Log().Infof("Stack[%d] = %s", i, s.arr[i])
	}
	return
}

func (s *Stack) IsFull() bool {
	return s.top+1 >= s.maxNum
}
func (s *Stack) IsEmpty() bool {
	return s.top == -1
}
