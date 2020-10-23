package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 处理url，得到返回结果。
func Fetch(url string)([]byte,error)  {
	req,_ := http.NewRequest("GET",url,nil)
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36")
	resp,err :=(&http.Client{}).Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
