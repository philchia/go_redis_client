package redis

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

var (
	errResultFormatMismatch = errors.New("format mismatch")
)

// Result interface
type Result interface {
	Error() error
	OK() bool
	PONG() bool
	String() (string, error)
	Strings() ([]string, error)
	StringMap() (map[string]string, error)
	Int() (int, error)
	Float64() (float64, error)
	Results() ([]Result, error)
	Bool() (bool, error)
	Message() (Message, error)
}

type redisResult struct {
	Value interface{}
}

func (rr *redisResult) Error() error {
	if err, ok := rr.Value.(error); ok {
		return err
	}
	return nil
}

func (rr *redisResult) OK() bool {
	return reflect.DeepEqual(rr, OK)
}

func (rr *redisResult) PONG() bool {
	return reflect.DeepEqual(rr, PONG)
}

func (rr *redisResult) String() (string, error) {
	switch val := rr.Value.(type) {
	case error:
		return "", val
	case string:
		return val, nil
	case []byte:
		return bytes2str(val), nil
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
	}
	return "", errResultFormatMismatch
}

func (rr *redisResult) Int() (int, error) {
	switch val := rr.Value.(type) {
	case error:
		return 0, val
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
		return strconv.Atoi(bytes2str(val))
	}
	return 0, errResultFormatMismatch
}

func (rr *redisResult) Float64() (float64, error) {
	switch val := rr.Value.(type) {
	case error:
		return 0, val
	case int8:
		return float64(val), nil
	case int16:
		return float64(val), nil
	case int:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case string:
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0, err
		}
		return f, nil
	}
	return 0, errResultFormatMismatch
}

func (rr *redisResult) Strings() ([]string, error) {
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

func (rr *redisResult) StringMap() (map[string]string, error) {
	switch r := rr.Value.(type) {
	case error:
		return nil, r

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
		if length%2 != 0 {
			return nil, errResultFormatMismatch
		}
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

func (rr *redisResult) Results() ([]Result, error) {
	switch r := rr.Value.(type) {
	case []Result:
		return r, nil
	}
	return nil, errResultFormatMismatch
}

func (rr *redisResult) Bool() (bool, error) {
	switch val := rr.Value.(type) {
	case bool:
		return val, nil
	case int:
		return val != 0, nil
	}
	return false, errResultFormatMismatch
}

func (rr *redisResult) Message() (Message, error) {

	res, err := rr.Results()
	if err != nil {
		return Message{}, err
	}

	t, err := res[0].String()
	if err != nil {
		return Message{}, err
	}

	switch t {
	case "subscribe", "psubscribe", "unsubscribe", "punsubscribe":
		if len(res) < 3 {
			return Message{}, errResultFormatMismatch
		}
		channel, err := res[1].String()
		if err != nil {
			return Message{}, err
		}
		count, err := res[2].Int()
		if err != nil {
			return Message{}, err
		}
		msg := Message{Type: t, Channel: channel, Count: count}
		return msg, nil

	case "message":
		if len(res) < 3 {
			return Message{}, errResultFormatMismatch
		}
		channel, err := res[1].String()
		if err != nil {
			return Message{}, err
		}
		data, err := res[2].String()
		if err != nil {
			return Message{}, err
		}
		msg := Message{Type: t, Channel: channel, Data: data}
		return msg, nil
	case "pmessage":
		if len(res) < 4 {
			return Message{}, errResultFormatMismatch
		}
		pattern, err := res[1].String()
		if err != nil {
			return Message{}, err
		}
		channel, err := res[2].String()
		if err != nil {
			return Message{}, err
		}
		data, err := res[3].String()
		if err != nil {
			return Message{}, err
		}
		msg := Message{Type: t, Pattern: pattern, Channel: channel, Data: data}
		return msg, nil

	default:
		return Message{}, errResultFormatMismatch

	}
}
