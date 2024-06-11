package cart

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

var entities = []*Entity{&Entity{
	ID:          "cart1",
	CustomersID: "customer1",
	TotalPrice:  100,
	IsCompleted: false,
},
	&Entity{
		ID:          "cart2",
		CustomersID: "customer2",
		TotalPrice:  100,
		IsCompleted: false,
	},
}

var entity2 = &Entity{
	ID:          "cart2",
	CustomersID: "customer2",
	TotalPrice:  100,
	IsCompleted: false,
}

func (m *serviceMock) GetAll() []*Entity {
	args := m.Called()
	return args.Get(0).([]*Entity)
}

func (m *serviceMock) GetById(id string) (*Entity, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Entity), args.Error(1)
}

func TestHandler_CreateCartHandler(t *testing.T) {
	type fields struct {
		service *serviceMock
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		cartId         string
		wantStatusCode int
	}{
		{
			name:   "cart is not found, returns 404",
			fields: fields{service: &serviceMock{}},
			args: args{
				r: httptest.NewRequest("GET", "/cart/notValid", strings.NewReader(`{}`)),
			},
			cartId:         "notValid",
			wantStatusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := &serviceMock{}
			m.On("GetById", mock.MatchedBy(func(id string) bool {
				return !slices.ContainsFunc(entities, func(e *Entity) bool {
					return e.ID == tt.cartId
				})
			})).Return(nil, errors.New("could not find cart")).Maybe()

			h := Handler{m}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.GetCartHandler)
			handler.ServeHTTP(rr, tt.args.r)

			if !assert.Equalf(t, tt.wantStatusCode, rr.Code, "status code not as expected") {
				return
			}
		})
	}
}
