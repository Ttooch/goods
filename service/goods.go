package service

import (
	"fmt"
	"github.com/micro/examples/pubsub/srv/proto"
	"github.com/pborman/uuid"
	"time"
	goods "github.com/ttooch/proto/goods"
)

func GoodsDetail() {
	req := rpcClient.NewRequest("Goods.Detail", goodsRpc, &goods.DetailRequest{
		BarCode: "111",
	})

	rsp := &goods.DetailResponse{}

	// Call service
	if err := rpcClient.Call(ctx, req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}

	fmt.Println(rsp)
}

func PubGoods() {

	msg := rpcClient.NewPublication(goodsTopic, &pubsub.Event{
		Id:        uuid.NewUUID().String(),
		Timestamp: time.Now().Unix(),
		Message:   fmt.Sprintf("Messaging you all day on %s", "1111"),
	})

	if err := rpcClient.Publish(ctx, msg); err != nil {
		fmt.Println("pub err: ", err)
		return
	}

	fmt.Printf("Published: %v\n", msg)
}
