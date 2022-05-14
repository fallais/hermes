package models

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Thing is something that has to be done.
type Thing struct {
	Name string `mapstructure:"name"`
	When string `mapstructure:"when"`
}
