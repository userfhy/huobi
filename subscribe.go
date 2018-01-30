package main

import (
    "github.com/gorilla/websocket"
    ui "github.com/gizak/termui"
    // "math"
    "os"
    "os/signal"
    "syscall"
    "fmt"
    "time"
    "strconv"
    "encoding/json"
    "./conf"
    "./common"
)

type urlWidgets struct {
    par     ui.Par
    lineChart *ui.LineChart
}

func buildUI(sinps []float64, Text string) map[string]*urlWidgets{
    if err := ui.Init(); err != nil {
        panic(err)
    }
    // defer ui.Close()

    widgets := make(map[string]*urlWidgets)

    par1 := ui.NewPar(Text)
    par1.Height = 5
    par1.BorderLabel = "标签"

    lc := ui.NewLineChart()
    lc.BorderLabel = "braille-mode Line Chart 走势"
    lc.Data = sinps
    lc.Height = 11
    lc.AxesColor = ui.ColorWhite
    lc.LineColor = ui.ColorYellow | ui.AttrBold

    widgets["test"] = &urlWidgets{
        par: *par1,
        lineChart: lc,
    }

    // build layout
    ui.Body.AddRows(
         ui.NewRow(
            ui.NewCol(6, 0, lc),
            ui.NewCol(6, 0, par1)))

    // calculate layout
    ui.Body.Align()

    ui.Render(ui.Body)

/*    ui.Handle("/timer/1s", func(e ui.Event) {
        t := e.Data.(ui.EvtTimer)
        i := t.Count
        if i > 103 {
            ui.StopLoop()
            return
        }

        lc.Data = sinps[2*i:]
        ui.Render(ui.Body)
    })*/

    // ui.Loop()

    return widgets

    // os.Exit(0)
}

func ToStr(i interface{}) string {
  return fmt.Sprintf("%v", i)
}

func main()  {
    go func() {
        //创建监听退出chan
        c := make(chan os.Signal)
        //监听指定信号 ctrl+c kill
        signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2, os.Interrupt)
        //阻塞直到有信号传入
        fmt.Println("启动")
        s := <-c
        fmt.Println("退出信号", s)
        os.Exit(0)
    }()

    dialer := websocket.Dialer{ /* set fields as needed */ }
    ws, _, err := dialer.Dial("wss://api.huobi.pro/ws", nil)
    if err != nil {
        // handle error
    }
    subStruct := new(common.SubStruct)
    subStruct.Id = strconv.Itoa(time.Now().Nanosecond())
    subStruct.Sub = conf.LtcTopic.KLineTopicDesc
    reqBytes , err := json.Marshal(subStruct)
    fmt.Println(subStruct)

    if err!=nil {
        return
    }
    if err := ws.WriteMessage(websocket.TextMessage, reqBytes); err == nil {

    }else{
        fmt.Errorf("send req response error %s", err.Error())
    }

    //使用make 创建 切片
    var slice1 []float64

    for {
        if _,p,err := ws.ReadMessage();err == nil {
            res := common.UnGzip(p)
            // fmt.Println(string(res))
            resMap := common.JsonDecodeByte(res)
            if  v, ok := resMap["ping"];ok  {
                pingMap := make(map[string]interface{})
                pingMap["pong"] = v
                pingParams := common.JsonEncodeMapToByte(pingMap)
                if err := ws.WriteMessage(websocket.TextMessage, pingParams); err == nil {

                }else{
                    fmt.Errorf("huobi server ping client error %s",err.Error())
                    continue
                }
            }

            if _, ok := resMap["ch"];ok  {
                var resStruct common.ResStruct
                json.Unmarshal(res, &resStruct)

                // 使用 append 添加元素，并且未超出 cap
                slice1 = append(slice1, resStruct.Tick.Close)

                showUi(slice1, resStruct)
            }
        }
    }
}

func showUi(sinps []float64, resStruct common.ResStruct) {
/*    sinps := (func() []float64 {
        n := 400
        ps := make([]float64, n)
        for i := range ps {
            ps[i] = 1 + math.Sin(float64(i)/5)
        }
        return ps
    })()*/

/*    showstr := fmt.Sprintf("当前价：%s\n最低价：%s\n最高价：%s\n",
        strconv.FormatFloat(resStruct.Tick.Close, 'f', 4, 64),
        strconv.FormatFloat(resStruct.Tick.Low, 'f', 4, 64),
        strconv.FormatFloat(resStruct.Tick.High, 'f', 4, 64))*/

    showstr := fmt.Sprintf("[%s](fg-green)\n当前价：%s",
    resStruct.Ch,
    strconv.FormatFloat(resStruct.Tick.Close, 'f', 4, 64))

/*fmt.Println(sinps)
fmt.Println(showstr)*/
    buildUI(sinps, showstr)

    ui.Handle("/sys/kbd/q", func(ui.Event) {
        ui.StopLoop()
        os.Exit(0)
    })

    ui.Handle("/sys/kbd/C-c", func(ui.Event) {
        ui.StopLoop()
        os.Exit(0)
    })

    ui.Handle("/sys/wnd/resize", func(e ui.Event) {
        ui.Body.Width = ui.TermWidth()
        ui.Body.Align()
        ui.Clear()
        ui.Render(ui.Body)
    })
}
