package zif

import (
	"encoding/json"
	"errors"

	"golang.org/x/crypto/ed25519"
)

// This contains the more "complex" structures that will be sent in message
// data fields.

type MessageHashList struct {
	HashList  []byte
	Signature []byte
}

func (mhl *MessageHashList) Verify(pk ed25519.PublicKey) error {
	// TODO: Check length of pk/hashlist/etc

	verified := ed25519.Verify(pk, mhl.HashList, mhl.Signature)

	if !verified {
		return errors.New("Invalid signature")
	}

	return nil
}

func (mhl *MessageHashList) Encode() ([]byte, error) {
	data, err := json.Marshal(mhl)
	return data, err
}

func MessageHashListDecode(data []byte) (*MessageHashList, error) {
	ret := MessageHashList{}

	err := json.Unmarshal(data, &ret)

	if err != nil {
		return nil, err
	}

	return &ret, nil
}