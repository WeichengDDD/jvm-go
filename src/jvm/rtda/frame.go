package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
}

func newFrame(maxLocals uint, maxStack uint) *Frame {
	return &Frame{
		localVars: newLocalVars(maxLocals),
		//	TODO
	}
}
