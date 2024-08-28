package upgrade

import (
	"encoding/json"
	"errors"
	"net/http"
)

func getAndDecode[T any](u *Upgrader, getUrl string, result *T) error {
	req, err := http.NewRequest("GET", getUrl, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("PRIVATE-TOKEN", u.Opts.Token)
	r, err := u.cl.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return errors.New(r.Status)
	}

	err = json.NewDecoder(r.Body).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func head(u *Upgrader, getUrl string) (*http.Response, error) {
	req, err := http.NewRequest("HEAD", getUrl, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("PRIVATE-TOKEN", u.Opts.Token)
	r, err := u.cl.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return r, nil
}
