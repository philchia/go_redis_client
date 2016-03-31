package redis

import "errors"

// Result interface
type Result interface {
	String() (string, error)
	StringArray() ([]string, error)
	StringMap() (map[string]string, error)
}

type redisResult struct {
	Res interface{}
}

func (rr *redisResult) String() (string, error) {
	switch rr.Res.(type) {
	case error:
		return "", rr.Res.(error)
	case string:
		return rr.Res.(string), nil
	case []byte:
		return string(rr.Res.([]byte)), nil
	}
	return "", errors.New("Result is not string format")
}

func (rr *redisResult) StringArray() ([]string, error) {
	switch rr.Res.(type) {
	case error:
		return nil, rr.Res.(error)
	case []string:
		return rr.Res.([]string), nil

	}
	return nil, errors.New("Result is not string array format")
}

func (rr *redisResult) StringMap() (map[string]string, error) {
	switch rr.Res.(type) {
	case error:
		return nil, rr.Res.(error)
	case map[string]string:
		return rr.Res.(map[string]string), nil
	}
	return nil, errors.New("Result is not string map format")
}
