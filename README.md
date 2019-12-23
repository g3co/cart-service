### Install

`make`

### Run

`bin/cart-service -c cfg/config.example.yml`

### Test

`make test`

### Routes

`POST@/addToCart` Data example `{"items":[{"quantity":1,"itemId":1}]}`

`GET@/getCart`

`DELETE@/cart`