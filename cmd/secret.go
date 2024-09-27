package cmd

import (
	"encoding/json"
	"fmt"
)

func parseSecretAsJSON(contents string) (map[string]string, error) {
	kv := map[string]string{}
	if err := json.Unmarshal([]byte(contents), &kv); err != nil {
		return nil, err
	}
	return kv, nil
}

func displayAsEnv(kv map[string]string) {
	for k, v := range kv {
		fmt.Printf("export %s='%s'\n", k, v)
	}

}
