package cart

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/charliegreeny/checkout/pkg/cartLineItem"
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
	IsComplete:  false,
}

func (m *serviceMock) GetById(id string) (*Entity, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Entity), args.Error(1)
}

func (m *serviceMock) Create(input *createInput) (*Entity, error) {
	args := m.Called(input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Entity), args.Error(1)
}

func TestHandler_GetCartHandler(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name           string
		args           args
		cartId         string
		wantStatusCode int
		wantBody       *output
	}{
		{
			name: "cart is not valid, returns 404 status",
			args: args{
				r: httptest.NewRequest("GET", "/cart/notValid", nil),
			},
			cartId:         "notValid",
			wantStatusCode: http.StatusNotFound,
		},
		{
			name: "generic error returned, returns 500 status",
			args: args{
				r: httptest.NewRequest("GET", "/cart/internalErr", nil),
			},
			cartId:         "internalErr",
			wantStatusCode: http.StatusNotFound,
		},
		{
			name: "cart is valid, returns cart and 200 status code",
			args: args{
				r: httptest.NewRequest("GET", "/cart/cart1", nil),
			},
			cartId:         "cart1",
			wantStatusCode: http.StatusOK,
			wantBody:       entity.ToOutput(),
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
				var gotBody *output
				json.NewDecoder(rr.Body).Decode(&gotBody)
				assert.Equalf(t, tt.wantBody, gotBody, "return body not as expected")
			}
		})
	}
}

func TestHandler_CreateCartHandler(t *testing.T) {
	tests := []struct {
		name           string
		reqBody        *createInput
		stubEntity     *Entity
		stubErr        error
		wantStatusCode int
		wantBody       *output
	}{
		{
			name: "valid request, cart return with ",
			reqBody: &createInput{
				CustomerId: "customer_id",
				LineItems: []*cartLineItem.Item{
					{
						ItemSKU:    "A",
						Quantity:   1,
						TotalPrice: 10,
					},
				},
			},
			stubEntity:     entity,
			stubErr:        nil,
			wantStatusCode: 201,
			wantBody:       entity.ToOutput(),
		},
		{
			name: "Invalid request, cart return with ",
			reqBody: &createInput{
				CustomerId: "customer_id",
				LineItems: []*cartLineItem.Item{
					{
						ItemSKU:    "A",
						Quantity:   1,
						TotalPrice: 10,
					},
				},
			},
			stubEntity:     entity,
			stubErr:        errors.New("db Error"),
			wantStatusCode: 500,
			wantBody:       entity.ToOutput(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &serviceMock{}
			m.On("Create", mock.Anything).Return(tt.stubEntity, tt.stubErr).Once()
			h := Handler{m}

			body, _ := json.Marshal(tt.reqBody)

			req, err := http.NewRequest("POST", "/cart/", bytes.NewBuffer(body))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			req.Header.Add("Content-Type", "application/json")
			handler := http.HandlerFunc(h.CreateCartHandler)
			handler.ServeHTTP(rr, req)
			if !assert.Equalf(t, tt.wantStatusCode, rr.Code, "status code not as expected") {
				return
			}
			if tt.wantStatusCode == http.StatusCreated {
				var gotBody *output
				json.NewDecoder(rr.Body).Decode(&gotBody)
				assert.Equalf(t, tt.wantBody, gotBody, "return body not as expected")
			}
		},
		)
	}
}
