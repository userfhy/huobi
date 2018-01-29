package main

import (
    "fmt"
    "time"
    "strconv"
    "encoding/json"
    "github.com/gorilla/websocket"
    "./common"
    "./conf"
)

func main()  {
    dialer := websocket.Dialer{ /* set fields as needed */ }
    ws, _, err := dialer.Dial("wss://api.huobi.pro/ws", nil)
    if err != nil {
    	// handle error
    }

    for {
        if _,p,err := ws.ReadMessage();err == nil {
            res := common.UnGzip(p)
            fmt.Println(string(res))
            resMap := common.JsonDecodeByte(res)
            if  v, ok := resMap["ping"];ok  {
                pingMap := make(map[string]interface{})
                pingMap["pong"] = v
                pingParams := common.JsonEncodeMapToByte(pingMap)
                if err := ws.WriteMessage(websocket.TextMessage, pingParams); err == nil {
                    reqMap := new(common.ReqStruct)
                    reqMap.Id = strconv.Itoa(time.Now().Nanosecond())
                    reqMap.Req = conf.LtcTopic.KLineTopicDesc
                    reqBytes , err := json.Marshal(reqMap)
                    if err!=nil {
                        continue
                    }
                    if err := ws.WriteMessage(websocket.TextMessage,reqBytes); err == nil {
                    }else{
                        fmt.Errorf("send req response error %s",err.Error())
                    }
                }else{
                    fmt.Errorf("huobi server ping client error %s",err.Error())
                    continue
                }
            }
            if  _, ok := resMap["rep"];ok  {
                var resStruct common.ResStruct
                json.Unmarshal(res,&resStruct)
                //resStruct.Status
                fmt.Println(resStruct)
            }
        }
    }
}