package main

import (
	"encoding/json"
	"fmt"
	"github.com/oliveagle/jsonpath"
	"strings"
)

func main() {
	str := `{
    "code":"0",
    "msg":"成功",
    "data":{
        "total":"55",
        "results":[
            {
                "ksInboundOrder_createUser":"admin",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666157072960",
                "ksInboundOrder_updateTime":"1666157155326",
                "ksInboundOrder_id":"780101148283400192",
                "ksInboundOrder_orderNo":"ASN-20221019-00027",
                "ksInboundOrder_customerOrderNo":"",
                "ksInboundOrder_lpnCode":"T10193003",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"Return",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "780101148316954624"
                ],
                "ksInboundOrder_inboundOrderStatus":"PUTTED_AWAY",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"",
                "ksInboundOrder_inboundOrderStatusCn":"上架完成"
            },
            {
                "ksInboundOrder_createUser":"admin",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666156543182",
                "ksInboundOrder_updateTime":"1666157155935",
                "ksInboundOrder_id":"780098926233415680",
                "ksInboundOrder_orderNo":"ASN-20221019-00025",
                "ksInboundOrder_customerOrderNo":"",
                "ksInboundOrder_lpnCode":"T10193002",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"Return",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "780098926279553024"
                ],
                "ksInboundOrder_inboundOrderStatus":"PUTTED_AWAY",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"",
                "ksInboundOrder_inboundOrderStatusCn":"上架完成"
            },
            {
                "ksInboundOrder_createUser":"admin",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666155912868",
                "ksInboundOrder_updateTime":"1666156019978",
                "ksInboundOrder_id":"780096282504884224",
                "ksInboundOrder_orderNo":"ASN-20221019-00023",
                "ksInboundOrder_customerOrderNo":"",
                "ksInboundOrder_lpnCode":"T10192003",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"Return",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "780096282513272832"
                ],
                "ksInboundOrder_inboundOrderStatus":"PUTTED_AWAY",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"",
                "ksInboundOrder_inboundOrderStatusCn":"上架完成"
            },
            {
                "ksInboundOrder_createUser":"admin",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666155775888",
                "ksInboundOrder_updateTime":"1666155822692",
                "ksInboundOrder_id":"780095707969122304",
                "ksInboundOrder_orderNo":"ASN-20221019-00022",
                "ksInboundOrder_customerOrderNo":"",
                "ksInboundOrder_lpnCode":"T10192001",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"Return",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "780095707985899520"
                ],
                "ksInboundOrder_inboundOrderStatus":"PUTTED_AWAY",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"",
                "ksInboundOrder_inboundOrderStatusCn":"上架完成"
            },
            {
                "ksInboundOrder_createUser":"admin",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666151671305",
                "ksInboundOrder_updateTime":"1666151756597",
                "ksInboundOrder_id":"780078492100227072",
                "ksInboundOrder_orderNo":"ASN-20221019-00021",
                "ksInboundOrder_customerOrderNo":"",
                "ksInboundOrder_lpnCode":"T10191016",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"Return",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "780078492137975808",
                    "780078492142170112"
                ],
                "ksInboundOrder_inboundOrderStatus":"PUTTED_AWAY",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"",
                "ksInboundOrder_inboundOrderStatusCn":"上架完成"
            },
            {
                "ksInboundOrder_createUser":"admin",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666151554877",
                "ksInboundOrder_updateTime":"1666151556482",
                "ksInboundOrder_id":"780078003765800960",
                "ksInboundOrder_orderNo":"ASN-20221019-00020",
                "ksInboundOrder_customerOrderNo":"",
                "ksInboundOrder_lpnCode":"T10191015",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"Return",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "780078003799355392",
                    "780078003803549696"
                ],
                "ksInboundOrder_inboundOrderStatus":"RECEIVED",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"",
                "ksInboundOrder_inboundOrderStatusCn":"收货完成"
            },
            {
                "ksInboundOrder_createUser":"admin",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666151524094",
                "ksInboundOrder_updateTime":"1666151717544",
                "ksInboundOrder_id":"780077874652540928",
                "ksInboundOrder_orderNo":"ASN-20221019-00019",
                "ksInboundOrder_customerOrderNo":"",
                "ksInboundOrder_lpnCode":"T10191014",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"Return",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "780077874665123840",
                    "780077874669318144"
                ],
                "ksInboundOrder_inboundOrderStatus":"PUTTED_AWAY",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"",
                "ksInboundOrder_inboundOrderStatusCn":"上架完成"
            },
            {
                "ksInboundOrder_createUser":"admin",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666151340245",
                "ksInboundOrder_updateTime":"1666151385933",
                "ksInboundOrder_id":"780077103533944832",
                "ksInboundOrder_orderNo":"ASN-20221019-00018",
                "ksInboundOrder_customerOrderNo":"",
                "ksInboundOrder_lpnCode":"T10191013",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"Return",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "780077103563304960"
                ],
                "ksInboundOrder_inboundOrderStatus":"PUTTED_AWAY",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"",
                "ksInboundOrder_inboundOrderStatusCn":"上架完成"
            },
            {
                "ksInboundOrder_createUser":"admin",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666151325258",
                "ksInboundOrder_updateTime":"1666151385232",
                "ksInboundOrder_id":"780077040673910784",
                "ksInboundOrder_orderNo":"ASN-20221019-00017",
                "ksInboundOrder_customerOrderNo":"",
                "ksInboundOrder_lpnCode":"T10191012",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"Return",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "780077040761991168"
                ],
                "ksInboundOrder_inboundOrderStatus":"PUTTED_AWAY",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"",
                "ksInboundOrder_inboundOrderStatusCn":"上架完成"
            },
            {
                "ksInboundOrder_createUser":"admin",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666151295072",
                "ksInboundOrder_updateTime":"1666151384316",
                "ksInboundOrder_id":"780076914064650240",
                "ksInboundOrder_orderNo":"ASN-20221019-00016",
                "ksInboundOrder_customerOrderNo":"",
                "ksInboundOrder_lpnCode":"T10191011",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"Return",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "780076914073038848"
                ],
                "ksInboundOrder_inboundOrderStatus":"PUTTED_AWAY",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"",
                "ksInboundOrder_inboundOrderStatusCn":"上架完成"
            },
            {
                "ksInboundOrder_createUser":"admin",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666150974222",
                "ksInboundOrder_updateTime":"1666151046971",
                "ksInboundOrder_id":"780075568322211840",
                "ksInboundOrder_orderNo":"ASN-20221019-00015",
                "ksInboundOrder_customerOrderNo":"",
                "ksInboundOrder_lpnCode":"T10191010",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"Return",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "780075568414486528"
                ],
                "ksInboundOrder_inboundOrderStatus":"PUTTED_AWAY",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"",
                "ksInboundOrder_inboundOrderStatusCn":"上架完成"
            },
            {
                "ksInboundOrder_createUser":"admin",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666150950378",
                "ksInboundOrder_updateTime":"1666151046127",
                "ksInboundOrder_id":"780075468313227264",
                "ksInboundOrder_orderNo":"ASN-20221019-00014",
                "ksInboundOrder_customerOrderNo":"",
                "ksInboundOrder_lpnCode":"T10191008",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"Return",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "780075468363558912"
                ],
                "ksInboundOrder_inboundOrderStatus":"PUTTED_AWAY",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"",
                "ksInboundOrder_inboundOrderStatusCn":"上架完成"
            },
            {
                "ksInboundOrder_createUser":"admin",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666150927735",
                "ksInboundOrder_updateTime":"1666151025560",
                "ksInboundOrder_id":"780075373341601792",
                "ksInboundOrder_orderNo":"ASN-20221019-00013",
                "ksInboundOrder_customerOrderNo":"",
                "ksInboundOrder_lpnCode":"T10191007",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"Return",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "780075373391933440"
                ],
                "ksInboundOrder_inboundOrderStatus":"PUTTED_AWAY",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"",
                "ksInboundOrder_inboundOrderStatusCn":"上架完成"
            },
            {
                "ksInboundOrder_createUser":"admin",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666150279589",
                "ksInboundOrder_updateTime":"1666150322978",
                "ksInboundOrder_id":"780072654820241408",
                "ksInboundOrder_orderNo":"ASN-20221019-00012",
                "ksInboundOrder_customerOrderNo":"",
                "ksInboundOrder_lpnCode":"T10191006",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"Return",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "780072654908321792",
                    "780072654920904704"
                ],
                "ksInboundOrder_inboundOrderStatus":"PUTTED_AWAY",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"",
                "ksInboundOrder_inboundOrderStatusCn":"上架完成"
            },
            {
                "ksInboundOrder_createUser":"admin",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666150016278",
                "ksInboundOrder_updateTime":"1666150066981",
                "ksInboundOrder_id":"780071550413860864",
                "ksInboundOrder_orderNo":"ASN-20221019-00011",
                "ksInboundOrder_customerOrderNo":"",
                "ksInboundOrder_lpnCode":"T10191005",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"Return",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "780071550426443776"
                ],
                "ksInboundOrder_inboundOrderStatus":"PUTTED_AWAY",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"",
                "ksInboundOrder_inboundOrderStatusCn":"上架完成"
            },
            {
                "ksInboundOrder_createUser":"admin",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666149508696",
                "ksInboundOrder_updateTime":"1666149696040",
                "ksInboundOrder_id":"780069421460647936",
                "ksInboundOrder_orderNo":"ASN-20221019-00010",
                "ksInboundOrder_customerOrderNo":"",
                "ksInboundOrder_lpnCode":"T10191004",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"Return",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "780069421469036544"
                ],
                "ksInboundOrder_inboundOrderStatus":"PUTTED_AWAY",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"",
                "ksInboundOrder_inboundOrderStatusCn":"上架完成"
            },
            {
                "ksInboundOrder_createUser":"admin",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666149486864",
                "ksInboundOrder_updateTime":"1666149548025",
                "ksInboundOrder_id":"780069329882214400",
                "ksInboundOrder_orderNo":"ASN-20221019-00009",
                "ksInboundOrder_customerOrderNo":"",
                "ksInboundOrder_lpnCode":"T10191003",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"Return",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "780069330008043520"
                ],
                "ksInboundOrder_inboundOrderStatus":"PUTTED_AWAY",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"",
                "ksInboundOrder_inboundOrderStatusCn":"上架完成"
            },
            {
                "ksInboundOrder_createUser":"admin",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666149454137",
                "ksInboundOrder_updateTime":"1666149547345",
                "ksInboundOrder_id":"780069192623616000",
                "ksInboundOrder_orderNo":"ASN-20221019-00008",
                "ksInboundOrder_customerOrderNo":"",
                "ksInboundOrder_lpnCode":"T10191002",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"Return",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "780069192657170432"
                ],
                "ksInboundOrder_inboundOrderStatus":"PUTTED_AWAY",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"",
                "ksInboundOrder_inboundOrderStatusCn":"上架完成"
            },
            {
                "ksInboundOrder_createUser":"admin",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666148975173",
                "ksInboundOrder_updateTime":"1666149048534",
                "ksInboundOrder_id":"780067183702994944",
                "ksInboundOrder_orderNo":"ASN-20221019-00007",
                "ksInboundOrder_customerOrderNo":"",
                "ksInboundOrder_lpnCode":"T10191001",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"Return",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "780067183971430400"
                ],
                "ksInboundOrder_inboundOrderStatus":"PUTTED_AWAY",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"",
                "ksInboundOrder_inboundOrderStatusCn":"上架完成"
            },
            {
                "ksInboundOrder_createUser":"openapi",
                "ksInboundOrder_updateUser":"admin",
                "ksInboundOrder_createTime":"1666111900718",
                "ksInboundOrder_updateTime":"1666113066611",
                "ksInboundOrder_id":"779911682168090624",
                "ksInboundOrder_orderNo":"ASN-20221019-00006",
                "ksInboundOrder_customerOrderNo":"PO_20221019005140",
                "ksInboundOrder_lpnCode":"LNP_20221019005140",
                "ksInboundOrder_warehouseCode":"herry001",
                "ksInboundOrder_warehouseName":"herry001测试仓",
                "ksInboundOrder_customerCode":"herry002",
                "ksInboundOrder_customerName":"herry002",
                "ksInboundOrder_supplier":"",
                "ksInboundOrder_inboundOrderType":"B2C",
                "ksInboundOrder_operationType":"",
                "ksInboundOrder_inboundReceipt":"RECEIVING_MODE",
                "ksInboundOrder_ksInboundOrderContainerDetailIds":[
                    "779911682197450752",
                    "779911682210033664"
                ],
                "ksInboundOrder_abnormal":"FALSE",
                "ksInboundOrder_scanBatchNumber":"FALSE",
                "ksInboundOrder_mergeSameSkuBatch":"TRUE",
                "ksInboundOrder_inboundOrderStatus":"RECEIVING",
                "ksInboundOrder_version":"1",
                "ksInboundOrder_stationCode":"11",
                "ksInboundOrder_attachment":{

                },
                "ksInboundOrder_stationCodes":[

                ],
                "ksInboundOrder_estimatedArrivalTime":"",
                "ksInboundOrder_remark":"",
                "ksInboundOrder_abnormalCn":"否",
                "ksInboundOrder_inboundOrderStatusCn":"收货中"
            }
        ]
    },
    "errorMsg":"",
    "traceId":"02d75cfbfc7b424488652c788567f9cd"
}`
	//m := make(map[string]interface{})
	//if err := json.Unmarshal([]byte(str), &m); err != nil {
	//	return
	//}
	var jsonData interface{}
	json.Unmarshal([]byte(str), &jsonData)

	res, err := jsonpath.JsonPathLookup(jsonData, "$.data.results")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

	////or reuse lookup pattern
	//pat, _ := jsonpath.Compile(`$.store.book[?(@.price < $.expensive)].price`)
	//res, err := pat.Lookup(json_data)
}

func jsonPath(m map[string]interface{}) interface{} {
	path := "data.results"
	pathSplit := strings.Split(path, ".")
	temp := m
	for index, p := range pathSplit {
		v, ok := temp[p]
		if !ok {
			break
		}
		if index == len(pathSplit)-1 {
			return v
		}
		switch value := v.(type) {
		case nil:
			break
		case string:
			break
		case int:
			break
		case float64:
			break
		case []interface{}:
			break
		case map[string]interface{}:
			temp = value
		default:
			fmt.Println(p, "is unknown type", fmt.Sprintf("%T", v))
		}
	}
	return nil
}
