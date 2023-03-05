package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ldaysjun/unbff/config"
	"github.com/ldaysjun/unbff/engine"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	kernel, err := engine.NewKernel(config.ForTest())
	if err != nil {
		panic(err)
	}
	s := `{getInboundOrderList {
    		ksInboundOrder_createUser
    		ksInboundOrder_updateUser
    		ksInboundOrder_createTime
    		ksInboundOrder_updateTime
    		ksInboundOrder_id
    		ksInboundOrder_orderNo
    		ksInboundOrder_customerOrderNo
   		 	ksInboundOrder_lpnCode
    		ksInboundOrder_warehouseCode
    		ksInboundOrder_warehouseName
    		ksInboundOrder_customerCode
    		ksInboundOrder_customerName
    		ksInboundOrder_supplier
    		ksInboundOrder_inboundOrderType
    		ksInboundOrder_operationType
    		ksInboundOrder_inboundReceipt
   		 	ksInboundOrder_inboundOrderStatus
    		ksInboundOrder_version
    		ksInboundOrder_stationCode
    		ksInboundOrder_estimatedArrivalTime
    		ksInboundOrder_remark
    		ksInboundOrder_abnormalCn
    		ksInboundOrder_inboundOrderStatusCn
    		ksInboundOrder_stationCodes
    		ksInboundOrder_ksInboundOrderContainerDetailIds
	}}`
	r := kernel.Do(context.Background(), engine.Params{
		DSL:            s,
		App:            "app",
		VariableValues: nil,
	})

	d, _ := json.Marshal(r)
	fmt.Printf(string(d))
}
