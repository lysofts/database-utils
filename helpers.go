package databaseutils

import (
	"log"
	"os"
)

func GetEnv(varName string) string {
	val := os.Getenv(varName)
	if len(val) == 0 {
		log.Panicf("environment variable %v not found", varName)
	}
	return val
}
