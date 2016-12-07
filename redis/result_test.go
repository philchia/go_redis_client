package redis

import (
	"errors"
	"reflect"
	"testing"
)

func Test_redisResult_Error(t *testing.T) {
	type fields struct {
		Value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"case1",
			fields{
				errors.New("test"),
			},
			true,
		},
		{
			"case2",
			fields{
				"test",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := &redisResult{
				Value: tt.fields.Value,
			}
			if err := rr.Error(); (err != nil) != tt.wantErr {
				t.Errorf("redisResult.Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_redisResult_OK(t *testing.T) {
	type fields struct {
		Value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"case1",
			fields{
				errors.New("test"),
			},
			false,
		},
		{
			"case2",
			fields{
				false,
			},
			false,
		},
		{
			"case3",
			fields{
				true,
			},
			false,
		},
		{
			"case4",
			fields{
				"OK",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := &redisResult{
				Value: tt.fields.Value,
			}
			if got := rr.OK(); got != tt.want {
				t.Errorf("redisResult.OK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisResult_PONG(t *testing.T) {
	type fields struct {
		Value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"case1",
			fields{
				errors.New("test"),
			},
			false,
		},
		{
			"case2",
			fields{
				false,
			},
			false,
		},
		{
			"case3",
			fields{
				"PONG",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := &redisResult{
				Value: tt.fields.Value,
			}
			if got := rr.PONG(); got != tt.want {
				t.Errorf("redisResult.PONG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisResult_String(t *testing.T) {
	type fields struct {
		Value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			"case1",
			fields{
				"PONG",
			},
			"PONG",
			false,
		},
		{
			"case2",
			fields{
				errors.New("test"),
			},
			"",
			true,
		},
		{
			"case3",
			fields{
				[]byte("test"),
			},
			"test",
			false,
		},
		{
			"case4",
			fields{
				int8(8),
			},
			"8",
			false,
		},
		{
			"case5",
			fields{
				int16(8),
			},
			"8",
			false,
		},
		{
			"case6",
			fields{
				int(8),
			},
			"8",
			false,
		},
		{
			"case7",
			fields{
				int32(8),
			},
			"8",
			false,
		},
		{
			"case8",
			fields{
				int64(8),
			},
			"8",
			false,
		},
		{
			"case9",
			fields{
				float64(8.8),
			},
			"8.8",
			false,
		},
		{
			"case10",
			fields{
				[]string{"hello"},
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := &redisResult{
				Value: tt.fields.Value,
			}
			got, err := rr.String()
			if (err != nil) != tt.wantErr {
				t.Errorf("redisResult.String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("redisResult.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisResult_Int(t *testing.T) {
	type fields struct {
		Value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			"case1",
			fields{
				"PONG",
			},
			0,
			true,
		},
		{
			"case2",
			fields{
				errors.New("test"),
			},
			0,
			true,
		},
		{
			"case3",
			fields{
				[]byte("test"),
			},
			0,
			true,
		},
		{
			"case4",
			fields{
				int8(8),
			},
			8,
			false,
		},
		{
			"case5",
			fields{
				int16(8),
			},
			8,
			false,
		},
		{
			"case6",
			fields{
				int(8),
			},
			8,
			false,
		},
		{
			"case7",
			fields{
				int32(8),
			},
			8,
			false,
		},
		{
			"case8",
			fields{
				int64(8),
			},
			8,
			false,
		},
		{
			"case9",
			fields{
				float64(8.8),
			},
			0,
			true,
		},
		{
			"case10",
			fields{
				[]string{"hello"},
			},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := &redisResult{
				Value: tt.fields.Value,
			}
			got, err := rr.Int()
			if (err != nil) != tt.wantErr {
				t.Errorf("redisResult.Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("redisResult.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisResult_Float64(t *testing.T) {
	type fields struct {
		Value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			"case1",
			fields{
				"PONG",
			},
			0,
			true,
		},
		{
			"case2",
			fields{
				errors.New("test"),
			},
			0,
			true,
		},
		{
			"case3",
			fields{
				"8",
			},
			8,
			false,
		},
		{
			"case4",
			fields{
				int8(8),
			},
			8,
			false,
		},
		{
			"case5",
			fields{
				int16(8),
			},
			8,
			false,
		},
		{
			"case6",
			fields{
				int(8),
			},
			8,
			false,
		},
		{
			"case7",
			fields{
				int32(8),
			},
			8,
			false,
		},
		{
			"case8",
			fields{
				int64(8),
			},
			8,
			false,
		},
		{
			"case8",
			fields{
				float64(8.8),
			},
			0,
			true,
		},
		{
			"case8",
			fields{
				[]string{"hello"},
			},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := &redisResult{
				Value: tt.fields.Value,
			}
			got, err := rr.Float64()
			if (err != nil) != tt.wantErr {
				t.Errorf("redisResult.Float64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("redisResult.Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisResult_Strings(t *testing.T) {
	type fields struct {
		Value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			"case1",
			fields{
				[]string{"1", "2"},
			},
			[]string{"1", "2"},
			false,
		},
		{
			"case2",
			fields{
				errors.New("test"),
			},
			nil,
			true,
		},
		{
			"case3",
			fields{
				[]Result{&redisResult{"hello"}},
			},
			[]string{"hello"},
			false,
		},
		{
			"case4",
			fields{
				"test",
			},
			nil,
			true,
		},
		{
			"case5",
			fields{
				[]Result{&redisResult{"hello"}, &redisResult{errors.New("test")}},
			},
			nil,
			true,
		},
		{
			"case6",
			fields{
				[]interface{}{"hello", "test"},
			},
			[]string{"hello", "test"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := &redisResult{
				Value: tt.fields.Value,
			}
			got, err := rr.Strings()
			if (err != nil) != tt.wantErr {
				t.Errorf("redisResult.Strings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("redisResult.Strings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisResult_StringMap(t *testing.T) {
	type fields struct {
		Value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]string
		wantErr bool
	}{
		{
			"case1",
			fields{
				[]string{"1", "2"},
			},
			map[string]string{"1": "2"},
			false,
		},
		{
			"case2",
			fields{
				map[string]string{"1": "2"},
			},
			map[string]string{"1": "2"},
			false,
		},
		{
			"case3",
			fields{
				[]Result{&redisResult{"hello"}, &redisResult{errors.New("test")}},
			},
			nil,
			true,
		},
		{
			"case4",
			fields{
				"test",
			},
			nil,
			true,
		},
		{
			"case5",
			fields{
				[]Result{&redisResult{"hello"}, &redisResult{errors.New("test")}},
			},
			nil,
			true,
		},
		{
			"case6",
			fields{
				[]Result{&redisResult{"hello"}, &redisResult{"hello"}, &redisResult{errors.New("test")}},
			},
			nil,
			true,
		},
		{
			"case7",
			fields{
				[]interface{}{"hello", "test"},
			},
			map[string]string{"hello": "test"},
			false,
		},
		{
			"case8",
			fields{
				[]interface{}{&redisResult{"hello"}, &redisResult{errors.New("test")}, &redisResult{"hello"}},
			},
			nil,
			true,
		},
		{
			"case9",
			fields{
				[]Result{&redisResult{"hello"}, &redisResult{"hello"}, &redisResult{"hello"}},
			},
			nil,
			true,
		},
		{
			"case10",
			fields{
				errors.New("test"),
			},
			nil,
			true,
		},
		{
			"case11",
			fields{
				[]Result{&redisResult{"hello"}, &redisResult{errors.New("test")}},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := &redisResult{
				Value: tt.fields.Value,
			}
			got, err := rr.StringMap()
			if (err != nil) != tt.wantErr {
				t.Errorf("redisResult.StringMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("redisResult.StringMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisResult_Results(t *testing.T) {
	type fields struct {
		Value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Result
		wantErr bool
	}{
		{
			"case1",
			fields{
				[]string{"1", "2"},
			},
			nil,
			true,
		},
		{
			"case2",
			fields{
				errors.New("test"),
			},
			nil,
			true,
		},
		{
			"case3",
			fields{
				[]Result{&redisResult{"hello"}},
			},
			[]Result{&redisResult{"hello"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := &redisResult{
				Value: tt.fields.Value,
			}
			got, err := rr.Results()
			if (err != nil) != tt.wantErr {
				t.Errorf("redisResult.Results() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("redisResult.Results() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisResult_Bool(t *testing.T) {
	type fields struct {
		Value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			"case1",
			fields{
				"PONG",
			},
			false,
			true,
		},
		{
			"case2",
			fields{
				true,
			},
			true,
			false,
		},
		{
			"case3",
			fields{
				0,
			},
			false,
			false,
		},
		{
			"case4",
			fields{
				1,
			},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := &redisResult{
				Value: tt.fields.Value,
			}
			got, err := rr.Bool()
			if (err != nil) != tt.wantErr {
				t.Errorf("redisResult.Bool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("redisResult.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisResult_Message(t *testing.T) {
	type fields struct {
		Value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    Message
		wantErr bool
	}{
		{
			"case1",
			fields{
				"hello",
			},
			Message{},
			true,
		},
		{
			"case2",
			fields{
				[]Result{&redisResult{"hello"}},
			},
			Message{},
			true,
		},
		{
			"case3",
			fields{
				[]Result{&redisResult{"subscribe"}, &redisResult{"name"}, &redisResult{1}, &redisResult{"subscribe"}},
			},
			Message{Type: "subscribe", Pattern: "", Channel: "name", Count: 1},
			false,
		},
		{
			"case4",
			fields{
				[]Result{&redisResult{"psubscribe"}, &redisResult{"name*"}, &redisResult{1}, &redisResult{"subscribe"}},
			},
			Message{Type: "psubscribe", Pattern: "", Channel: "name*", Count: 1},
			false,
		},
		{
			"case5",
			fields{
				[]Result{&redisResult{"psubscribe"}},
			},
			Message{},
			true,
		},
		{
			"case6",
			fields{
				[]Result{&redisResult{"psubscribe"}, &redisResult{errors.New("test")}},
			},
			Message{},
			true,
		},
		{
			"case7",
			fields{
				[]Result{&redisResult{errors.New("test")}, &redisResult{errors.New("test")}},
			},
			Message{},
			true,
		},
		{
			"case8",
			fields{
				[]Result{&redisResult{"psubscribe"}, &redisResult{"name*"}, &redisResult{"name*"}, &redisResult{"subscribe"}},
			},
			Message{},
			true,
		},
		{
			"case9",
			fields{
				[]Result{&redisResult{"message"}},
			},
			Message{},
			true,
		},
		{
			"case10",
			fields{
				[]Result{&redisResult{"message"}, &redisResult{"name"}, &redisResult{"phil"}},
			},
			Message{Type: "message", Channel: "name", Data: "phil"},
			false,
		},
		{
			"case11",
			fields{
				[]Result{&redisResult{"pmessage"}},
			},
			Message{},
			true,
		},
		{
			"case12",
			fields{
				[]Result{&redisResult{"pmessage"}, &redisResult{"name*"}, &redisResult{"name"}, &redisResult{"phil"}},
			},
			Message{Type: "pmessage", Pattern: "name*", Channel: "name", Data: "phil"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := &redisResult{
				Value: tt.fields.Value,
			}
			got, err := rr.Message()
			if (err != nil) != tt.wantErr {
				t.Errorf("redisResult.Message() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("redisResult.Message() = %v, want %v", got, tt.want)
			}
		})
	}
}
