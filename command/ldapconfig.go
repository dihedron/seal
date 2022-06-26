package command

import (
	"encoding/json"
	"net/url"

	"github.com/dihedron/rawdata"
)

type LDAPConfiguration struct {
	Endpoint URL    `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	BaseDN   string `json:"basedn,omitempty" yaml:"basedn,omitempty"`
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
	Insecure bool   `json:"insecure,omitempty" yaml:"insecure,omitempty"`
}

func (c *LDAPConfiguration) UnmarshalFlag(value string) error {
	return rawdata.UnmarshalInto(value, c)
}

type URL struct {
	*url.URL
}

func (u *URL) UnmarshalJSON(data []byte) error {
	s := ""
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	v, err := url.Parse(s)
	if err != nil {
		return err
	}
	u.URL = v
	return nil
}

// func (u *URL) UnmarshalFlag(value string) error {
// 	v, err := url.Parse(value)
// 	if err != nil {
// 		return err
// 	}
// 	u.URL = v
// 	return nil
// }
