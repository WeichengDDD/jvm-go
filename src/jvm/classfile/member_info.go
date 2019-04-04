package classfile

import "reflect"

type MemberInfo struct {
	cp ConstantPool
	accessFlags uint16
	nameIndex uint16
	descriptorIndex uint16
	attributes []AttributeInfo
}

func readAttributes()  {

}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp: cp,
		accessFlags: reader.readUint16()
		nameIndex: reader.readUint16()
		descriptorIndex: reader.readUint16()
		attributes:
	}
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
	}

	return members
}