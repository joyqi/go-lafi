package httptool

import "net/url"

// MakeURL builds a URL for the specified url string and params
func MakeURL(uri string, params url.Values) string {
	u, err := url.Parse(uri)
	if err != nil {
		return uri
	}

	query := u.Query()
	for key, val := range params {
		if len(val) > 0 {
			query.Set(key, val[0])
		}
	}

	u.RawQuery = query.Encode()
	return u.String()
}
