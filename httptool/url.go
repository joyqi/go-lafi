package httptool

import (
	"github.com/google/go-querystring/query"
	"net/url"
	"regexp"
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

// MakeTemplateURL builds a URL by replacing the placeholders in the uri with the values in the params.
// e.g. MakeTemplateURL("https://example.com/:id", map[string]string{"id": "123"}) returns "https://example.com/123"
func MakeTemplateURL(uri string, params map[string]string) string {
	re := regexp.MustCompile(`:([_a-z\d-]+)`)
	return string(re.ReplaceAllFunc([]byte(uri), func(bytes []byte) []byte {
		key := string(bytes[1:])
		if val, ok := params[key]; ok {
			return []byte(val)
		} else {
			return bytes
		}
	}))
}

// MakeStructureURL builds a URL by using the structured params as the query string of the url.
// e.g. MakeStructureURL("https://example.com/", &testParams{State: "test", UserName: "hello"})
// returns "https://example.com/?state=test&user_name=hello"
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
