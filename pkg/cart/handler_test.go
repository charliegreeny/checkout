package cart

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/charliegreeny/checkout/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

var entity = &Entity{
	ID:          "cart1",
	CustomersID: "customer1",
	TotalPrice:  100,
	IsCompleted: false,
}

func (m *serviceMock) GetById(id string) (*Entity, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Entity), args.Error(1)
}

func TestHandler_GetCartHandler(t *testing.T) {
	type fields struct {
		service *serviceMock
	}
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name           string
		matchFn        func(id string) bool
		fields         fields
		args           args
		cartId         string
		wantStatusCode int
		wantBody       *Entity
	}{
		{
			name:   "cart is not valid, returns 404 status",
			fields: fields{service: &serviceMock{}},
			args: args{
				r: httptest.NewRequest("GET", "/cart/notValid", nil),
			},
			cartId:         "notValid",
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:   "generic error returned, returns 500 status",
			fields: fields{service: &serviceMock{}},
			args: args{
				r: httptest.NewRequest("GET", "/cart/internalErr", nil),
			},
			cartId:         "internalErr",
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:   "cart is valid, returns cart and 200 status code",
			fields: fields{service: &serviceMock{}},
			args: args{
				r: httptest.NewRequest("GET", "/cart/cart1", nil),
			},
			cartId:         "cart1",
			wantStatusCode: http.StatusOK,
			wantBody:       entity,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := &serviceMock{}
			m.On("GetById", mock.MatchedBy(func(id string) bool {
				return tt.cartId != entity.ID
			})).Return(nil, model.ErrNotFound{Err: errors.New("not found")}).Maybe()

			m.On("GetById", mock.MatchedBy(func(id string) bool {
				return tt.cartId == entity.ID
			})).Return(entity, nil).Maybe()

			m.On("GetById", mock.MatchedBy(func(id string) bool {
				return id == "internalErr"
			})).Return(nil, errors.New("Internal error")).Maybe()

			h := Handler{m}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.GetCartHandler)
			handler.ServeHTTP(rr, tt.args.r)

			if !assert.Equalf(t, tt.wantStatusCode, rr.Code, "status code not as expected") {
				return
			}

			if tt.wantStatusCode == http.StatusOK {
				var gotBody *Entity
				json.NewDecoder(rr.Body).Decode(&gotBody)
				assert.Equalf(t, tt.wantBody, gotBody, "return body not as expected")
			}
		})
	}
}
