package handlers

import (
	"context"
	"testing"

	"tServerOra/internal/models"

	"github.com/stretchr/testify/mock"
)

var repoMock *RepoMock

type RepoMock struct {
	mock.Mock
}

func (r *RepoMock) SaveCard(ctx context.Context, cTC *models.CardTC) error {
	args := r.Called(ctx, cTC)
	return args.Error(0)
}

func (r *RepoMock) CheckDBConnection(context.Context) error {
	return nil
}

func (r *RepoMock) CreateUser(context.Context) (string, error) {
	return "", nil
}

func TestHandlerApiUrlPost(t *testing.T) {
	// str := &struct {
	// 	URL string
	// }{
	// 	URL: "www.example.com",
	// }
	// bOut, err := json.Marshal(str)
	// if err != nil {
	// 	t.Error("Ошибка серилизации")
	// }
	// ctx := context.WithValue(context.Background(), models.UserKey, "aasdasdSQW")
	// repoMock.On("SaveURL", ctx, "www.example.com", opt.RespBaseURL()+"/", "aasdasdSQW").Return(opt.RespBaseURL()+"/123123asdasd", nil)
	// handler := http.HandlerFunc(HandlerAPIURLPost)
	// r := httptest.NewRequest("POST", "http://localhost:8080", bytes.NewBuffer(bOut))
	// w := httptest.NewRecorder()

	// handler.ServeHTTP(w, r.WithContext(ctx))
	// res := w.Result()
	// b, _ := io.ReadAll(res.Body)
	// defer res.Body.Close()
	// assert.Equal(t, http.StatusCreated, res.StatusCode, "Не верный код ответа POST")
	// assert.Equal(t, `{"result":"`+opt.RespBaseURL()+`/123123asdasd"}`, string(b), "Не верный ответ POST")

}

func InitMocks() {
	repoMock = new(RepoMock)
	NewHandlers(repoMock)
}
