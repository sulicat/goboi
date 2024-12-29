package regex_utils

import "regexp"

// crediting this post:
// https://stackoverflow.com/questions/12643009/regular-expression-for-floating-point-numbers
var FloatingPointRE *regexp.Regexp = regexp.MustCompile("[+-]?([0-9]+([.][0-9]*)?|[.][0-9]+)")
