package mediaType

var (
	subtypeNameRegexp = `[A-Za-z0-9][A-Za-z0-9!#$&^_.-]{0,126}$`
	typeNameRegexp    = `^[A-Za-z0-9][A-Za-z0-9!#$&^_-]{0,126}$`
	nameRegexp        = `^ *([A-Za-z0-9][A-Za-z0-9!#$&^_-]{0,126})\/([A-Za-z0-9][A-Za-z0-9!#$&^_.+-]{0,126}) *$`
)

// MediaType mediaType object.
type MediaType struct {
	Type    string
	SubType string
	Suffix  string
}

// Format format object to media type.
func Format(mt *MediaType) {

}

// Parse parse media type to object.
func Parse() {

}
