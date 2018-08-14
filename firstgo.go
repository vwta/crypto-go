package main

import (
"crypto/tls"
"crypto/x509"
"fmt"
"io/ioutil"
"log"
"net/http"
	"encoding/json"
)
type cfmargin struct {
	Contracts float64 `json:"contracts"`
	InitialMargin float64 `json:"initialMargin"`
	MaintenanceMargin float64 `json:"maintenanceMargin"`
}
type cfinstrument struct {
	Symbol string `json:"symbol"`
	Type string `json:"type"`
	Underlying string `json:"underlying"`
	LastTradingTime string `json:"lastTradingTime"`
	TickSize float64 `json:"tickSize"`
	contractSize float64 `json:"contractSize"`
	tradeable bool `json:"tradeable"`
	marginLevels []cfmargin `json:"marginLevels"`
}
type response2 struct {
	Result   string      `json:"result"`
	Instruments []cfinstrument `json:"instruments"`
}

func main() {
	fmt.Printf("go start")
	caCert, err := ioutil.ReadFile("server.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	cert, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := client.Get("https://www.cryptofacilities.com/derivatives/api/v3/instruments")
	if err != nil {
		log.Println(err)
		return
	}

	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("%v\n", resp.Status)

	str := string(htmlData)
	//fmt.Printf(str)
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	//fmt.Println(res)
	fmt.Println(res.Result)

	//fmt.Println(res.Instruments)

	for i := 0; i < len( res.Instruments); i++ {
		fmt.Println(res.Instruments[i])
	}


}
