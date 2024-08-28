package upgrade

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

func getAndDecode[T any](u *Upgrader, getUrl string, result *T) error {
	req, err := http.NewRequest("GET", getUrl, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("PRIVATE-TOKEN", u.Opts.Token)
	cl := &http.Client{Timeout: time.Second * time.Duration(u.Opts.RequestTimeout)}
	r, err := cl.Do(req)
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
	cl := &http.Client{Timeout: time.Second * time.Duration(u.Opts.RequestTimeout)}
	r, err := cl.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return r, nil
}

func getDownload(u *Upgrader, getUrl string, output io.WriteCloser) error {
	req, err := http.NewRequest("GET", getUrl, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("PRIVATE-TOKEN", u.Opts.Token)
	cl := &http.Client{Timeout: time.Second * time.Duration(u.Opts.DownloadTimeout)}
	r, err := cl.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return errors.New(r.Status)
	}

	_, err = io.Copy(output, r.Body)
	return err
}
