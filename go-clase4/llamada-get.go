package main

import (
	"fmt"
	"log"
	"net/http"
)

func requestIp() {
	const url = "https://httpbin.org/ip"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	type IpResponse struct {

	}

	var ip IpResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer el body")
	}
	json.Unmarshal(body, &ip)
	if resp.StatusCode == 200 http.StatusOk {
		fmt.Printf("Valor devuelto: %s", ip.MiIp)
	}
}

func main() {

}
