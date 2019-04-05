package classfile

//常量表
type ConstantPool []ConstantInfo

//解析常量表
func readConstantPool(reader *ClassReader) ConstantPool {
	//常量表容量
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)

	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		//long, double为8字节，占2个位置
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}

	return cp
}

//按常量表索引查找常量
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}

	panic("Invalid constant pool index!")
}

//从常量表查找类名
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

//解析字段或方法描述符
func (self ConstantPool) getNameAndType(index uint16) (name string, descriptor string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name = self.getUtf8(ntInfo.nameIndex)
	descriptor = self.getUtf8(ntInfo.descriptorIndex)

	return name, descriptor
}

//从常量表中取uft-8字符串
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
