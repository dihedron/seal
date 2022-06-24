package command

import (
	"net/url"
)

type URL url.URL

func (u *URL) UnmarshalFlag(value string) error {
	v, err := url.Parse(value)
	if err != nil {
		return err
	}
	*u = URL(*v)
	return nil
}
