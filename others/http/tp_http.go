package tphttp

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/spf13/viper"
)

func Post(targetUrl string, payload string) (*http.Response, error) {
	req, _ := http.NewRequest("POST", targetUrl, strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		logs.Info(err.Error())
	}
	return response, err
}

func PostWithDeviceInfo(targetUrl, payload, deviceId, accessToken string) (*http.Response, error) {

	timeout := time.Duration(viper.GetInt("data_transpond.http_timeout")) * time.Second
	if timeout == 0 {
		timeout = 3 * time.Second
	}

	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("POST", targetUrl, strings.NewReader(payload))
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("DeviceID", deviceId)
	req.Header.Add("AccessToken", accessToken)
	response, err := client.Do(req)
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}

	return response, nil
}

func Delete(targetUrl string, payload string) (*http.Response, error) {
	logs.Info("Delete:", targetUrl, payload)
	req, _ := http.NewRequest("DELETE", targetUrl, strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(req)
	if err == nil {
		logs.Info(response.Body)
	} else {
		logs.Info(err.Error())
	}
	return response, err
}

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		logs.Info("Response: ", string(body))
		return body, err
	} else {
		return nil, errors.New("Get failed with error: " + resp.Status)
	}
}
func PostJson(targetUrl string, payload []byte) (*http.Response, error) {
	req, _ := http.NewRequest("POST", targetUrl, bytes.NewReader(payload))
	req.Header.Add("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(req)
	return response, err
}

func SendSignedRequest(url, message, secret string) error {
	signature := GenerateHMAC(message, secret)

	// Creating the request
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(message))
	if err != nil {
		return fmt.Errorf("创建请求失败: %v", err)
	}

	// Adding the signature to the request header
	req.Header.Set("X-Signature-256", "sha256="+signature)
	req.Header.Set("Content-Type", "application/json")

	// Sending the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("请求已发送。状态码: %d\n", resp.StatusCode)
	return nil
}

func GenerateHMAC(message, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	signature := h.Sum(nil)
	return hex.EncodeToString(signature)
}
