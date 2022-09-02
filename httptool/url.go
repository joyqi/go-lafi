package httptool

import (
	"github.com/google/go-querystring/query"
	"net/url"
)

// MakeURL builds a URL for the specified url string and params
func MakeURL(uri string, params url.Values) string {
	u, err := url.Parse(uri)
	if err != nil {
		return uri
	}

	v := u.Query()
	for key, val := range params {
		if len(val) > 0 {
			v.Set(key, val[0])
		}
	}

	u.RawQuery = v.Encode()
	return u.String()
}

func MakeStructureURL(uri string, params interface{}) string {
	u, err := url.Parse(uri)
	if err != nil {
		return uri
	}

	v, err := query.Values(params)
	if err != nil {
		return uri
	}

	u.RawQuery = v.Encode()
	return u.String()
}
