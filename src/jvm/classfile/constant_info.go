package classfile

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

//根据tag创建常量
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{}
	case CONSTANT_Class:
	case CONSTANT_Fieldref:
	case CONSTANT_Methodref:
	case CONSTANT_InterfaceMethodref:
	case CONSTANT_NameAndType:
	case CONSTANT_MethodType:
	case CONSTANT_MethodHandle:
	case CONSTANT_InvokeDynamic:
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")

	}
}

//读取常量信息
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	//获取tag值
	tag := reader.readUint8()
	//生成常量结构体
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}
