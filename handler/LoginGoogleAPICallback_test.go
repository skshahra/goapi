package handler

import (
	"net/http"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestDatabaseCollections_GoogleLoginAPICallback(t *testing.T) {
	type fields struct {
		MySqlDB *gorm.DB
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			H := &DatabaseCollections{
				MySqlDB: tt.fields.MySqlDB,
			}
			H.GoogleLoginAPICallback(tt.args.w, tt.args.r)
		})
	}
}

func Test_getUserInfo(t *testing.T) {
	type args struct {
		state string
		code  string
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getUserInfo(tt.args.state, tt.args.code)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUserInfo() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getUserInfo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
