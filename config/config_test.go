package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_readJson(t *testing.T) {
	conf, err := readJson("./data/config.json")
	assert.Nil(t, err)
	assert.NotNil(t, conf)

	assert.Equal(t, "10.0.0.1", conf.Database.Address)
	assert.Equal(t, 3306, conf.Database.Port)
}

func Test_writeJson(t *testing.T) {
	err := writeJson("./data/test.json", []byte(`
	{
		"database": {
			"address": "10.0.0.1",
			"port": 3306
		}
	}
	`))
	assert.Nil(t, err)
}
