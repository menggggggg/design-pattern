package factory

import (
	"reflect"
	"testing"
)

func TestNewIConfigParserMap(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IConfigParserMap
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: &JsonConfigParse{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: &YamlConfigParse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIConfigParserMap(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIConfigParserMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIConfigParserSlice(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IConfigParserSlice
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: &JsonConfigParse{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: &YamlConfigParse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIConfigParserSlice(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
