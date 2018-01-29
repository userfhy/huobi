package common

import (
    "bytes"
    "compress/gzip"
    "io/ioutil"
    "encoding/json"
)

func Gzip(data []byte) {
    var b bytes.Buffer
    w := gzip.NewWriter(&b)
    defer w.Close()
    w.Write(data)
    w.Flush()
    //fmt.Println("gzip size:", len(b.Bytes()))
}

func UnGzip(byte []byte) []byte {
    b := bytes.NewBuffer(byte)
    r, _ := gzip.NewReader(b)
    defer r.Close()
    undatas, _ := ioutil.ReadAll(r)
    //fmt.Println("ungzip size:", len(undatas))
    return undatas
}

func JsonDecodeString(String string) map[string]interface{} {
    jsonMap := make(map[string]interface{})
    json.Unmarshal([]byte(String), &jsonMap)
    return jsonMap
}

func JsonDecodeByte(bytes []byte) map[string]interface{} {
    jsonMap := make(map[string]interface{})
    json.Unmarshal(bytes, &jsonMap)
    return jsonMap
}
func JsonEncodeMapToByte(stringMap map[string]interface{}) []byte {
    jsonBytes, err := json.Marshal(stringMap)
    if err != nil {
        return nil
    }
    return jsonBytes
}
