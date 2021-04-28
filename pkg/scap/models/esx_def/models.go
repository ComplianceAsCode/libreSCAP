// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-definitions-5#esx
package esx_def

import (
	"encoding/xml"
	"github.com/complianceascode/librescap/pkg/scap/models/oval"
	"github.com/complianceascode/librescap/pkg/scap/models/oval_def"
	"github.com/complianceascode/librescap/pkg/scap/models/xml_dsig"
)

// Element
type Patch56Test struct {
	XMLName xml.Name `xml:patch56_test`

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
type Patch56Object struct {
	XMLName xml.Name `xml:patch56_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type Patch56State struct {
	XMLName xml.Name `xml:patch56_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	PatchName *oval_def.EntityStateStringType `xml:"patch_name"`

	KnowledgeBaseId *oval_def.EntityStateIntType `xml:"knowledge_base_id"`

	BundleId *oval_def.EntityStateIntType `xml:"bundle_id"`

	Classification *EntityStateClassificationType `xml:"classification"`

	SupportLevel *EntityStateSupportLevelType `xml:"support_level"`

	Status *oval_def.EntityStateBoolType `xml:"status"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PatchTest struct {
	XMLName xml.Name `xml:patch_test`

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
type PatchObject struct {
	XMLName xml.Name `xml:patch_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type PatchState struct {
	XMLName xml.Name `xml:patch_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	PatchNumber *oval_def.EntityStateStringType `xml:"patch_number"`

	Status *oval_def.EntityStateBoolType `xml:"status"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type VersionTest struct {
	XMLName xml.Name `xml:version_test`

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
type VersionObject struct {
	XMLName xml.Name `xml:version_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type VersionState struct {
	XMLName xml.Name `xml:version_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Release *oval_def.EntityStateVersionType `xml:"release"`

	Build *oval_def.EntityStateIntType `xml:"build"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type VisdkmanagedobjectTest struct {
	XMLName xml.Name `xml:visdkmanagedobject_test`

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
type VisdkmanagedobjectObject struct {
	XMLName xml.Name `xml:visdkmanagedobject_object`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Set oval_def.Set `xml:"set"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// Element
type VisdkmanagedobjectState struct {
	XMLName xml.Name `xml:visdkmanagedobject_state`

	Id string `xml:"id,attr"`

	Version string `xml:"version,attr"`

	Operator string `xml:"operator,attr"`

	Comment string `xml:"comment,attr"`

	Deprecated string `xml:"deprecated,attr"`

	Property *oval_def.EntityStateStringType `xml:"property"`

	Value *oval_def.EntityStateAnySimpleType `xml:"value"`

	Signature *xml_dsig.Signature `xml:"Signature"`

	Notes *oval.Notes `xml:"notes"`
}

// XSD ComplexType declarations

type Patch56Behaviors struct {
	Supersedence string `xml:"supersedence,attr"`
}

type PatchBehaviors struct {
	Supersedence string `xml:"supersedence,attr"`
}

type ViSdkManagedEntityBehaviors struct {
	ManagedEntityType string `xml:"managed_entity_type,attr"`
}

type EntityStateClassificationType struct {
}

type EntityStateSupportLevelType struct {
}
