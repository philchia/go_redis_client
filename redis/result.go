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
	if str, ok := rr.Res.(string); ok {
		return str, nil
	}
	return "", errors.New("Result is not string type")
}

func (rr *redisResult) StringArray() ([]string, error) {
	if strs, ok := rr.Res.([]string); ok {
		return strs, nil
	}
	return []string{}, errors.New("Result is not array type")
}

func (rr *redisResult) StringMap() (map[string]string, error) {
	if mp, ok := rr.Res.(map[string]string); ok {
		return mp, nil
	}
	return map[string]string{}, errors.New("Result is not map type")
}
