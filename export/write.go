package export

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func WriteFile(filename string, buffer io.Reader) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		logrus.WithField("type", "writefile").WithError(err).Errorf("open file %s failed", filename)
		return err
	}
	defer func() { _ = file.Close() }()

	if _, err = io.Copy(file, buffer); err != nil {
		logrus.WithField("type", "writefile").WithError(err).Errorf("io.Copy failed, filename: %s", filename)
		return err
	}
	return nil
}
