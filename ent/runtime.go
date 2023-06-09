// Code generated by ent, DO NOT EDIT.

package ent

import (
	"webapp/ent/schema"
	"webapp/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[2].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[6].Descriptor()
	// user.DefaultAvatar holds the default value on creation for the avatar field.
	user.DefaultAvatar = userDescAvatar.Default.(string)
	// userDescCareer is the schema descriptor for career field.
	userDescCareer := userFields[7].Descriptor()
	// user.DefaultCareer holds the default value on creation for the career field.
	user.DefaultCareer = userDescCareer.Default.(string)
}
