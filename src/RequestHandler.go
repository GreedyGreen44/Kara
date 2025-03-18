package main

import (
	"errors"
	"io"
	"net/http"
	"strconv"
)

func handleRequest(configMap map[string]string) (data []byte, resultCode int) {
	req, resultCode := createRequest(configMap)
	if resultCode != 0 {
		return nil, resultCode
	}

	if req == nil {
		return nil, handleError([2]byte{0x01, 0x03}, errors.New("nil request created"))
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, handleError([2]byte{0x00, 0x06}, err)
	}
	defer res.Body.Close()
	resultCode = checkResponse(res)
	if resultCode != 0 {
		return nil, resultCode
	}
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, handleError([2]byte{0x00, 0x02}, err)
	}
	return data, resultCode
}

func createRequest(configMap map[string]string) (req *http.Request, resultCode int) {
	url := "https://" + configMap["Host"] + "/api/pub/mp2?bbox="
	zoom, err := strconv.ParseInt(configMap["Zoom"], 10, 32)
	if err != nil {
		return nil, handleError([2]byte{0x00, 0x03}, err)
	}
	boxBot, resultCode := calculateIntCoordinate(configMap["BoxBot"], int(zoom))
	if resultCode != 0 {
		return nil, resultCode
	}
	boxTop, resultCode := calculateIntCoordinate(configMap["BoxTop"], int(zoom))
	if resultCode != 0 {
		return nil, resultCode
	}
	boxLeft, resultCode := calculateIntCoordinate(configMap["BoxLeft"], int(zoom))
	if resultCode != 0 {
		return nil, resultCode
	}
	boxRight, resultCode := calculateIntCoordinate(configMap["BoxRight"], int(zoom))
	if resultCode != 0 {
		return nil, resultCode
	}
	url += boxLeft + "%2C" + boxBot + "%2C" + boxRight + "%2C" + boxTop + "&zoom=" + configMap["Zoom"]
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, handleError([2]byte{0x01, 0x03}, err)
	}
	req.Header.Add("Host", configMap["Host"])
	if userAgent, ok := configMap["UserAgent"]; ok {
		req.Header.Add("User-Agent", userAgent)
	} else {
		req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Referer", configMap["Host"])
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")

	return req, 0
}

func checkResponse(res *http.Response) (resultCode int) {
	if res.StatusCode != 200 {
		return handleError([2]byte{0x00, 0x02}, errors.New("response is not successful"))
	}
	if res.Header.Get("Content-Type") != "application/octet-stream" {
		return handleError([2]byte{0x00, 0x02}, errors.New("not a octet-stream response"))
	}
	return 0
}
