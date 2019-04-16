package heap

import (
	"github.com/therecipe/qt/nfc"
	"jvm/classfile"
	"jvm/rtda"
	"jvmgo/rtda/heap"
)

type Class struct {
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *classfile.ConstantPool
	fields            []*heap.Field
	methods           []*heap.Method
	loader            *heap.ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        *rtda.Slot
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	//TODO
	class.constantPool = newConstantPool()
	class.fields = newFields()
	class.methods = newMethods()
	return class
}
