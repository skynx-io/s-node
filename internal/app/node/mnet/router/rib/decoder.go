package rib

import (
	"encoding/base64"
	"encoding/json"

	"skynx.io/s-api-go/grpc/network/routing"
	"skynx.io/s-lib/pkg/errors"
)

func (r *ribData) decoder(msg string) error {
	jsonData, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return errors.Wrapf(err, "[%v] function base64.StdEncoding.DecodeString()", errors.Trace())
	}

	d := &routing.RIBData{}

	if err := json.Unmarshal(jsonData, d); err != nil {
		return errors.Wrapf(err, "[%v] function json.Unmarshal()", errors.Trace())
	}

	r.rxQueue <- d

	return nil
}
