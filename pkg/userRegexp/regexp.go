package userRegexp

import "regexp"

var (
	BoldWithUderscore, _     = regexp.Compile(`__[a-zA-z0-9]*__`)       // __abc__
	BoldWithStar, _          = regexp.Compile(`\*\*[a-zA-Z0-9]*\*\*`)   // **abc**
	ItalicsWithStar, _       = regexp.Compile(`\*[a-zA-Z0-9]*\*`)       // *abc*
	ItalicsWithUnderscore, _ = regexp.Compile(`_[a-zA-Z0-9]*_`)         // _abc_
	BoldAndItalics, _        = regexp.Compile(`\*\*\*[a-zA-Z0-9]*\*\*`) // ***abc***
)
