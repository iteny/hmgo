package reflectm

import (
	"reflect"
	"sync"
)

// FieldInfo is metadata for a struct field.
// FieldInfo 是结构体字段元数据。
type FieldInfo struct {
	Index    []int
	Path     string
	Field    reflect.StructField
	Zero     reflect.Value
	Name     string
	Options  map[string]string
	Embedded bool
	Children []*FieldInfo
	Parent   *FieldInfo
}

// A StructMap is an index of field metadata for a struct.
// StructMap是一个结构体字段元数据索引
type StructMap struct {
	Tree  *FieldInfo
	Index []*FieldInfo
	Paths map[string]*FieldInfo
	Names map[string]*FieldInfo
}

// Mapper is a general purpose mapper of names to struct fields.  A Mapper
// behaves like most marshallers in the standard library, obeying a field tag
// for name mapping but also providing a basic transform function.
// Mapper是一个名为结构体字段的常规目标映射器。Mapper表现的像标准库中大多数装配器一样，遵循字段标签为了名称映射，但也提供了基本转换功能。
type Mapper struct {
	cache      map[reflect.Type]*StructMap
	tagName    string
	tagMapFunc func(string) string
	mapFunc    func(string) string
	mutex      sync.Mutex
}

// NewMapperFunc returns a new mapper which optionally obeys a field tag and
// a struct field name mapper func given by f.  Tags will take precedence, but
// for any other field, the mapped name will be f(field.Name)
// NewMapperFunc返回一个新的映射器，可选的遵循一个字段标签并且给予f结构体字段名称映射函数。标签将优先，但是对于其他任何字段，映射器名称将为f(field.Name)
func NewMapperFunc(tagName string, f func(string) string) *Mapper {
	return &Mapper{
		cache:   make(map[reflect.Type]*StructMap),
		tagName: tagName,
		mapFunc: f,
	}
}
