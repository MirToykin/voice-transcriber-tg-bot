package e

import "github.com/pkg/errors"

func WrapIfErr(msg string, err error) error {
	if err == nil {
		return nil
	}

	return errors.Wrap(err, msg)
}
