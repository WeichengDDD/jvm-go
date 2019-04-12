package rtda

import "math"

//操作数栈
type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}

	return nil
}

//int进出栈操作
func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++
}

func (self *OperandStack) PopInt() int32 {
	if self.size > 0 {
		self.size--
		return self.slots[self.size].num
	}

	return 0
}

//float进出栈操作
func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}

func (self *OperandStack) PopFloat() float32 {
	if self.size > 0 {
		self.size--
		bits := uint32(self.slots[self.size].num)
		return math.Float32frombits(bits)
	}

	return 0
}

//long进出栈操作
func (self *OperandStack) PushLong(val int64) {
	//低端32bit
	self.slots[self.size].num = int32(val)
	//高端32bit
	self.slots[self.size+1].num = int32(val >> 32)
	self.size++
}

func (self *OperandStack) PopLong() int64 {
	if self.size > 1 {
		self.size--
		high := uint32(self.slots[self.size].num)
		self.size--
		low := uint32(self.slots[self.size].num)
		return int64(high)<<32 | int64(low)
	}

	return 0
}

//double进出栈操作
func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}

func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

//ref进出栈操作
func (self *OperandStack) PushRef(ref *Object) {
	self.slots[self.size].ref = ref
	self.size++
}

func (self *OperandStack) PopRef() *Object {
	if self.size > 0 {
		self.size--
		ref := self.slots[self.size].ref
		self.slots[self.size].ref = nil
		return ref
	}

	return nil
}

//Slot进出栈操作
func (self *OperandStack) PushSlot(slot Slot) {
	self.slots[self.size] = slot
	self.size++
}

func (self *OperandStack) PopSlot() Slot {
	self.size--
	return self.slots[self.size]
}
