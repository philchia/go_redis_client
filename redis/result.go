package redis

import "errors"

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
	case Result:
		return rr.Res.(Result).String()
	}
	return "", errors.New("Result is not string format")
}

func (rr *redisResult) Int() (int, error) {
	switch rr.Res.(type) {
	case error:
		return -1, rr.Res.(error)
	case int:
		return rr.Res.(int), nil
	}
	return -1, errors.New("Result is not int format")
}

func (rr *redisResult) Int32() (int32, error) {
	switch rr.Res.(type) {
	case error:
		return -1, rr.Res.(error)
	case int:
		return int32(rr.Res.(int)), nil
	case int64:
		return int32(rr.Res.(int64)), nil
	case int32:
		return rr.Res.(int32), nil
	}
	return -1, errors.New("Result is not int32 format")
}

func (rr *redisResult) Int64() (int64, error) {
	switch rr.Res.(type) {
	case error:
		return -1, rr.Res.(error)
	case int:
		return int64(rr.Res.(int)), nil
	case int64:
		return rr.Res.(int64), nil
	case int32:
		return int64(rr.Res.(int32)), nil
	}
	return -1, errors.New("Result is not int64 format")
}

func (rr *redisResult) StringArray() ([]string, error) {
	switch rr.Res.(type) {
	case error:
		return nil, rr.Res.(error)
	case []string:
		return rr.Res.([]string), nil
	case []Result:
		var arr []string
		results := rr.Res.([]Result)
		for i, r := range results {
			str, err := r.String()
			if err != nil {
				return nil, err
			}
			arr = append(arr, str)
		}
		return arr, nil
	}
	return nil, errors.New("Result is not string array format")
}

func (rr *redisResult) StringMap() (map[string]string, error) {
	switch rr.Res.(type) {
	case error:
		return nil, rr.Res.(error)
	case map[string]string:
		return rr.Res.(map[string]string), nil
	case []Result:
		results := rr.Res.([]Result)

		length := len(results)
		if length%2 != 0 {
			return nil, errors.New("Result is not a string map format")
		}
		m := make(map[string]string)
		for i := 0; i < length; i++ {
			k, err := results[i].String()
			if err != nil {
				return nil, err
			}
			i++
			v, err := results[i].String()
			if err != nil {
				return nil, err
			}
			m[k] = v
		}
		return m, nil
	}
	return nil, errors.New("Result is not string map format")
}

func (rr *redisResult) Array() ([]Result, error) {
	switch rr.Res.(type) {
	case []Result:
		return rr.Res.([]Result), nil
	}
	return nil, errors.New("Result is not Array of result format")
}

func (rr *redisResult) Bool() (bool, error) {
	switch rr.Res.(type) {
	case bool:
		return rr.Res.(bool), nil
	case int:
		return rr.Res.(int) != 0, nil
	}
	return false, errors.New("Result is not Array of result format")
}
