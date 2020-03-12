package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
)

type Secrets struct {
	DatabaseSecrets
}

type Configuration struct {
	ServerConfiguration
	DatabaseConfiguration
	Secrets
}

func (c *Configuration) GetConfiguration(v viper.Viper) {
	c.DatabaseConfiguration = GetDatabaseConfig(v)
	c.ServerConfiguration = GetServerConfig(v)
}

func (c *Configuration) GetSecrets() error {
	var secretsMap Secrets

	secrets, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", c.SrvSecretsPath, c.SrvSecretsFile))
	if err != nil { return err }

	if err := json.Unmarshal(secrets, &secretsMap); err != nil {
		return err
	}

	c.Secrets = secretsMap
	return nil
}
