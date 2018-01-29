package conf

import (
    "fmt"
)

const (
    //$symbol 为币种，可选值： { ethbtc, ltcbtc, etcbtc, bccbtc...... }
    LTCUSDT = "iostusdt"
)

const (
    //$period 可选值：{ 1min, 5min, 15min, 30min, 60min, 1day, 1mon, 1week, 1year }
    PERIOD1min  = "1min"
    PERIOD5min  = "5min"
    PERIOD15min = "15min"
    PERIOD30min = "30min"
    PERIOD60min = "60min"
    PERIOD1day  = "1day"
    PERIOD1mon  = "1mon"
    PERIOD1week = "1week"
    PERIOD1year = "1year"
)
const (
    //$type 可选值：{ step0, step1, step2, step3, step4, step5 } （合并深度0-5）；step0时，不合并深度
    TYPE0 = "step0"
    TYPE1 = "step1"
    TYPE2 = "step2"
    TYPE3 = "step3"
    TYPE4 = "step4"
    TYPE5 = "step5"
)
const (
    KLINE        = iota
    MARKETDEPTH  = iota + 1
    TRADEDETAIL  = iota + 2
    MARKETDETAIL = iota + 3
)
const (
    KLine        = "market.%s.kline.%s"     //$symbol $period
    MarketDepth  = "market.%s.depth.%s"     //$symbol $type
    TradeDetail  = "market.%s.trade.detail" //$symbol
    MarketDetail = "market.%s.detail"       //$symbol
)

type Topic struct {
    KLineTopicDesc        string
    MarketDepthTopicDesc  string
    TradeDetailTopicDesc  string
    MarketDetailTopicDesc string
}
func (t *Topic) Build(symbol string) {
    t.KLineTopicDesc = fmt.Sprintf(KLine, symbol,PERIOD1min)
    t.MarketDepthTopicDesc = fmt.Sprintf(MarketDepth, symbol,TYPE0)
    t.TradeDetailTopicDesc = fmt.Sprintf(TradeDetail, symbol)
    t.MarketDetailTopicDesc = fmt.Sprintf(MarketDetail, symbol)
}

var LtcTopic = new(Topic)

func init() {
    LtcTopic.Build(LTCUSDT)
}
