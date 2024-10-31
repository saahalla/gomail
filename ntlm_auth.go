package mail

import (
	"bytes"
	"errors"

	ntlm "github.com/Azure/go-ntlmssp"
)

// loginAuth is an smtp.Auth that implements the LOGIN authentication mechanism.
type ntlmAuth struct {
	username string
	password string
	host     string
}

func (a *ntlmAuth) Start(server *ServerInfo) (string, []byte, error) {
	if !server.TLS {
		advertised := false
		for _, mechanism := range server.Auth {
			if mechanism == "NTLM" {
				advertised = true
				break
			}
		}
		if !advertised {
			return "", nil, errors.New("gomail: unencrypted connection")
		}
	}
	if server.Name != a.host {
		return "", nil, errors.New("gomail: wrong host name")
	}
	return "NTLM", nil, nil
}

func (a *ntlmAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if !more {
		return nil, nil
	}

	switch {
	case bytes.Equal(fromServer, []byte("NTLM supported")):
		negotiate, err := ntlm.NewNegotiateMessage(a.host, "")
		if err != nil {
			return []byte{}, errors.New("error generate negotiate message ntlm")
		}

		return negotiate, nil
	default:
		challengeMsg, err := ntlm.ProcessChallenge(fromServer, a.username, a.password, true)
		if err != nil {
			return []byte{}, errors.New("error process challenge message ntlm")
		}

		return challengeMsg, nil
	}
}
