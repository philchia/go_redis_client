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

func (rr redisResult) String() (string, error) {
	switch rr.Value.(type) {
	case error:
		return "", rr.Value.(error)
	case string:
		return rr.Value.(string), nil
	case []byte:
		return string(rr.Value.([]byte)), nil
	case interface{}:
		return fmt.Sprint(rr.Value.(interface{})), nil
	}
	return "", errors.New("Result is not string format")
}

func (rr redisResult) Int() (int, error) {
	switch rr.Value.(type) {
	case error:
		return -1, rr.Value.(error)
	case int:
		return rr.Value.(int), nil
	case string:
		return strconv.Atoi(rr.Value.(string))
	}
	return -1, errors.New("Result is not int format")
}

func (rr redisResult) Int32() (int32, error) {
	switch rr.Value.(type) {
	case error:
		return -1, rr.Value.(error)
	case int:
		return int32(rr.Value.(int)), nil
	case int64:
		return int32(rr.Value.(int64)), nil
	case int32:
		return rr.Value.(int32), nil
	case string:
		i, err := strconv.Atoi(rr.Value.(string))
		if err != nil {
			return -1, err
		}
		return int32(i), nil
	}
	return -1, errors.New("Result is not int32 format")
}

func (rr redisResult) Int64() (int64, error) {
	switch rr.Value.(type) {
	case error:
		return -1, rr.Value.(error)
	case int:
		return int64(rr.Value.(int)), nil
	case int64:
		return rr.Value.(int64), nil
	case int32:
		return int64(rr.Value.(int32)), nil
	case string:
		i, err := strconv.Atoi(rr.Value.(string))
		if err != nil {
			return -1, err
		}
		return int64(i), nil
	}
	return -1, errors.New("Result is not int64 format")
}

func (rr redisResult) Float32() (float32, error) {
	switch rr.Value.(type) {
	case error:
		return -1, rr.Value.(error)
	case int:
		return float32(rr.Value.(int)), nil
	case int64:
		return float32(rr.Value.(int64)), nil
	case int32:
		return float32(rr.Value.(int32)), nil
	case string:
		f, err := strconv.ParseFloat(rr.Value.(string), 32)
		if err != nil {
			return -1, err
		}
		return float32(f), nil
	}
	return -1, errors.New("Result is not float32 format")
}

func (rr redisResult) Float64() (float64, error) {
	switch rr.Value.(type) {
	case error:
		return -1, rr.Value.(error)
	case int:
		return float64(rr.Value.(int)), nil
	case int64:
		return float64(rr.Value.(int64)), nil
	case int32:
		return float64(rr.Value.(int32)), nil
	case string:
		f, err := strconv.ParseFloat(rr.Value.(string), 64)
		if err != nil {
			return -1, err
		}
		return f, nil
	}
	return -1, errors.New("Result is not float64 format")
}

func (rr redisResult) StringArray() ([]string, error) {
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

func (rr redisResult) Array() ([]Result, error) {
	switch rr.Value.(type) {
	case []Result:
		return rr.Value.([]Result), nil
	}
	return nil, errors.New("Result is not Array of result format")
}

func (rr redisResult) Bool() (bool, error) {
	switch rr.Value.(type) {
	case bool:
		return rr.Value.(bool), nil
	case int:
		return rr.Value.(int) != 0, nil
	case string:
		return rr.Value.(string) == "OK", nil
	}
	return false, errors.New("Result is not Array of result format")
}
