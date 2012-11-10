package xsd

import (
	"fmt"
	"path"
	"strings"
	"time"

	util "github.com/metaleap/go-util"
	ustr "github.com/metaleap/go-util/str"

	xsdt "github.com/metaleap/go-xsd/types"
)

func (me *All) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsElement.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Annotation) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemsAppInfo.makePkg(bag)
	me.hasElemsDocumentation.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Any) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *AnyAttribute) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *AppInfo) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Attribute) makePkg (bag *PkgBag) {
	var safeName, typeName, tmp, key, defVal, impName string
	var defName = "Default"
	me.elemBase.beforeMakePkg(bag)
	if len(me.Form) == 0 { me.Form = bag.Schema.AttributeFormDefault }
	me.hasElemsSimpleType.makePkg(bag)
	if len(me.Ref) > 0 {
		me.hasElemAnnotation.makePkg(bag)
	} else {
		safeName = bag.safeName(me.Name.String())
		if typeName = me.Type.String(); (len(typeName) == 0) && (len(me.SimpleTypes) > 0) {
			typeName = me.SimpleTypes[0].Name.String()
		} else {
			if len(typeName) == 0 { typeName = bag.xsdStringTypeRef() }
			typeName = bag.resolveTypeRef(typeName, &impName)
		}
		if defVal = me.Default; len(defVal) == 0 { defName, defVal = "Fixed", me.Fixed }
		if key = safeName + "_" + bag.safeName(typeName) + "_" + bag.safeName(defVal); len(perPkgState.attsCache[key]) == 0 {
			bag.impsUsed[impName] = true
			tmp = "_XsdHasAttr_" + key
			perPkgState.attsCache[key] = tmp
			bag.appendFmt(false, "type %v struct {", tmp)
			me.hasElemAnnotation.makePkg(bag)
			bag.appendFmt(false, "\t%v %v `xml:\"%v,attr\"`", safeName, typeName, util.Ifs((me.Form == "qualified") && len(bag.Schema.TargetNamespace) > 0, bag.Schema.TargetNamespace.String() + " ", "") + me.Name.String())
			bag.appendFmt(true, "}")
			if isPt := bag.isParseType(typeName); len(defVal) > 0 {
				bag.appendFmt(false, "\t//\tReturns the %v value for %v -- " + util.Ifs(isPt, "%v", "%#v"), defName, safeName, defVal)
				if isPt {
					bag.appendFmt(true, "\tfunc (me *%v) %v%v () %v { return %v(%v) }", tmp, safeName, defName, typeName, typeName, defVal)
				} else {
					bag.appendFmt(true, "\tfunc (me *%v) %v%v () %v { return %v(%#v) }", tmp, safeName, defName, typeName, typeName, defVal)
				}
			}
		}
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *AttributeGroup) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Choice) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsAny.makePkg(bag)
	me.hasElemsChoice.makePkg(bag)
	me.hasElemsElement.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsSequence.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *ComplexContent) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemExtensionComplexContent.makePkg(bag)
	me.hasElemRestrictionComplexContent.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *ComplexType) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemAll.makePkg(bag)
	me.hasElemChoice.makePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsSequence.makePkg(bag)
	me.hasElemComplexContent.makePkg(bag)
	me.hasElemSimpleContent.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Documentation) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	if len(me.CDATA) > 0 { var s, ln string;
		for _, ln = range ustr.Split(me.CDATA, "\n") { if s = strings.Trim(ln, " \t\r\n"); len(s) > 0 { bag.appendFmt(false, "//\t%s", s) } }
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *Element) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	me.hasElemUnique.makePkg(bag)
	me.hasElemsKey.makePkg(bag)
	me.hasElemComplexType.makePkg(bag)
	me.hasElemKeyRef.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *ExtensionComplexContent) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemAll.makePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsChoice.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsSequence.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *ExtensionSimpleContent) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Field) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Group) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemAll.makePkg(bag)
	me.hasElemChoice.makePkg(bag)
	me.hasElemSequence.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Import) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	var impName, impPath string
	var pos int
	me.hasElemAnnotation.makePkg(bag)
	for k, v := range bag.Schema.XMLNamespaces { if v == me.Namespace { impName = k; break } }
	if len(impName) > 0 {
		if pos, impPath = strings.Index(me.SchemaLocation.String(), protSep), me.SchemaLocation.String(); pos > 0 {
			impPath = impPath[pos + len(protSep) :]
		} else {
			impPath = path.Join(path.Dir(bag.Schema.loadUri), impPath)
		}
		impPath = path.Join(path.Dir(impPath), goPkgPrefix + path.Base(impPath) + goPkgSuffix)
		bag.imports[impName] = path.Join(PkgGen.BasePath, impPath)
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *Key) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemField.makePkg(bag)
	me.hasElemSelector.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *KeyRef) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemField.makePkg(bag)
	me.hasElemSelector.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *List) makePkg (bag *PkgBag) {
	var rtr = bag.resolveTypeRef(me.ItemType.String(), nil)
	var safeName string
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	if len(rtr) == 0 {
		rtr = me.SimpleTypes[0].Name.String()
		if len(rtr) == 0 {
			rtr = bag.safeName(bag.Stacks.CurSimpleType().Name.String() + "SubList")
			me.SimpleTypes[0].Name = xsdt.NCName(rtr)
		}
	}
	bag.impsUsed[bag.impName] = true
	safeName = bag.safeName(bag.Stacks.CurSimpleType().Name.String())
	bag.appendFmt(false, "\t//\tThis %v contains a whitespace-separated list of %v values. Its Values() method returns a slice with all elements in that list.", safeName, rtr)
	if bag.isParseType(rtr) {
		bag.appendFmt(false, `	func (me %v) Values () (list []%v) {
		var btv = new(%v)
		var svals = xsdt.ListValues(string(me))
		list = make([]%v, len(svals))
		for i, s := range svals { btv.SetFromString(s); list[i] = *btv }
		return
	}
		`, safeName, rtr, rtr, rtr)
	} else {
		bag.appendFmt(false, `	func (me %v) Values () (list []%v) {
		var svals = xsdt.ListValues(string(me))
		list = make([]%v, len(svals))
		for i, s := range svals { list[i] = %v(s) }
		return
	}
		`, safeName, rtr, rtr, rtr)
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *Notation) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	bag.appendFmt(false, "Notations.Add(%#v, %#v, %#v, %#v)", me.Id, me.Name, me.Public, me.System)
	me.elemBase.afterMakePkg(bag)
}

