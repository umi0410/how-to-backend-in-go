// Code generated by entc, DO NOT EDIT.

package ent

import (
	"db/ent/schema"
	"db/ent/tourproduct"
	"db/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	tourproductFields := schema.TourProduct{}.Fields()
	_ = tourproductFields
	// tourproductDescForSale is the schema descriptor for forSale field.
	tourproductDescForSale := tourproductFields[2].Descriptor()
	// tourproduct.DefaultForSale holds the default value on creation for the forSale field.
	tourproduct.DefaultForSale = tourproductDescForSale.Default.(bool)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescIsActivated is the schema descriptor for isActivated field.
	userDescIsActivated := userFields[2].Descriptor()
	// user.DefaultIsActivated holds the default value on creation for the isActivated field.
	user.DefaultIsActivated = userDescIsActivated.Default.(bool)
}