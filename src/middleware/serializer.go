package middleware

import (
	"bytes"
	"encoding/gob"
	"webdawgengine/config"
	"webdawgengine/crypt"
	"webdawgengine/models"
)

// A cookie serializer is a better way to handle session data. they are still
// generated, validated, and read only by the server, but they are stored on the
// client in a cookie.

// For example, when a user logs into a web service, all of their auth data is
// packed into a serialized encrypted string, which is sent via a cookie. this
// cookie can be sent back to the page, decrypted, and de-serialized to retrieve
// auth information in code. this is extremely fast and cheap, since you do not
// need to store this data in a database, or even in memory.

// Of course with this approach you must be careful not to leak the encryption
// key, since it can be used to decrypt legitimate keys, and sign faulty ones.
// The key should not be checked into VCS, and be regenerated if theft is
// suspected.

func EncryptSession(data map[string]interface{}) (string, error) {
	// serialize
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(data)
	if err != nil {
		return "", err
	}

	// encrypt
	sessionString, err := crypt.EncryptSecret(b.Bytes(), config.Config.IdentityPrivateKey)
	if err != nil {
		return "", err
	}

	return sessionString, nil
}

func DecryptSession(sessionString string) (map[string]interface{}, error) {
	session := make(map[string]interface{})

	secret, err := crypt.DecryptSecret(sessionString, config.Config.IdentityPrivateKey)
	if err != nil {
		return nil, err
	}

	// de-serialized
	b := bytes.Buffer{}
	b.Write(secret)

	d := gob.NewDecoder(&b)
	gobErr := d.Decode(&session)
	if gobErr != nil {
		return nil, gobErr
	}

	return session, nil
}

func EncryptIdentity(data *models.Identity) (string, error) {
	// serialize
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(data)
	if err != nil {
		return "", err
	}

	// encrypt
	auth, err := crypt.EncryptSecret(b.Bytes(), config.Config.IdentityPrivateKey)
	if err != nil {
		return "", err
	}

	return auth, nil
}

func DecryptIdentity(authString string) (*models.Identity, error) {
	cookie := models.Identity{}

	// decrypt
	secret, err := crypt.DecryptSecret(authString, config.Config.IdentityPrivateKey)
	if err != nil {
		return nil, err
	}

	// de-serialized
	b := bytes.Buffer{}
	b.Write(secret)

	d := gob.NewDecoder(&b)
	gobErr := d.Decode(&cookie)
	if gobErr != nil {
		return nil, gobErr
	}

	return &cookie, nil
}