func (me *Redefine) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.hasElemsComplexType.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionComplexContent) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemAll.makePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsChoice.makePkg(bag)
	me.hasElemsSequence.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleContent) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	me.hasElemLength.makePkg(bag)
	me.hasElemPattern.makePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsEnumeration.makePkg(bag)
	me.hasElemFractionDigits.makePkg(bag)
	me.hasElemMaxExclusive.makePkg(bag)
	me.hasElemMaxInclusive.makePkg(bag)
	me.hasElemMaxLength.makePkg(bag)
	me.hasElemMinExclusive.makePkg(bag)
	me.hasElemMinInclusive.makePkg(bag)
	me.hasElemMinLength.makePkg(bag)
	me.hasElemTotalDigits.makePkg(bag)
	me.hasElemWhiteSpace.makePkg(bag)
	me.hasElemsAnyAttribute.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleEnumeration) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	safeName := bag.safeName(bag.Stacks.CurSimpleType().Name.String())
	bag.appendFmt(false, "\t//\tReturns true if the value of this enumerated %v is %#v.", safeName, me.Value)
	bag.appendFmt(true, "\tfunc (me %v) Is%v () bool { return me == %#v }", safeName, bag.safeName(me.Value), me.Value)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleFractionDigits) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleLength) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleMaxExclusive) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleMaxInclusive) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleMaxLength) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleMinExclusive) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleMinInclusive) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleMinLength) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimplePattern) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleTotalDigits) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleType) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	me.hasElemLength.makePkg(bag)
	me.hasElemPattern.makePkg(bag)
	me.hasElemsEnumeration.makePkg(bag)
	me.hasElemFractionDigits.makePkg(bag)
	me.hasElemMaxExclusive.makePkg(bag)
	me.hasElemMaxInclusive.makePkg(bag)
	me.hasElemMaxLength.makePkg(bag)
	me.hasElemMinExclusive.makePkg(bag)
	me.hasElemMinInclusive.makePkg(bag)
	me.hasElemMinLength.makePkg(bag)
	me.hasElemTotalDigits.makePkg(bag)
	me.hasElemWhiteSpace.makePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *RestrictionSimpleWhiteSpace) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Schema) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	var impPos int
	bag.reinit(); bag.now = time.Now().UnixNano(); bag.snow = fmt.Sprintf("%v", bag.now)
	for k, _ := range me.XMLNamespaces { if k == bag.impName { bag.impName += bag.snow } }
	bag.imports[bag.impName] = "github.com/metaleap/go-xsd/types"
	bag.ParseTypes = []string {}
	for _, pt := range []string { "Boolean", "Byte", "Double", "Float", "Int", "Integer", "Long", "NegativeInteger", "NonNegativeInteger", "NonPositiveInteger", "PositiveInteger", "Short", "UnsignedByte", "UnsignedInt", "UnsignedLong", "UnsignedShort" } {
		bag.ParseTypes = append(bag.ParseTypes, bag.impName + "." + pt)
	}
	// for _, im := range []string { "strings" } { bag.imports[im] = "" }
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsImport.makePkg(bag)
	impPos = len(bag.lines) + 1
	bag.append("import (", ")", "")
	if me.XSDParentSchema == nil {
		perPkgState.anonCount = 0
		perPkgState.attsCache = map[string]string {}
		bag.append("type _XsdHasCdata struct { CDATA string `xml:\",chardata\"` }", "")
	}
	me.hasElemsSimpleType.makePkg(bag)
	me.hasElemsAttribute.makePkg(bag)
	me.hasElemsElement.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsNotation.makePkg(bag)
	me.hasElemsRedefine.makePkg(bag)
	me.hasElemsAttributeGroup.makePkg(bag)
	me.hasElemsComplexType.makePkg(bag)
	for impName, impPath := range bag.imports {
		if bag.impsUsed[impName] {
			if len(impPath) > 0 {
				bag.insertFmt(impPos, "\t%v \"%v\"", impName, impPath)
			} else {
				bag.insertFmt(impPos, "\t\"%v\"", impName)
			}
		}
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *Selector) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *Sequence) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsAny.makePkg(bag)
	me.hasElemsChoice.makePkg(bag)
	me.hasElemsElement.makePkg(bag)
	me.hasElemsGroup.makePkg(bag)
	me.hasElemsSequence.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *SimpleContent) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemExtensionSimpleContent.makePkg(bag)
	me.hasElemRestrictionSimpleContent.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}

