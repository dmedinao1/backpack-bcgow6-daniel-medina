package products

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetFunStubRepository() Repository {
	return &funRepository{}
}

type funRepository struct {
}

func (f *funRepository) GetAllBySeller(sellerID string) ([]Product, error) {
	return []Product{}, nil
}

func GetErrorStubRepository() Repository {
	return &errorRepository{}
}

type errorRepository struct {
}

func (e errorRepository) GetAllBySeller(sellerID string) ([]Product, error) {
	return nil, errors.New("error in repository")
}

func Test_service_GetAllBySeller(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		sellerID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Product
		wantErr assert.ErrorAssertionFunc
	}{
		{"Ok when valid", fields{repo: GetFunStubRepository()}, args{sellerID: ""}, []Product{}, assert.NoError},
		{"Error when repository has error", fields{repo: GetErrorStubRepository()}, args{sellerID: ""}, nil, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.GetAllBySeller(tt.args.sellerID)
			if !tt.wantErr(t, err, fmt.Sprintf("GetAllBySeller(%v)", tt.args.sellerID)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetAllBySeller(%v)", tt.args.sellerID)
		})
	}
}
