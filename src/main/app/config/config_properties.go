package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/go-chassis/go-archaius"

	"github.com/src/main/app/helpers/files"

	"github.com/src/main/app/config/env"
)

const (
	File = "config.yml"
)

func init() {
	log.Println("INFO: Trying to load config ...")
	_, caller, _, _ := runtime.Caller(0)
	root := path.Join(path.Dir(caller), "../../..")
	err := os.Chdir(root)
	if err != nil {
		log.Println(root)
		log.Fatalln(err)
	}

	propertiesPath, environment, scope :=
		fmt.Sprintf("%s/resources/config", root),
		env.GetEnv(),
		env.GetScope()

	var compositeConfig []string

	scopeConfig := fmt.Sprintf("%s/%s/%s.%s", propertiesPath, environment, scope, File)
	if files.Exist(scopeConfig) {
		compositeConfig = append(compositeConfig, scopeConfig)
	}

	envConfig := fmt.Sprintf("%s/%s/%s", propertiesPath, environment, File)
	if files.Exist(envConfig) {
		compositeConfig = append(compositeConfig, envConfig)
	}

	sharedConfig := fmt.Sprintf("%s/%s", propertiesPath, File)
	if files.Exist(sharedConfig) {
		compositeConfig = append(compositeConfig, sharedConfig)
	}

	err = archaius.Init(
		archaius.WithENVSource(),
		archaius.WithRequiredFiles(compositeConfig),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("INFO: ENV: %s, SCOPE: %s", environment, scope)
}

func String(key string) string {
	return archaius.GetString(key, "")
}

func Int(key string) int {
	return archaius.GetInt(key, 0)
}
