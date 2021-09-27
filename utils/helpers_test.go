package utils

import (
	"reflect"
	"testing"
)

func TestJsonToBson(t *testing.T) {
	type args struct {
		payload interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    Map
		wantErr bool
	}{
		{
			name: "happy",
			args: args{
				map[string]interface{}{
					"id":   123,
					"name": "test",
				},
			},
			want: Map{
				"_id":  123,
				"name": "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonToBson(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonToBson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonToBson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBsonToJson(t *testing.T) {
	type args struct {
		payload interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    Map
		wantErr bool
	}{
		{
			name: "happy",
			args: args{
				map[string]interface{}{
					"_id":  123,
					"name": "test",
				},
			},
			want: Map{
				"id":   123,
				"name": "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BsonToJson(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("BsonToJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BsonToJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
