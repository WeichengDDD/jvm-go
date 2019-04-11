package classfile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

//创建属性
func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
	case "ConstantValue":
	case "Deprecated":
	case "Exceptions":
	case "LineNumberTable":
	case "LocalVariableTable":
	case "SourceFile":
	case "Synthetic":
	default:

	}
	return nil
}

//读取属性字段
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)

	return attrInfo
}

//读取属性列表
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}

	return attributes
}
