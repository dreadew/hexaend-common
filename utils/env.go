package utils

import (
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func EnvError(name string) error {
	return errors.Errorf("failed to get %s env", name)
}

func GetEnvAsInt(env string) (int, error) {
	valStr := os.Getenv(env)
	if len(valStr) == 0 {
		return 0, errors.Errorf("cannot find env variable with the name %s", env)
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		return 0, errors.Errorf("invalid value for %s: %v", env, err)
	}

	return val, nil
}

func GetEnvAsFloat(env string, bitSize int) (float64, error) {
	valStr := os.Getenv(env)
	if len(valStr) == 0 {
		return 0, errors.Errorf("cannot find env variable with the name %s", env)
	}

	val, err := strconv.ParseFloat(valStr, bitSize)
	if err != nil {
		return 0, errors.Errorf("invalid value for %s: %v", env, err)
	}

	return val, nil
}

func GetEnvAsStringArray(env string) ([]string, error) {
	val := os.Getenv(env)
	if len(val) == 0 {
		return nil, errors.Errorf("cannot find env variable with the name %s", env)
	}

	return strings.Split(val, ","), nil
}
