package gen

import (
	"fmt"
	"strings"

	"github.com/progrium/macschema/schema"
)

type classBuilder struct {
	Class           schema.Class
	Imports         []PackageContents
	consumedImports map[Import]bool
}

func (cb *classBuilder) EachTypeMethod(f func(schema.Method)) {
	f(schema.Method{
		Name:   "alloc",
		Return: schema.DataType{Name: "instancetype"},
	})
	for _, m := range cb.Class.TypeMethods {
		f(m)
	}
	for _, p := range cb.Class.TypeProperties {
		for _, m := range propertyMethods(p) {
			f(m)
		}
	}
}

func (cb *classBuilder) EachInstanceMethod(f func(schema.Method)) {
	seen := make(map[string]bool)
	for _, m := range cb.Class.InstanceMethods {
		seen[m.Name] = true
		f(m)
	}
	if !seen["init"] {
		f(schema.Method{
			Name:   "init",
			Return: schema.DataType{Name: "instancetype"},
		})
	}
	for _, p := range cb.Class.InstanceProperties {
		func() {
			defer ignoreIfUnimplemented(fmt.Sprintf("%s.%s", cb.Class.Name, p.Name))
			for _, m := range propertyMethods(p) {
				// properties sometimes specify a getter or setter method that is also
				// in `InstanceMethods`, so skip those if they've already been handled
				if !seen[m.Name] {
					f(m)
				}
			}
		}()
	}
}

func (cb *classBuilder) instanceMethod(method schema.Method) MethodDef {
	r := MethodDef{
		Name:        toExportedName(selectorNameToGoIdent(method.Name)),
		WrappedFunc: cb.cgoWrapperFunc(method, false),
	}
	if isInstanceType(method.Return) {
		r.Name += "_as" + cb.Class.Name
	}
	return r
}

func (cb *classBuilder) msgSend(method schema.Method, isTypeMethod bool) CGoMsgSend {
	msg := CGoMsgSend{
		Name:   msgSendFuncName(cb.Class, method.Name, isTypeMethod),
		Class:  cb.Class.Name,
		Return: cb.toMsgSendReturn(method.Return),
	}
	for i, key := range strings.SplitAfter(method.Name, ":") {
		if key == "" {
			continue
		}
		part := SelectorPart{
			Key: key,
		}
		if i < len(method.Args) {
			arg := method.Args[i]
			typ := cb.mapType(arg.Type)
			part.Arg = &Arg{
				Name: arg.Name,
				Type: typ.CType,
			}
		}
		msg.Selector = append(msg.Selector, part)
	}
	return msg
}

func (cb *classBuilder) toMsgSendReturn(dt schema.DataType) string {
	if isVoid(dt) {
		return "void"
	}
	typ := cb.mapType(dt)
	return typ.CType
}

func (cb *classBuilder) cgoWrapperFunc(method schema.Method, isTypeMethod bool) CGoWrapperFunc {
	r := CGoWrapperFunc{
		Name:    msgSendFuncName(cb.Class, method.Name, isTypeMethod),
		Args:    []CGoWrapperArg{},
		Returns: cb.toCGoWrapperReturn(method.Return),
	}
	for _, arg := range method.Args {
		typ := cb.mapType(arg.Type)
		goType := typ.GoSimpleRefType
		if goType == "" {
			goType = typ.GoType
		}
		r.Args = append(r.Args, CGoWrapperArg{
			Name:     arg.Name,
			Type:     goType,
			ToCGoFmt: typ.ToCGoFmt,
		})
	}
	return r
}

func (cb *classBuilder) toCGoWrapperReturn(dt schema.DataType) []CGoWrapperReturn {
	if isVoid(dt) {
		return nil
	}
	typ := cb.mapType(dt)
	return []CGoWrapperReturn{
		{Type: typ.GoType, FromCGoFmt: typ.FromCGoFmt},
	}
}