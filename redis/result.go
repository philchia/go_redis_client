package redis

import (
	"errors"
	"fmt"
	"strconv"
)

type redisResult struct {
	Value interface{}
	Err   error
}

func (rr redisResult) Error() error {
	if rr.Err != nil {
		return rr.Err
	}
	return nil
}

func (rr redisResult) OK() bool {
	return rr.Value == OK
}

func (rr redisResult) PONG() bool {
	return rr.Value == PONG
}

func (rr redisResult) String() (string, error) {
	switch val := rr.Value.(type) {
	case error:
		return "", rr.Value.(error)
	case string:
		return val, nil
	case []byte:
		return string(val), nil
	case int8:
		return strconv.FormatInt(int64(val), 10), nil
	case int16:
		return strconv.FormatInt(int64(val), 10), nil
	case int:
		return strconv.FormatInt(int64(val), 10), nil
	case int32:
		return strconv.FormatInt(int64(val), 10), nil
	case int64:
		return strconv.FormatInt(val, 10), nil
	case float64:
		return strconv.FormatFloat(val, 'g', -1, 64), nil
	case interface{}:
		return fmt.Sprint(rr.Value.(interface{})), nil
	}
	return "", errors.New("Result is not string format")
}

func (rr redisResult) Int() (int, error) {
	switch val := rr.Value.(type) {
	case error:
		return -1, val
	case int8:
		return int(val), nil
	case int16:
		return int(val), nil
	case int:
		return val, nil
	case int32:
		return int(val), nil
	case int64:
		return int(val), nil
	case string:
		return strconv.Atoi(val)
	case []byte:
		return strconv.Atoi(string(val))

	}
	return -1, errors.New("Result is not int format")
}

func (rr redisResult) Float64() (float64, error) {
	switch val := rr.Value.(type) {
	case error:
		return -1, val
	case int:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case string:
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return -1, err
		}
		return f, nil
	}
	return -1, errors.New("Result is not float64 format")
}

func (rr redisResult) Strings() ([]string, error) {
	switch rr.Value.(type) {
	case error:
		return nil, rr.Value.(error)
	case []string:
		return rr.Value.([]string), nil
	case []Result:
		var arr []string
		results := rr.Value.([]Result)
		for _, r := range results {
			str, err := r.String()
			if err != nil {
				return nil, err
			}
			arr = append(arr, str)
		}
		return arr, nil
	case []interface{}:
		var arr []string
		results := rr.Value.([]interface{})
		for _, v := range results {

			str := fmt.Sprint(v)
			arr = append(arr, str)
		}
		return arr, nil
	}

	return nil, errors.New("Result is not string array format")
}

func (rr redisResult) StringMap() (map[string]string, error) {
	switch rr.Value.(type) {
	case error:
		return nil, rr.Value.(error)
	case map[string]string:
		return rr.Value.(map[string]string), nil
	case []Result:
		results := rr.Value.([]Result)

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

	case []string:
		arr := rr.Value.([]string)
		length := len(arr)

		m := make(map[string]string)
		for i := 0; i < length; i++ {
			k := arr[i]

			i++

			v := arr[i]
			m[k] = v
		}
		return m, nil
	case []interface{}:

		m := make(map[string]string)
		results := rr.Value.([]interface{})
		if len(results)%2 != 0 {
			return nil, errors.New("Result is not string map format")
		}
		for i := 0; i < len(results); i++ {
			k := fmt.Sprint(results[i])
			i++
			v := fmt.Sprint(results[i])
			m[k] = v
		}
		return m, nil
	}

	return nil, errors.New("Result is not string map format")
}

func (rr redisResult) Results() ([]Result, error) {
	switch rr.Value.(type) {
	case []Result:
		return rr.Value.([]Result), nil
	}
	return nil, errors.New("Result is not Array of result format")
}

func (rr redisResult) Bool() (bool, error) {
	switch val := rr.Value.(type) {
	case bool:
		return val, nil
	case int:
		return val != 0, nil
	}
	return false, errors.New("Result is not Array of result format")
}
