package main

import (
	"os"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
)

type Conf struct {
	Database struct {
		Address string
		Port    int
	}
}

func readJson(filePath string) (*Conf, error) {
	if err := config.Load(
		file.NewSource(
			file.WithPath(filePath),
		),
	); err != nil {
		return nil, err
	}
	var conf Conf
	if err := config.Get().Scan(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

func writeJson(filePath string, data []byte) error {
	fh, err := os.Create(filePath)
	if err != nil {
		return nil
	}
	defer func() {
		fh.Close()
	}()

	_, err = fh.Write(data)
	if err != nil {
		return err
	}

	return nil
}
