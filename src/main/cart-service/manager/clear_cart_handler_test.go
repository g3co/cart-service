package manager

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
	"main/cart-service/repository/mock"
	"main/cart-service/structs"
	"testing"
)

func TestManager_ClearCartHandler(t *testing.T) {
	Convey("Test ClearCartHandler", t, func() {
		ctrl := gomock.NewController(t)

		Reset(func() {
			ctrl.Finish()
		})

		repo := mock_repository.NewMockIRepository(ctrl)
		client := structs.Client{UserId: 10}
		router := mux.NewRouter()

		m := new(Manager)
		m.Repo = repo
		m.Config = *new(structs.Config)
		m.Router = router

		Convey("Wrong incoming data", func() {
			data := []byte("wrong json")
			repo.EXPECT().ClearCart(client.UserId).Return(nil)

			res, err := m.ClearCartHandler(client, data)

			So(err, ShouldBeNil)
			So(res.GetJson(), ShouldEqual, `{"status":true}`)
		})

		Convey("Db error", func() {
			data := []byte("wrong json")
			repo.EXPECT().ClearCart(client.UserId).Return(errors.New("some error"))

			_, err := m.ClearCartHandler(client, data)

			So(err, ShouldNotBeNil)
		})

	})
}
