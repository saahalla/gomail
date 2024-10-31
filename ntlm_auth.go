package mail

import (
	"bytes"
	"errors"
	"log"

	ntlm "github.com/bigkraig/go-ntlm/ntlm"
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
		session, err := ntlm.CreateClientSession(ntlm.Version1, ntlm.ConnectionlessMode)
		if err != nil {
			return []byte{}, errors.New("error create ntlm session")
		}
		session.SetUserInfo(a.username, a.password, a.host)
		negotiate, err := session.GenerateNegotiateMessage()
		if err != nil {
			return []byte{}, errors.New("error generate negotiate message ntlm")
		}

		return negotiate.Bytes, nil
	default:
		session, err := ntlm.CreateClientSession(ntlm.Version1, ntlm.ConnectionlessMode)
		if err != nil {
			return []byte{}, errors.New("error create ntlm session")
		}
		session.SetUserInfo(a.username, a.password, a.host)

		challengeMessage, err := ntlm.ParseChallengeMessage(fromServer)
		if err != nil {
			return []byte{}, errors.New("error parse challenge message ntlm")
		}
		err = session.ProcessChallengeMessage(challengeMessage)
		if err != nil {
			return []byte{}, errors.New("error process challenge message ntlm")
		}
		authenticationMsssage, err := session.GenerateAuthenticateMessage()
		if err != nil {
			return []byte{}, errors.New("error generate authentication message ntlm")
		}
		log.Println(authenticationMsssage.String())
		return authenticationMsssage.Bytes(), nil
	}
}
