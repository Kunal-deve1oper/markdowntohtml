package userRegexp

import "regexp"

var (
	BoldWithUderscore, _     = regexp.Compile(`__(.*?)__`)          // __abc__
	BoldWithStar, _          = regexp.Compile(`\*\*(.*?)\*\*`)      // **abc**
	ItalicsWithStar, _       = regexp.Compile(`\*(.*?)\*`)          // *abc*
	ItalicsWithUnderscore, _ = regexp.Compile(`_(.*?)_`)            // _abc_
	BoldAndItalics, _        = regexp.Compile(`\*\*\*(.*?)\*\*\*`)  // ***abc***
	Link, _                  = regexp.Compile(`\[(.*?)\]\((.*?)\)`) //[abc](http://abc.com)
)
