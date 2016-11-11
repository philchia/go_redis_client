package redis

import "testing"

func TestResultString(t *testing.T) {
	res := []redisResult{
		redisResult{
			Value: "hello",
		},
		redisResult{
			Value: 12,
		},
		redisResult{
			Value: int8(12),
		},
		redisResult{
			Value: int16(12),
		},
		redisResult{
			Value: int32(12),
		},
		redisResult{
			Value: int64(12),
		},
		redisResult{
			Value: float32(3.14),
		},
		redisResult{
			Value: float64(3.14),
		},
		redisResult{
			Value: []byte("HELLO"),
		},
	}

	for _, v := range res {
		_, err := v.String()
		if err != nil {
			t.Fail()
		}
	}
}

func TestResultInt(t *testing.T) {
	res := []redisResult{
		redisResult{
			Value: "12",
		},
		redisResult{
			Value: 12,
		},
		redisResult{
			Value: int8(12),
		},
		redisResult{
			Value: int16(12),
		},
		redisResult{
			Value: int32(12),
		},
		redisResult{
			Value: int64(12),
		},
		redisResult{
			Value: []byte("12"),
		},
	}

	for _, v := range res {
		_, err := v.Int()
		if err != nil {
			t.Fail()
		}
	}
}

func TestResultBool(t *testing.T) {
	res := []redisResult{
		redisResult{
			Value: "1",
		},
		redisResult{
			Value: 1,
		},
		redisResult{
			Value: int8(2),
		},
		redisResult{
			Value: int16(12),
		},
		redisResult{
			Value: int32(12),
		},
		redisResult{
			Value: int64(12),
		},
		redisResult{
			Value: float32(3.14),
		},
		redisResult{
			Value: float64(3.14),
		},
		redisResult{
			Value: []byte("HELLO"),
		},
	}

	for _, v := range res {
		_, err := v.String()
		if err != nil {
			t.Fail()
		}
	}
}
