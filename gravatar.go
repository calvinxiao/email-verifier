package emailverifier

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Gravatar is detail about the Gravatar
type Gravatar struct {
	HasGravatar bool   // whether has gravatar
	GravatarUrl string // gravatar url
}

// CheckGravatar will return the Gravatar records for the given email.
func (v *Verifier) CheckGravatar(email string) (Gravatar, error) {
	ret := Gravatar{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err, emailMd5 := getMD5Hash(strings.ToLower(strings.TrimSpace(email)))
	if err != nil {
		return ret, err
	}
	gravatarUrl := gravatarBaseUrl + emailMd5
	cli := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", gravatarUrl, nil)
	if err != nil {
		return ret, err
	}
	resp, err := cli.Do(req)
	if err != nil {
		return ret, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ret, err
	}
	// check body
	err, md5Body := getMD5Hash(string(body))
	if err != nil {
		return ret, err
	}
	if md5Body == gravatarDefaultMd5 || resp.StatusCode != 200 {
		return ret, nil
	}
	return Gravatar{
		HasGravatar: true,
		GravatarUrl: gravatarUrl,
	}, nil
}
