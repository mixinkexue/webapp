package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

type dataJSON struct {
	Key    string `json:"key"`
	Hash   string `json:"hash"`
	Fsize  int    `json:"fsize"`
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

type imageJSON struct {
	Info   string   `json:"info"`
	Status int      `json:"status"`
	Data   dataJSON `json:"data"`
}

var qnyUrl = "https://be-prod.redrock.cqupt.edu.cn/red-qny/upload/redrockoss/magipoke_intergral?key=%s"

func UploadFile(k string, filename string, content io.Reader) (string, error) {
	buf := new(bytes.Buffer)           // 实例化一个结构体
	writer := multipart.NewWriter(buf) // 返回一个writer指针

	formFile, err := writer.CreateFormFile("file", filename) // 提供表单中的字段名和文件名返回值是可写的接口io.Writer
	if err != nil {
		fmt.Println("Create form file failed:", err)
		return "", err
	}
	_, err = io.Copy(formFile, content)
	if err != nil {
		fmt.Println("Write to form file failed: ", err)
		return "", err
	}
	writer.Close() // 发送之前必须调用Close()以写入结尾行

	client := http.DefaultClient

	httpRequest, err := http.NewRequest("GET", fmt.Sprintf(qnyUrl, k), buf)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	httpRequest.Header.Add("Content-Type", writer.FormDataContentType())

	// 发送请求
	resp, err := client.Do(httpRequest)
	if err != nil {
		log.Println("fail:", err)
		return "", err
	}
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)
	var info imageJSON
	if err := json.Unmarshal(response, &info); err != nil {
		log.Println(string(response), err)
		return "", err
	}
	if info.Status != 10000 {
		fmt.Println(info)
		return "", errors.New("Upload error")
	}
	return info.Data.Name, nil
}
