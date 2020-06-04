package git

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/client"
	githttp "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	gitssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

type Config struct {
	Username              string
	Password              string
	PrivateKeyFile        string
	InsecureIgnoreHostKey bool
	InsecureSkipTLSVerify bool
}

type Meta struct {
	Auth transport.AuthMethod
}

func (c *Config) Client() (interface{}, error) {
	var meta Meta

	if c.PrivateKeyFile != "" {
		auth, err := gitssh.NewPublicKeysFromFile(gitssh.DefaultUsername, c.PrivateKeyFile, "")
		if err != nil {
			return nil, fmt.Errorf("unable to get ssh key: %s", err)
		}

		if c.InsecureIgnoreHostKey {
			auth.HostKeyCallbackHelper = gitssh.HostKeyCallbackHelper{
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			}
		}

		meta.Auth = auth
	}

	if c.Username != "" && c.Password != "" {
		meta.Auth = &githttp.BasicAuth{
			Username: c.Username,
			Password: c.Password,
		}
	}

	httpsClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: c.InsecureSkipTLSVerify},
		},
	}

	client.InstallProtocol("https", githttp.NewClient(httpsClient))

	return &meta, nil
}
