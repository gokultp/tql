package fs

import (
	"io"
	"os"
)

func GetFileReader(fileName *string) (io.Reader, error) {
	if fileName != nil {
		return os.Open(*fileName)
	}
	stat, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return os.Stdin, nil
	}
	return nil, nil
}
