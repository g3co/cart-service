package manager

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
	"main/cart-service/repository/mock"
	"main/cart-service/structs"
	"testing"
)

func TestManager_AddToCartHandler(t *testing.T) {
	Convey("Test AddToCartHandler", t, func() {
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
		request := `{"items":[{"quantity":1,"itemId":1}]}`

		Convey("Wrong incoming data", func() {
			data := []byte("wrong json")

			_, err := m.AddToCartHandler(client, data)

			So(err, ShouldNotBeNil)
		})

		Convey("db error", func() {
			data := []byte(request)
			var rq structs.AddCartRequest

			_ = json.Unmarshal(data, &rq)

			repo.EXPECT().AddToCart(client.UserId, rq.Items).Return(errors.New("some error"))

			_, err := m.AddToCartHandler(client, data)

			So(err, ShouldNotBeNil)
		})

		Convey("success", func() {
			data := []byte(request)
			var rq structs.AddCartRequest

			_ = json.Unmarshal(data, &rq)

			repo.EXPECT().AddToCart(client.UserId, rq.Items).Return(nil)

			res, err := m.AddToCartHandler(client, data)

			So(err, ShouldBeNil)
			So(res.GetJson(), ShouldEqual, `{"status":true}`)
		})
	})
}
