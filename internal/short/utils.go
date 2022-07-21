package short

import (
	"crypto/rand"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// FormatURL trims URL trailing slash, set schema and host to lowercase
func FormatURL(URL string) (string, error) {
	if !strings.HasPrefix(URL, "http://") && !strings.HasPrefix(URL, "https://") {
		URL = fmt.Sprintf("http://%s", URL)
	}

	URL = strings.TrimSuffix(URL, "/")
	parsedURL, err := url.Parse(URL)

	if err != nil {
		return "", err
	}

	schema := strings.ToLower(parsedURL.Scheme)
	host := strings.ToLower(parsedURL.Host)
	path := parsedURL.Path
	rquery := parsedURL.RawQuery

	if rquery != "" {
		rquery = fmt.Sprintf("?%s", rquery)
	}

	URL = fmt.Sprintf("%s://%s%s%s", schema, host, path, rquery)

	return URL, nil
}

// RandomCode generates a random code
func RandomCode(length int, chars []byte) (string, error) {
	if length == 0 {
		return "", errors.New("uniuri: length must be a positive integer")
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		return "", errors.New("uniuri: wrong charset length for NewLenChars")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4))
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			return "", errors.New("uniuri: error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b), nil
			}
		}
	}
}
