// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#unix
package unix_def

import (
	"encoding/xml"
	"github.com/complianceascode/librescap/pkg/scap/models/oval"
	"github.com/complianceascode/librescap/pkg/scap/models/oval_def"
	"github.com/complianceascode/librescap/pkg/scap/models/xml_dsig"
)

// Element
type DnscacheTest struct {
	XMLName xml.Name `xml:dnscache_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type DnscacheObject struct {
	XMLName xml.Name `xml:dnscache_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type DnscacheState struct {
	XMLName xml.Name `xml:dnscache_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	DomainName *oval_def.EntityStateStringType `xml:"domain_name"`

	Ttl *oval_def.EntityStateIntType `xml:"ttl"`

	IpAddress *oval_def.EntityStateIPAddressStringType `xml:"ip_address"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type FileTest struct {
	XMLName xml.Name `xml:file_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type FileObject struct {
	XMLName xml.Name `xml:file_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type FileState struct {
	XMLName xml.Name `xml:file_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Filename *oval_def.EntityStateStringType `xml:"filename"`

	Type *oval_def.EntityStateStringType `xml:"type"`

	GroupId *FileStateGroupId `xml:"group_id"`

	UserId *FileStateUserId `xml:"user_id"`

	ATime *FileStateATime `xml:"a_time"`

	CTime *FileStateCTime `xml:"c_time"`

	MTime *FileStateMTime `xml:"m_time"`

	Size *oval_def.EntityStateIntType `xml:"size"`

	Suid *oval_def.EntityStateBoolType `xml:"suid"`

	Sgid *oval_def.EntityStateBoolType `xml:"sgid"`

	Sticky *oval_def.EntityStateBoolType `xml:"sticky"`

	Uread *oval_def.EntityStateBoolType `xml:"uread"`

	Uwrite *oval_def.EntityStateBoolType `xml:"uwrite"`

	Uexec *oval_def.EntityStateBoolType `xml:"uexec"`

	Gread *oval_def.EntityStateBoolType `xml:"gread"`

	Gwrite *oval_def.EntityStateBoolType `xml:"gwrite"`

	Gexec *oval_def.EntityStateBoolType `xml:"gexec"`

	Oread *oval_def.EntityStateBoolType `xml:"oread"`

	Owrite *oval_def.EntityStateBoolType `xml:"owrite"`

	Oexec *oval_def.EntityStateBoolType `xml:"oexec"`

	HasExtendedAcl *oval_def.EntityStateBoolType `xml:"has_extended_acl"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type FileextendedattributeTest struct {
	XMLName xml.Name `xml:fileextendedattribute_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type FileextendedattributeObject struct {
	XMLName xml.Name `xml:fileextendedattribute_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type FileextendedattributeState struct {
	XMLName xml.Name `xml:fileextendedattribute_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Filename *oval_def.EntityStateStringType `xml:"filename"`

	AttributeName *oval_def.EntityStateStringType `xml:"attribute_name"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type GconfTest struct {
	XMLName xml.Name `xml:gconf_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type GconfObject struct {
	XMLName xml.Name `xml:gconf_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type GconfState struct {
	XMLName xml.Name `xml:gconf_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Key *oval_def.EntityStateStringType `xml:"key"`

	Source *oval_def.EntityStateStringType `xml:"source"`

	Type *EntityStateGconfTypeType `xml:"type"`

	IsWritable *oval_def.EntityStateBoolType `xml:"is_writable"`

	ModUser *oval_def.EntityStateStringType `xml:"mod_user"`

	ModTime *oval_def.EntityStateIntType `xml:"mod_time"`

	IsDefault *oval_def.EntityStateBoolType `xml:"is_default"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type InetdTest struct {
	XMLName xml.Name `xml:inetd_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type InetdObject struct {
	XMLName xml.Name `xml:inetd_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type InetdState struct {
	XMLName xml.Name `xml:inetd_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Protocol *oval_def.EntityStateStringType `xml:"protocol"`

	ServiceName *oval_def.EntityStateStringType `xml:"service_name"`

	ServerProgram *oval_def.EntityStateStringType `xml:"server_program"`

	ServerArguments *oval_def.EntityStateStringType `xml:"server_arguments"`

	EndpointType *EntityStateEndpointType `xml:"endpoint_type"`

	ExecAsUser *oval_def.EntityStateStringType `xml:"exec_as_user"`

	WaitStatus *EntityStateWaitStatusType `xml:"wait_status"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type InterfaceTest struct {
	XMLName xml.Name `xml:interface_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type InterfaceObject struct {
	XMLName xml.Name `xml:interface_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type InterfaceState struct {
	XMLName xml.Name `xml:interface_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Type *EntityStateInterfaceType `xml:"type"`

	HardwareAddr *oval_def.EntityStateStringType `xml:"hardware_addr"`

	InetAddr *oval_def.EntityStateIPAddressStringType `xml:"inet_addr"`

	BroadcastAddr *oval_def.EntityStateIPAddressStringType `xml:"broadcast_addr"`

	Netmask *oval_def.EntityStateIPAddressStringType `xml:"netmask"`

	Flag *oval_def.EntityStateStringType `xml:"flag"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PasswordTest struct {
	XMLName xml.Name `xml:password_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PasswordObject struct {
	XMLName xml.Name `xml:password_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PasswordState struct {
	XMLName xml.Name `xml:password_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Username *oval_def.EntityStateStringType `xml:"username"`

	Password *oval_def.EntityStateStringType `xml:"password"`

	UserId *PasswordStateUserId `xml:"user_id"`

	GroupId *PasswordStateGroupId `xml:"group_id"`

	Gcos *oval_def.EntityStateStringType `xml:"gcos"`

	HomeDir *oval_def.EntityStateStringType `xml:"home_dir"`

	LoginShell *oval_def.EntityStateStringType `xml:"login_shell"`

	LastLogin *oval_def.EntityStateIntType `xml:"last_login"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type ProcessTest struct {
	XMLName xml.Name `xml:process_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type ProcessObject struct {
	XMLName xml.Name `xml:process_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Command oval_def.EntityObjectStringType `xml:"command"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type ProcessState struct {
	XMLName xml.Name `xml:process_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Command *oval_def.EntityStateStringType `xml:"command"`

	ExecTime *oval_def.EntityStateStringType `xml:"exec_time"`

	Pid *oval_def.EntityStateIntType `xml:"pid"`

	Ppid *oval_def.EntityStateIntType `xml:"ppid"`

	Priority *oval_def.EntityStateIntType `xml:"priority"`

	Ruid *oval_def.EntityStateIntType `xml:"ruid"`

	SchedulingClass *oval_def.EntityStateStringType `xml:"scheduling_class"`

	StartTime *oval_def.EntityStateStringType `xml:"start_time"`

	Tty *oval_def.EntityStateStringType `xml:"tty"`

	UserId *oval_def.EntityStateIntType `xml:"user_id"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Process58Test struct {
	XMLName xml.Name `xml:process58_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Process58Object struct {
	XMLName xml.Name `xml:process58_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Process58State struct {
	XMLName xml.Name `xml:process58_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	CommandLine *oval_def.EntityStateStringType `xml:"command_line"`

	ExecTime *oval_def.EntityStateStringType `xml:"exec_time"`

	Pid *oval_def.EntityStateIntType `xml:"pid"`

	Ppid *oval_def.EntityStateIntType `xml:"ppid"`

	Priority *oval_def.EntityStateIntType `xml:"priority"`

	Ruid *oval_def.EntityStateIntType `xml:"ruid"`

	SchedulingClass *oval_def.EntityStateStringType `xml:"scheduling_class"`

	StartTime *oval_def.EntityStateStringType `xml:"start_time"`

	Tty *oval_def.EntityStateStringType `xml:"tty"`

	UserId *oval_def.EntityStateIntType `xml:"user_id"`

	ExecShield *oval_def.EntityStateBoolType `xml:"exec_shield"`

	Loginuid *oval_def.EntityStateIntType `xml:"loginuid"`

	PosixCapability *EntityStateCapabilityType `xml:"posix_capability"`

	SelinuxDomainLabel *oval_def.EntityStateStringType `xml:"selinux_domain_label"`

	SessionId *oval_def.EntityStateIntType `xml:"session_id"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type RoutingtableTest struct {
	XMLName xml.Name `xml:routingtable_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type RoutingtableObject struct {
	XMLName xml.Name `xml:routingtable_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type RoutingtableState struct {
	XMLName xml.Name `xml:routingtable_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Destination *oval_def.EntityStateIPAddressType `xml:"destination"`

	Gateway *oval_def.EntityStateIPAddressType `xml:"gateway"`

	Flags *EntityStateRoutingTableFlagsType `xml:"flags"`

	InterfaceName *oval_def.EntityStateStringType `xml:"interface_name"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type RunlevelTest struct {
	XMLName xml.Name `xml:runlevel_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type RunlevelObject struct {
	XMLName xml.Name `xml:runlevel_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type RunlevelState struct {
	XMLName xml.Name `xml:runlevel_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	ServiceName *oval_def.EntityStateStringType `xml:"service_name"`

	Runlevel *oval_def.EntityStateStringType `xml:"runlevel"`

	Start *oval_def.EntityStateBoolType `xml:"start"`

	Kill *oval_def.EntityStateBoolType `xml:"kill"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SccsTest struct {
	XMLName xml.Name `xml:sccs_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SccsObject struct {
	XMLName xml.Name `xml:sccs_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SccsState struct {
	XMLName xml.Name `xml:sccs_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	Path *oval_def.EntityStateStringType `xml:"path"`

	Filename *oval_def.EntityStateStringType `xml:"filename"`

	ModuleName *oval_def.EntityStateStringType `xml:"module_name"`

	ModuleType *oval_def.EntityStateStringType `xml:"module_type"`

	Release *oval_def.EntityStateStringType `xml:"release"`

	Level *oval_def.EntityStateStringType `xml:"level"`

	Branch *oval_def.EntityStateStringType `xml:"branch"`

	Sequence *oval_def.EntityStateStringType `xml:"sequence"`

	WhatString *oval_def.EntityStateStringType `xml:"what_string"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type ShadowTest struct {
	XMLName xml.Name `xml:shadow_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type ShadowObject struct {
	XMLName xml.Name `xml:shadow_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type ShadowState struct {
	XMLName xml.Name `xml:shadow_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Username *oval_def.EntityStateStringType `xml:"username"`

	Password *oval_def.EntityStateStringType `xml:"password"`

	ChgLst *ShadowStateChgLst `xml:"chg_lst"`

	ChgAllow *ShadowStateChgAllow `xml:"chg_allow"`

	ChgReq *ShadowStateChgReq `xml:"chg_req"`

	ExpWarn *ShadowStateExpWarn `xml:"exp_warn"`

	ExpInact *ShadowStateExpInact `xml:"exp_inact"`

	ExpDate *ShadowStateExpDate `xml:"exp_date"`

	Flag *ShadowStateFlag `xml:"flag"`

	EncryptMethod *EntityStateEncryptMethodType `xml:"encrypt_method"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SymlinkTest struct {
	XMLName xml.Name `xml:symlink_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SymlinkObject struct {
	XMLName xml.Name `xml:symlink_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SymlinkState struct {
	XMLName xml.Name `xml:symlink_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Filepath *oval_def.EntityStateStringType `xml:"filepath"`

	CanonicalPath *oval_def.EntityStateStringType `xml:"canonical_path"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SysctlTest struct {
	XMLName xml.Name `xml:sysctl_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SysctlObject struct {
	XMLName xml.Name `xml:sysctl_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type SysctlState struct {
	XMLName xml.Name `xml:sysctl_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Name *oval_def.EntityStateStringType `xml:"name"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type UnameTest struct {
	XMLName xml.Name `xml:uname_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type UnameObject struct {
	XMLName xml.Name `xml:uname_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type UnameState struct {
	XMLName xml.Name `xml:uname_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	MachineClass *oval_def.EntityStateStringType `xml:"machine_class"`

	NodeName *oval_def.EntityStateStringType `xml:"node_name"`

	OsName *oval_def.EntityStateStringType `xml:"os_name"`

	OsRelease *oval_def.EntityStateStringType `xml:"os_release"`

	OsVersion *oval_def.EntityStateStringType `xml:"os_version"`

	ProcessorType *oval_def.EntityStateStringType `xml:"processor_type"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type XinetdTest struct {
	XMLName xml.Name `xml:xinetd_test`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	CheckExistence string `xml:"check_existence,attr"`

	Check string `xml:"check,attr"`

	StateOperator string `xml:"state_operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Object oval_def.ObjectRefType `xml:"object"`

	State []oval_def.StateRefType `xml:"state"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type XinetdObject struct {
	XMLName xml.Name `xml:xinetd_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type XinetdState struct {
	XMLName xml.Name `xml:xinetd_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Protocol *oval_def.EntityStateStringType `xml:"protocol"`

	ServiceName *oval_def.EntityStateStringType `xml:"service_name"`

	Flags *oval_def.EntityStateStringType `xml:"flags"`

	NoAccess *oval_def.EntityStateStringType `xml:"no_access"`

	OnlyFrom *oval_def.EntityStateIPAddressStringType `xml:"only_from"`

	Port *oval_def.EntityStateIntType `xml:"port"`

	Server *oval_def.EntityStateStringType `xml:"server"`

	ServerArguments *oval_def.EntityStateStringType `xml:"server_arguments"`

	SocketType *oval_def.EntityStateStringType `xml:"socket_type"`

	Type *EntityStateXinetdTypeStatusType `xml:"type"`

	User *oval_def.EntityStateStringType `xml:"user"`

	Wait *oval_def.EntityStateBoolType `xml:"wait"`

	Disabled *oval_def.EntityStateBoolType `xml:"disabled"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type FileStateGroupId struct {
	XMLName xml.Name `xml:group_id`
}

// Element
type FileStateUserId struct {
	XMLName xml.Name `xml:user_id`
}

// Element
type FileStateATime struct {
	XMLName xml.Name `xml:a_time`
}

// Element
type FileStateCTime struct {
	XMLName xml.Name `xml:c_time`
}

// Element
type FileStateMTime struct {
	XMLName xml.Name `xml:m_time`
}

// Element
type PasswordStateUserId struct {
	XMLName xml.Name `xml:user_id`
}

// Element
type PasswordStateGroupId struct {
	XMLName xml.Name `xml:group_id`
}

// Element
type ShadowStateChgLst struct {
	XMLName xml.Name `xml:chg_lst`
}

// Element
type ShadowStateChgAllow struct {
	XMLName xml.Name `xml:chg_allow`
}

// Element
type ShadowStateChgReq struct {
	XMLName xml.Name `xml:chg_req`
}

// Element
type ShadowStateExpWarn struct {
	XMLName xml.Name `xml:exp_warn`
}

// Element
type ShadowStateExpInact struct {
	XMLName xml.Name `xml:exp_inact`
}

// Element
type ShadowStateExpDate struct {
	XMLName xml.Name `xml:exp_date`
}

// Element
type ShadowStateFlag struct {
	XMLName xml.Name `xml:flag`
}

// XSD ComplexType declarations

type FileBehaviors struct {
	MaxDepth string `xml:"max_depth,attr"`

	Recurse string `xml:"recurse,attr"`

	RecurseDirection string `xml:"recurse_direction,attr"`

	RecurseFileSystem string `xml:"recurse_file_system,attr"`
}

type EntityStateCapabilityType struct {
}

type EntityStateEndpointType struct {
}

type EntityStateGconfTypeType struct {
}

type EntityStateRoutingTableFlagsType struct {
}

type EntityStateXinetdTypeStatusType struct {
}

type EntityStateWaitStatusType struct {
}

type EntityStateEncryptMethodType struct {
}

type EntityStateInterfaceType struct {
}
