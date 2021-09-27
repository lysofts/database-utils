package utils

import (
	"log"
	"os"

	"github.com/mitchellh/mapstructure"
)

func GetEnv(varName string) string {
	val := os.Getenv(varName)
	if len(val) == 0 {
		log.Panicf("environment variable %v not found", varName)
	}
	return val
}

//JsonToBson updates the id field into _id for bson
func JsonToBson(payload interface{}) (Map, error) {

	newPayload := Map{}
	err := mapstructure.Decode(payload, &newPayload)

	if err != nil {
		return nil, err
	}

	value, found := newPayload["id"]
	if found {
		newPayload["_id"] = value
		delete(newPayload, "id")
	}

	return newPayload, nil
}

//JsonToBson updates the _id field into id for json
func BsonToJson(payload interface{}) (Map, error) {

	newPayload := Map{}
	err := mapstructure.Decode(payload, &newPayload)

	if err != nil {
		return nil, err
	}

	value, found := newPayload["_id"]
	if found {
		newPayload["id"] = value
		delete(newPayload, "_id")
	}

	return newPayload, nil
}
