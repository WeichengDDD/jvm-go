package rtda

import "context"

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {

}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame) {
	//	TODO
}

func (selft *Thread) PopFrame() *Frame {
	//	TODO
}

func (self *Thread) CurrentFrame() *Frame {
	//	TODO
}
