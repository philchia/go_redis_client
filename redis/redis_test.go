package redis

import "testing"

func TestConnect(t *testing.T) {
	type args struct {
		host    string
		port    string
		options []*Option
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"case1",
			args{
				"127.0.0.1",
				"6379",
				nil,
			},
			false,
		},
		{
			"case2",
			args{
				"tcp://a.a.a.a",
				"bbb",
				nil,
			},
			true,
		},
		{
			"case3",
			args{
				"127.0.0.1",
				"6666",
				nil,
			},
			true,
		},
		{
			"case4",
			args{
				"127.0.0.1",
				"6379",
				[]*Option{&Option{Auth: "password"}},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Connect(tt.args.host, tt.args.port, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
