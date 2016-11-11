package redis

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	errResultFormatMismatch = errors.New("format mismatch")
)

type redisResult struct {
	Value interface{}
	Err   error
}

func (rr redisResult) Error() error {
	return rr.Err
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
	return "", errResultFormatMismatch
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
	return -1, errResultFormatMismatch
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
	return -1, errResultFormatMismatch
}

func (rr redisResult) Strings() ([]string, error) {
	switch r := rr.Value.(type) {
	case error:
		return nil, rr.Value.(error)
	case []string:
		return r, nil
	case []Result:
		var arr []string

		for _, res := range r {
			str, err := res.String()
			if err != nil {
				return nil, err
			}
			arr = append(arr, str)
		}
		return arr, nil
	case []interface{}:
		var arr []string

		for _, res := range r {
			str := fmt.Sprint(res)
			arr = append(arr, str)
		}
		return arr, nil
	}

	return nil, errResultFormatMismatch
}

func (rr redisResult) StringMap() (map[string]string, error) {
	switch r := rr.Value.(type) {
	case error:
		return nil, rr.Value.(error)
	case map[string]string:
		return r, nil
	case []Result:

		length := len(r)
		if length%2 != 0 {
			return nil, errResultFormatMismatch
		}
		m := make(map[string]string)
		for i := 0; i < length; i++ {
			k, err := r[i].String()
			if err != nil {
				return nil, err
			}
			i++
			v, err := r[i].String()
			if err != nil {
				return nil, err
			}
			m[k] = v
		}
		return m, nil

	case []string:

		length := len(r)

		m := make(map[string]string)
		for i := 0; i < length; i++ {
			k := r[i]

			i++

			v := r[i]
			m[k] = v
		}
		return m, nil
	case []interface{}:

		m := make(map[string]string)
		if len(r)%2 != 0 {
			return nil, errResultFormatMismatch
		}
		for i := 0; i < len(r); i++ {
			k := fmt.Sprint(r[i])
			i++
			v := fmt.Sprint(r[i])
			m[k] = v
		}
		return m, nil
	}

	return nil, errResultFormatMismatch
}

func (rr redisResult) Results() ([]Result, error) {
	switch r := rr.Value.(type) {
	case []Result:
		return r, nil
	}
	return nil, errResultFormatMismatch
}

func (rr redisResult) Bool() (bool, error) {
	switch val := rr.Value.(type) {
	case bool:
		return val, nil
	case int:
		return val != 0, nil
	}
	return false, errResultFormatMismatch
}