func (me *SimpleType) makePkg (bag *PkgBag) {
	var typeName = me.Name
	var baseType, safeName = "", ""
	var resolve = true
	if len(typeName) == 0 { typeName = xsdt.NCName(bag.AnonName()); me.hasAttrName.Name = typeName }
	me.elemBase.beforeMakePkg(bag)
	bag.Stacks.SimpleType.Push(me)
	safeName = bag.safeName(typeName.String())
	me.hasElemAnnotation.makePkg(bag)
	if me.RestrictionSimpleType != nil {
		if baseType = me.RestrictionSimpleType.Base.String(); (len(baseType) == 0) && (len(me.RestrictionSimpleType.SimpleTypes) > 0) {
			resolve, baseType = false, me.RestrictionSimpleType.SimpleTypes[0].Name.String()
		}
	}
	if len(baseType) == 0 { baseType = bag.xsdStringTypeRef() }
	if resolve { baseType = bag.resolveTypeRef(baseType, nil) }
	if bag.isParseType(baseType) { bag.ParseTypes = append(bag.ParseTypes, safeName) }
	bag.appendFmt(false, `type %s %s

	//	If base type is string, simply sets the value from the specified string. If base type is parseable (bool or numeric), sets the value by parsing the specified string.
	func (me *%s) SetFromString (s string) { (*%v)(me).SetFromString(s) }

	//	If base type is string, returns the underlying value, otherwise returns the string representation of the underlying value.
	func (me %s) String () string { return %v(me).String() }
	`, safeName, baseType, safeName, baseType, safeName, baseType)
	me.hasElemRestrictionSimpleType.makePkg(bag)
	me.hasElemList.makePkg(bag)
	me.hasElemUnion.makePkg(bag)
	bag.Stacks.SimpleType.Pop()
	me.elemBase.afterMakePkg(bag)
}

func (me *Union) makePkg (bag *PkgBag) {
	var memberTypes []string
	var rtn, rtnSafeName, safeName string
	var isParseType = false
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsSimpleType.makePkg(bag)
	memberTypes = ustr.Split(me.MemberTypes, " ")
	for _, st := range me.SimpleTypes { memberTypes = append(memberTypes, st.Name.String()) }
	for _, mt := range memberTypes {
		rtn = bag.resolveTypeRef(mt, nil)
		safeName, rtnSafeName = bag.safeName(bag.Stacks.CurSimpleType().Name.String()), bag.safeName(rtn)
		bag.appendFmt(false, "\t//\t%v is an XSD union type of several types. This is a simple type conversion to %v, but keep in mind the actual value may or may not be a valid %v value.", safeName, rtnSafeName, rtnSafeName)
		if isParseType = bag.isParseType(rtn); isParseType {
			bag.appendFmt(true, "\tfunc (me %v) To%v () %v { var x = new(%v); x.SetFromString(me.String()); return *x }", safeName, rtnSafeName, rtn, rtn)
		} else {
			bag.appendFmt(true, "\tfunc (me %v) To%v () %v { return %v(me) }", safeName, rtnSafeName, rtn, rtn)
		}
	}
	me.elemBase.afterMakePkg(bag)
}

func (me *Unique) makePkg (bag *PkgBag) {
	me.elemBase.beforeMakePkg(bag)
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemField.makePkg(bag)
	me.hasElemSelector.makePkg(bag)
	me.elemBase.afterMakePkg(bag)
}
