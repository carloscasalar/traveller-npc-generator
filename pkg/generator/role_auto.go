// Code generated by "enumer -type=Role -output=role_auto.go -trimprefix=Role -transform=snake"; DO NOT EDIT.

package generator

import (
	"fmt"
)

const _RoleName = "pilotnavigatorengineerstewardmedicmarinegunnerscouttechnicianleaderdiplomatentertainertraderthug"

var _RoleIndex = [...]uint8{0, 5, 14, 22, 29, 34, 40, 46, 51, 61, 67, 75, 86, 92, 96}

func (i Role) String() string {
	if i < 0 || i >= Role(len(_RoleIndex)-1) {
		return fmt.Sprintf("Role(%d)", i)
	}
	return _RoleName[_RoleIndex[i]:_RoleIndex[i+1]]
}

var _RoleValues = []Role{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

var _RoleNameToValueMap = map[string]Role{
	_RoleName[0:5]:   0,
	_RoleName[5:14]:  1,
	_RoleName[14:22]: 2,
	_RoleName[22:29]: 3,
	_RoleName[29:34]: 4,
	_RoleName[34:40]: 5,
	_RoleName[40:46]: 6,
	_RoleName[46:51]: 7,
	_RoleName[51:61]: 8,
	_RoleName[61:67]: 9,
	_RoleName[67:75]: 10,
	_RoleName[75:86]: 11,
	_RoleName[86:92]: 12,
	_RoleName[92:96]: 13,
}

// RoleString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func RoleString(s string) (Role, error) {
	if val, ok := _RoleNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Role values", s)
}

// RoleValues returns all values of the enum
func RoleValues() []Role {
	return _RoleValues
}

// IsARole returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Role) IsARole() bool {
	for _, v := range _RoleValues {
		if i == v {
			return true
		}
	}
	return false
}
