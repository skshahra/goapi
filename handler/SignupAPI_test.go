package handler

import (
	"net/http"
	"testing"

	"gorm.io/gorm"
)

func TestDatabaseCollections_SignupAPI(t *testing.T) {
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
			H.SignupAPI(tt.args.w, tt.args.r)
		})
	}
}
