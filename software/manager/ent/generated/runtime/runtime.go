// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/google/uuid"
	"github.com/niwla23/lagersystem/manager/ent/generated/box"
	"github.com/niwla23/lagersystem/manager/ent/generated/part"
	"github.com/niwla23/lagersystem/manager/ent/generated/position"
	"github.com/niwla23/lagersystem/manager/ent/generated/property"
	"github.com/niwla23/lagersystem/manager/ent/generated/tag"
	"github.com/niwla23/lagersystem/manager/ent/generated/warehouse"
	"github.com/niwla23/lagersystem/manager/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	boxHooks := schema.Box{}.Hooks()
	box.Hooks[0] = boxHooks[0]
	boxFields := schema.Box{}.Fields()
	_ = boxFields
	// boxDescCreatedAt is the schema descriptor for createdAt field.
	boxDescCreatedAt := boxFields[1].Descriptor()
	// box.DefaultCreatedAt holds the default value on creation for the createdAt field.
	box.DefaultCreatedAt = boxDescCreatedAt.Default.(func() time.Time)
	// boxDescUpdatedAt is the schema descriptor for updatedAt field.
	boxDescUpdatedAt := boxFields[2].Descriptor()
	// box.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	box.DefaultUpdatedAt = boxDescUpdatedAt.Default.(func() time.Time)
	// boxDescID is the schema descriptor for id field.
	boxDescID := boxFields[0].Descriptor()
	// box.DefaultID holds the default value on creation for the id field.
	box.DefaultID = boxDescID.Default.(func() uuid.UUID)
	partHooks := schema.Part{}.Hooks()
	part.Hooks[0] = partHooks[0]
	partFields := schema.Part{}.Fields()
	_ = partFields
	// partDescCreatedAt is the schema descriptor for createdAt field.
	partDescCreatedAt := partFields[0].Descriptor()
	// part.DefaultCreatedAt holds the default value on creation for the createdAt field.
	part.DefaultCreatedAt = partDescCreatedAt.Default.(func() time.Time)
	// partDescUpdatedAt is the schema descriptor for updatedAt field.
	partDescUpdatedAt := partFields[1].Descriptor()
	// part.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	part.DefaultUpdatedAt = partDescUpdatedAt.Default.(func() time.Time)
	// partDescDeleted is the schema descriptor for deleted field.
	partDescDeleted := partFields[2].Descriptor()
	// part.DefaultDeleted holds the default value on creation for the deleted field.
	part.DefaultDeleted = partDescDeleted.Default.(bool)
	// partDescName is the schema descriptor for name field.
	partDescName := partFields[3].Descriptor()
	// part.NameValidator is a validator for the "name" field. It is called by the builders before save.
	part.NameValidator = partDescName.Validators[0].(func(string) error)
	// partDescAmount is the schema descriptor for amount field.
	partDescAmount := partFields[5].Descriptor()
	// part.DefaultAmount holds the default value on creation for the amount field.
	part.DefaultAmount = partDescAmount.Default.(int)
	positionHooks := schema.Position{}.Hooks()
	position.Hooks[0] = positionHooks[0]
	positionFields := schema.Position{}.Fields()
	_ = positionFields
	// positionDescCreatedAt is the schema descriptor for createdAt field.
	positionDescCreatedAt := positionFields[0].Descriptor()
	// position.DefaultCreatedAt holds the default value on creation for the createdAt field.
	position.DefaultCreatedAt = positionDescCreatedAt.Default.(func() time.Time)
	// positionDescUpdatedAt is the schema descriptor for updatedAt field.
	positionDescUpdatedAt := positionFields[1].Descriptor()
	// position.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	position.DefaultUpdatedAt = positionDescUpdatedAt.Default.(func() time.Time)
	propertyHooks := schema.Property{}.Hooks()
	property.Hooks[0] = propertyHooks[0]
	property.Hooks[1] = propertyHooks[1]
	propertyFields := schema.Property{}.Fields()
	_ = propertyFields
	// propertyDescCreatedAt is the schema descriptor for createdAt field.
	propertyDescCreatedAt := propertyFields[0].Descriptor()
	// property.DefaultCreatedAt holds the default value on creation for the createdAt field.
	property.DefaultCreatedAt = propertyDescCreatedAt.Default.(func() time.Time)
	// propertyDescUpdatedAt is the schema descriptor for updatedAt field.
	propertyDescUpdatedAt := propertyFields[1].Descriptor()
	// property.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	property.DefaultUpdatedAt = propertyDescUpdatedAt.Default.(func() time.Time)
	// propertyDescName is the schema descriptor for name field.
	propertyDescName := propertyFields[2].Descriptor()
	// property.NameValidator is a validator for the "name" field. It is called by the builders before save.
	property.NameValidator = propertyDescName.Validators[0].(func(string) error)
	// propertyDescValue is the schema descriptor for value field.
	propertyDescValue := propertyFields[3].Descriptor()
	// property.ValueValidator is a validator for the "value" field. It is called by the builders before save.
	property.ValueValidator = propertyDescValue.Validators[0].(func(string) error)
	// propertyDescType is the schema descriptor for type field.
	propertyDescType := propertyFields[4].Descriptor()
	// property.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	property.TypeValidator = propertyDescType.Validators[0].(func(string) error)
	tagHooks := schema.Tag{}.Hooks()
	tag.Hooks[0] = tagHooks[0]
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescCreatedAt is the schema descriptor for createdAt field.
	tagDescCreatedAt := tagFields[0].Descriptor()
	// tag.DefaultCreatedAt holds the default value on creation for the createdAt field.
	tag.DefaultCreatedAt = tagDescCreatedAt.Default.(func() time.Time)
	// tagDescUpdatedAt is the schema descriptor for updatedAt field.
	tagDescUpdatedAt := tagFields[1].Descriptor()
	// tag.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	tag.DefaultUpdatedAt = tagDescUpdatedAt.Default.(func() time.Time)
	// tagDescName is the schema descriptor for name field.
	tagDescName := tagFields[2].Descriptor()
	// tag.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tag.NameValidator = tagDescName.Validators[0].(func(string) error)
	warehouseHooks := schema.Warehouse{}.Hooks()
	warehouse.Hooks[0] = warehouseHooks[0]
	warehouseFields := schema.Warehouse{}.Fields()
	_ = warehouseFields
	// warehouseDescCreatedAt is the schema descriptor for createdAt field.
	warehouseDescCreatedAt := warehouseFields[0].Descriptor()
	// warehouse.DefaultCreatedAt holds the default value on creation for the createdAt field.
	warehouse.DefaultCreatedAt = warehouseDescCreatedAt.Default.(func() time.Time)
	// warehouseDescUpdatedAt is the schema descriptor for updatedAt field.
	warehouseDescUpdatedAt := warehouseFields[1].Descriptor()
	// warehouse.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	warehouse.DefaultUpdatedAt = warehouseDescUpdatedAt.Default.(func() time.Time)
	// warehouseDescName is the schema descriptor for name field.
	warehouseDescName := warehouseFields[2].Descriptor()
	// warehouse.NameValidator is a validator for the "name" field. It is called by the builders before save.
	warehouse.NameValidator = warehouseDescName.Validators[0].(func(string) error)
}

const (
	Version = "v0.11.7"                                         // Version of ent codegen.
	Sum     = "h1:V+wKFh0jhAbY/FoU+PPbdMOf2Ma5vh07R/IdF+N/nFg=" // Sum of ent codegen.
)
