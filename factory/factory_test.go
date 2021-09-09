package factory

import (
	"reflect"
	"testing"
)

func TestNewIConfigParser(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IConfigParser
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: &JsonConfigParseInterface{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: &YamlConfigParseInterface{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIConfigParser(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
