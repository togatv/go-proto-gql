// Code generated by github.com/vektah/gqlgen, DO NOT EDIT.

package gqlgen

type Description struct {
	Description string `json:"description"`
}
type Details interface{}
type Error struct {
	Code    *string   `json:"code"`
	Message *string   `json:"message"`
	Details []Details `json:"details"`
}
type Validation struct {
	Field       *string `json:"field"`
	Description *string `json:"description"`
}