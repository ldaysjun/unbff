package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/graphql-go/handler"
	"github.com/ldaysjun/unbff/config"
	"github.com/ldaysjun/unbff/engine"
	"net/http"
)

func main() {
	kernel, err := engine.NewKernel(config.ForTest())
	if err != nil {
		panic(err)
	}
	schema, err := kernel.NewSchemaWithApp(context.Background(), "app")
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	if err != nil {
		panic(err)
	}
	http.Handle("/graphql", h)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

//s := `{getInboundOrderList {
//		ksInboundOrder_createUser
//		ksInboundOrder_updateUser
//		ksInboundOrder_createTime
//		ksInboundOrder_updateTime
//		ksInboundOrder_id
//		ksInboundOrder_orderNo
//		ksInboundOrder_customerOrderNo
//	 	ksInboundOrder_lpnCode
//		ksInboundOrder_warehouseCode
//		ksInboundOrder_warehouseName
//		ksInboundOrder_customerCode
//		ksInboundOrder_customerName
//		ksInboundOrder_supplier
//		ksInboundOrder_inboundOrderType
//		ksInboundOrder_operationType
//		ksInboundOrder_inboundReceipt
//	 	ksInboundOrder_inboundOrderStatus
//		ksInboundOrder_version
//		ksInboundOrder_stationCode
//		ksInboundOrder_estimatedArrivalTime
//		ksInboundOrder_remark
//		ksInboundOrder_abnormalCn
//		ksInboundOrder_inboundOrderStatusCn
//		ksInboundOrder_stationCodes
//		ksInboundOrder_ksInboundOrderContainerDetailIds
//}}`
