package classfile

import "fmt"

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
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)

	return
}

func (self *ClassFile) read(reader *ClassReader) {
	//检查魔数
	self.readAndCheckMagic(reader)
	//检查版本
	self.readAndCheckVersion(reader)
	//读取常量池
	self.constantPool = readConstantPool(reader)
	//类权限
	self.accessFlags = reader.readUint16()
	//u2索引值，指向本类描述符(CONSTANT_Class_info)
	self.thisClass = reader.readUint16()
	//u2索引值，指向父类描述符
	self.superClass = reader.readUint16()
	//索引数组，指向接口描述符
	self.interfaces = reader.readUnit16s()
	//字段表
	self.fields = readMembers(reader, self.constantPool)
	//方法表
	self.methods = readMembers(reader, self.constantPool)
	//属性表
	self.attributs = readAttributes(reader, self.constantPool)
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

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}

	//java.lang.Object没有超类
	return ""
}

func (self *ClassFile) InterfaceNames() []string {
	return nil
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}
