package heap

import "jvm/classfile"

//类成员
type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

//复制类成员信息
func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}
