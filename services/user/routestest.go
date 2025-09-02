package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sikozonpc/ecom/types"
)

func TestUserserviceHandler(t *testing.T) {
	userstore := &mokeUserStore{}
	handler := NewHandler(userstore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "123",
			Email:     "",
			Password:  "asd",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err!= nil{
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		
	})
}

type mokeUserStore struct{}

func (m *mokeUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mokeUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
func (m *mokeUserStore) CreateUser(types.User) error {
	return nil
}
