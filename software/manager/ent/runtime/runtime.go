// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/niwla23/lagersystem/manager/ent/box"
	"github.com/niwla23/lagersystem/manager/ent/part"
	"github.com/niwla23/lagersystem/manager/ent/position"
	"github.com/niwla23/lagersystem/manager/ent/property"
	"github.com/niwla23/lagersystem/manager/ent/schema"
	"github.com/niwla23/lagersystem/manager/ent/section"
	"github.com/niwla23/lagersystem/manager/ent/tag"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	boxFields := schema.Box{}.Fields()
	_ = boxFields
	// boxDescCreatedAt is the schema descriptor for createdAt field.
	boxDescCreatedAt := boxFields[0].Descriptor()
	// box.DefaultCreatedAt holds the default value on creation for the createdAt field.
	box.DefaultCreatedAt = boxDescCreatedAt.Default.(func() time.Time)
	partHooks := schema.Part{}.Hooks()
	part.Hooks[0] = partHooks[0]
	partFields := schema.Part{}.Fields()
	_ = partFields
	// partDescCreatedAt is the schema descriptor for createdAt field.
	partDescCreatedAt := partFields[0].Descriptor()
	// part.DefaultCreatedAt holds the default value on creation for the createdAt field.
	part.DefaultCreatedAt = partDescCreatedAt.Default.(func() time.Time)
	// partDescName is the schema descriptor for name field.
	partDescName := partFields[1].Descriptor()
	// part.NameValidator is a validator for the "name" field. It is called by the builders before save.
	part.NameValidator = partDescName.Validators[0].(func(string) error)
	positionFields := schema.Position{}.Fields()
	_ = positionFields
	// positionDescCreatedAt is the schema descriptor for createdAt field.
	positionDescCreatedAt := positionFields[0].Descriptor()
	// position.DefaultCreatedAt holds the default value on creation for the createdAt field.
	position.DefaultCreatedAt = positionDescCreatedAt.Default.(func() time.Time)
	// positionDescPositionId is the schema descriptor for positionId field.
	positionDescPositionId := positionFields[1].Descriptor()
	// position.PositionIdValidator is a validator for the "positionId" field. It is called by the builders before save.
	position.PositionIdValidator = positionDescPositionId.Validators[0].(func(int) error)
	propertyFields := schema.Property{}.Fields()
	_ = propertyFields
	// propertyDescCreatedAt is the schema descriptor for createdAt field.
	propertyDescCreatedAt := propertyFields[0].Descriptor()
	// property.DefaultCreatedAt holds the default value on creation for the createdAt field.
	property.DefaultCreatedAt = propertyDescCreatedAt.Default.(func() time.Time)
	// propertyDescName is the schema descriptor for name field.
	propertyDescName := propertyFields[1].Descriptor()
	// property.NameValidator is a validator for the "name" field. It is called by the builders before save.
	property.NameValidator = propertyDescName.Validators[0].(func(string) error)
	// propertyDescValue is the schema descriptor for value field.
	propertyDescValue := propertyFields[2].Descriptor()
	// property.ValueValidator is a validator for the "value" field. It is called by the builders before save.
	property.ValueValidator = propertyDescValue.Validators[0].(func(string) error)
	// propertyDescType is the schema descriptor for type field.
	propertyDescType := propertyFields[3].Descriptor()
	// property.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	property.TypeValidator = propertyDescType.Validators[0].(func(string) error)
	sectionFields := schema.Section{}.Fields()
	_ = sectionFields
	// sectionDescCreatedAt is the schema descriptor for createdAt field.
	sectionDescCreatedAt := sectionFields[0].Descriptor()
	// section.DefaultCreatedAt holds the default value on creation for the createdAt field.
	section.DefaultCreatedAt = sectionDescCreatedAt.Default.(func() time.Time)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescCreatedAt is the schema descriptor for createdAt field.
	tagDescCreatedAt := tagFields[0].Descriptor()
	// tag.DefaultCreatedAt holds the default value on creation for the createdAt field.
	tag.DefaultCreatedAt = tagDescCreatedAt.Default.(func() time.Time)
	// tagDescName is the schema descriptor for name field.
	tagDescName := tagFields[1].Descriptor()
	// tag.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tag.NameValidator = tagDescName.Validators[0].(func(string) error)
}

const (
	Version = "v0.11.7"                                         // Version of ent codegen.
	Sum     = "h1:V+wKFh0jhAbY/FoU+PPbdMOf2Ma5vh07R/IdF+N/nFg=" // Sum of ent codegen.
)
