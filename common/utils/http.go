package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var policy = bluemonday.StrictPolicy()

func ParseQuery(query string) (url.Values, error) {
	var err error
	m := make(url.Values)

	for query != "" {
		key := query
		if i := strings.IndexAny(key, "&"); i >= 0 {
			key, query = key[:i], key[i+1:]
		} else {
			query = ""
		}
		if key == "" {
			continue
		}
		value := ""
		if i := strings.Index(key, "="); i >= 0 {
			key, value = key[:i], key[i+1:]
		}
		key, err1 := url.QueryUnescape(key)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		value, err1 = url.QueryUnescape(value)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		value = policy.Sanitize(value)
		m[key] = append(m[key], value)
	}
	return m, err
}

// ParseResponse function to parse json response body to a given model
func ParseResponse(r *http.Response, model interface{}) error {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("cannot read response: error=%s, resp=%+v", err, r.Body)
		return errors.New(msg)
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	defer r.Body.Close()

	err = json.Unmarshal(bodyBytes, &model)
	if err != nil {
		msg := fmt.Sprintf("cannot parse json response: error=%s, resp=%s", err, string(bodyBytes))
		return errors.New(msg)
	}

	return nil
}

func ParseResponseBody(bodyBytes []byte, model interface{}) error {
	if err := json.Unmarshal(bodyBytes, &model); err != nil {
		msg := fmt.Sprintf("cannot parse json response: error=%s, resp=%s", err, string(bodyBytes))
		return errors.New(msg)
	}

	return nil
}
