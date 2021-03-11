package config

type Config struct {
	ApiUrl      string
	SaveSecrets bool
}

type Secrets struct {
	Secrets map[string]string
}
