package api

import (
	"github.com/gorilla/schema"
)

// NewAPIDecoder returns a schema.Decoder to support Moby API types
func NewAPIDecoder() *schema.Decoder {
	d := schema.NewDecoder()
	d.IgnoreUnknownKeys(true)

	// TODO register custom converter
	/// d.RegisterConverter(...)
	return d
}
