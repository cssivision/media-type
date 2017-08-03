package mediaType

import (
	"errors"
	"regexp"
	"strings"
)

/**
* RegExp to match type in RFC 6838
*
* type-name = restricted-name
* subtype-name = restricted-name
* restricted-name = restricted-name-first *126restricted-name-chars
* restricted-name-first  = ALPHA / DIGIT
* restricted-name-chars  = ALPHA / DIGIT / "!" / "#" /
*                          "$" / "&" / "-" / "^" / "_"
* restricted-name-chars =/ "." ; Characters before first dot always
*                              ; specify a facet name
* restricted-name-chars =/ "+" ; Characters after last plus always
*                              ; specify a structured syntax suffix
* ALPHA =  %x41-5A / %x61-7A   ; A-Z / a-z
* DIGIT =  %x30-39             ; 0-9
 */

var (
	subtypeNameRegexp = regexp.MustCompile(`[A-Za-z0-9][A-Za-z0-9!#$&^_.-]{0,126}$`)
	typeNameRegexp    = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9!#$&^_-]{0,126}$`)
	nameRegexp        = regexp.MustCompile(`^ *([A-Za-z0-9][A-Za-z0-9!#$&^_-]{0,126})\/([A-Za-z0-9][A-Za-z0-9!#$&^_.+-]{0,126}) *$`)
)

// MediaType mediaType object.
type MediaType struct {
	Type       string
	SubType    string
	Suffix     string
	Parameters map[string]string
}

// Format format object to media type.
func (mt *MediaType) Format() (string, error) {
	if mt.Type == "" || !typeNameRegexp.MatchString(mt.Type) {
		return "", errors.New("invalid type")
	}

	if mt.SubType == "" || !subtypeNameRegexp.MatchString(mt.SubType) {
		return "", errors.New("invalid subtype")
	}

	str := mt.Type + "/" + mt.SubType
	if mt.Suffix != "" {
		if !typeNameRegexp.MatchString(mt.Suffix) {
			return "", errors.New("invliad suffix")
		}
		str += "+" + mt.Suffix
	}

	return str, nil
}

// Parse parse media type to object.
func Parse(str string) (*MediaType, error) {
	if str == "" {
		return nil, errors.New("invliad parse string")
	}

	match := nameRegexp.FindStringSubmatch(str)
	if len(match) == 0 || len(match) < 3 {
		return nil, errors.New("invlaid media type")
	}
	mt := &MediaType{}
	mt.Type = match[1]
	mt.SubType = match[2]
	index := strings.LastIndex(mt.SubType, "+")
	if index != -1 {
		mt.Suffix = mt.SubType[index+1:]
		mt.SubType = mt.SubType[0:index]
	}
	return mt, nil
}
