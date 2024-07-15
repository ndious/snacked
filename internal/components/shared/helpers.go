package components

import (
	"fmt"

	"github.com/a-h/templ"
)

func HxURL(url string, params ...any) string {
    return string(URL(url, params...))
}

func URL(url string, params ...any) templ.SafeURL {
    return templ.URL(fmt.Sprintf(url, params...))
}
