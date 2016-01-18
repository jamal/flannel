package flannel

import (
	"net/http"

	"github.com/jamal/pint"
)

var formDecoder *pint.Decoder

func init() {
	formDecoder = pint.NewDecoder()
	formDecoder.Tag = "form"
}

// DecodeForm decodes the post form into v.
func DecodeForm(r *http.Request, v interface{}) error {
	return formDecoder.Decode(r, v)
}
