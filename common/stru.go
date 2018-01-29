package common

type ReqStruct struct {
    Req string `json:"req"`
    Id  string `json:"id"`
}
type DataStruct struct {
    Amonut float64 `json:"amonut"`
    Count  int64   `json:"count"`
    Id     int64   `json:"id"`
    Open   float64 `json:"open"`
    Close  float64 `json:"close"`
    Low    float64 `json:"low"`
    High   float64 `json:"high"`
    Vol    float64 `json:"vol"`
}
type ResStruct struct {
    Status   string       `json:"status"`
    Rep      string       `json:"rep"`
    Subbed   string       `json:"subbed"`
    Unsubbed string       `json:"unsubbed"`
    Data     []DataStruct `json:"data"`
    Tick     DataStruct   `json:"tick"`
    ErrCode  string       `json:"err-code"`
    ErrMsg   string       `json:"err-msg"`
    Ts       int64        `json:"ts"`
    Ch       string       `json:"ch"`
}
type SubStruct struct {
    Sub   string `json:"sub"`
    Id    string `json:"id"`
    Unsub string `json:"unsub"`
}