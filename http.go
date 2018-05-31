package main

import (
	"encoding/json"
	"fmt"

	"github.com/mikemintang/go-curl"
)

func main() {
	//	url := "https://hawkular-metrics.master.example.com/hawkular/metrics/metrics"
	url := "https://hawkular-metrics.master.example.com/hawkular/metrics/counters/hello-pod%2F692777fb-6246-11e8-b642-fcaa14a963a6%2Fcpu%2Fusage/data"
	headers := map[string]string{
		"Accept":          "application/json",
		"Content-Type":    "application/json",
		"Hawkular-Tenant": "k8s-test",
		"Authorization":   "Bearer LFTsk5oUBHwQdPgTr0x5tU1BOzfOVB7R0fgX5Tz_HKA",
	}
	//the queries
	queries := map[string]string{
		"buckets": "5",
		"start":   "1527732390000",
	}
	req := curl.NewRequest()
	resp, err := req.
		SetUrl(url).
		SetHeaders(headers).
		SetQueries(queries).
		Get()

	if err != nil {
		fmt.Println(err)
	} else {
		if resp.IsOk() {
			//header and body
			for headername := range resp.Headers {
				fmt.Printf("%s: %s\n", headername, resp.Headers[headername])
			}
			//json body
			var dat []map[string]interface{}
			if err := json.Unmarshal([]byte(resp.Body), &dat); err == nil {
				for _, di := range dat {
					fmt.Printf("min: %v max: %v\n", di["min"], di["max"])
				}

			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(resp.Raw)
		}
	}

}

/*



















 */
