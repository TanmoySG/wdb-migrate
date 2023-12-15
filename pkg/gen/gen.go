package gen

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

const wdbRetroIdKey = "_id"
const jsonSchemaRequiredFieldKey = "required"

func sampleData(data map[string]interface{}) ([]byte, error) {
	samplerData := map[string]interface{}{}

	for _, rb := range data {
		// picks first element from map
		samplerData = rb.(map[string]interface{})
		break
	}

	if len(samplerData) == 0 {
		return nil, fmt.Errorf("no data found for sampling")
	}

	delete(samplerData, wdbRetroIdKey)

	samplerDataBytes, err := json.Marshal(samplerData)
	if err != nil {
		return nil, err
	}

	return samplerDataBytes, nil
}

func GenerateJsonSchema(data map[string]interface{}) (map[string]interface{}, error) {
	sampledData, err := sampleData(data)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("echo", string(sampledData))
	catOutput, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	gensonCmd := exec.Command("genson")
	gensonCmd.Stdin = strings.NewReader(string(catOutput))

	generatedJSONSchemaBytes, err := gensonCmd.Output()
	if err != nil {
		return nil, err
	}

	var generatedJSONSchema map[string]interface{}

	err = json.Unmarshal(generatedJSONSchemaBytes, &generatedJSONSchema)
	if err != nil {
		return nil, err
	}

	// for safe migration remove `required` field for json schema
	// migrated schema should be flexible as incoming data is from v0
	// and might be incompatible
	delete(generatedJSONSchema, jsonSchemaRequiredFieldKey)

	return generatedJSONSchema, nil
}
