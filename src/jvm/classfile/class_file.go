package classfile

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributs    []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	return nil, nil
}

func (self *ClassFile) read(reader *ClassReader) {

}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) ClassName() string {
	return ""
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		//TODO
		return ""
	}

	//java.lang.Object没有超类
	return ""
}

func (self *ClassFile) InterfaceNames() []string {
	return nil
}

//魔数检查
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

//版本号检查
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	//读取版本号
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()

	//版本号合法性判断，45以后次版本号均为零。
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}

	panic("java.lang.UnsupportedClassVersionError!")
}
