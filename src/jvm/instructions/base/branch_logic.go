package base

import "jvm/rtda"

//跳转逻辑指令
func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
