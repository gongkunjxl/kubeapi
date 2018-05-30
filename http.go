package main

import (
	//	"crypto/tls"
	//	//	"encoding/json"
	"fmt"

	"github.com/mikemintang/go-curl"
	//	"io"
	//	//	"io/ioutil"
	//	"log"
	//	"net/http"
	//	//	"strings"
)

func main() {
	url := "https://hawkular-metrics.master.example.com/hawkular/metrics/metrics"
	headers := map[string]string{
		"Accept":          "application/json",
		"Content-Type":    "application/json",
		"Hawkular-Tenant": "default",
		"Authorization":   "Bearer Ej2p1ZAqgx72oSsGHU3RIikXDoN3UalYfIpZsCNpsZ0",
	}
	req := curl.NewRequest()
	resp, err := req.
		SetUrl(url).
		SetHeaders(headers).Get()

	if err != nil {
		fmt.Println(err)
	} else {
		if resp.IsOk() {
			fmt.Println(resp.Body)
		} else {
			fmt.Println(resp.Raw)
		}
	}

}

/*



















 */
