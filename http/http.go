package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func NewViperFromFile(path string) (*viper.Viper, error) {
	vp := viper.New()
	vp.SetConfigFile(path)
	if err := vp.ReadInConfig(); err != nil {
		return vp, err
	}
	return vp, nil
}


func DoHttpRequest(configPath string, requestName string) error{
	vp, err := NewViperFromFile(configPath)
	if err != nil {
		return fmt.Errorf("NewViperFromFile: %v", err)
	}
	var param RequestSchema
	if err = vp.Sub(requestName).Unmarshal(&param); err != nil {
		return fmt.Errorf("UnmarshalError: %v", err)
	}
	if len(param.Headers)==0{
		param.Headers = make(map[string][]string)
	}
	var req *http.Request
	switch param.Method {
	case "GET":
		getUrl := param.Url
		if len(param.Values) > 0 {
			v := url.Values(param.Values)
			getUrl += "?" + v.Encode()
		}
		req, err = http.NewRequest(param.Method, getUrl, nil)
		req.Header = param.Headers
	case "POST":
		var bodyIO io.Reader
		if len(param.FileToUpload) >0{
			if data, err := os.ReadFile(param.FileToUpload);err!=nil{
				return fmt.Errorf("ReadFile: %v", err)
			}else{
				bodyIO = bytes.NewBuffer(data)
			}
		}else{
			switch param.ContentType{
			case "application/json":
				param.Headers["Content-Type"] = []string{"application/json"}
				if data, err := json.Marshal(param.Body);err!=nil{
					return err
				}else{
					bodyIO = bytes.NewBuffer(data)
				}
			default:
				bodyIO = strings.NewReader(fmt.Sprintf("%v", param.Body))
			}
		}
		req, err = http.NewRequest(param.Method, param.Url, bodyIO)
		req.Header = param.Headers
	default:
		return fmt.Errorf("unknown method: %s", param.Method)
	}
	FPrintln("%v", req)
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	rspData, err := io.ReadAll(rsp.Body)
	if err!=nil{
		return err
	}
	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status code: %d, body=%s", rsp.StatusCode, string(rspData))
	}
	defer rsp.Body.Close()
	if len(param.OutputFile)>0{
		return os.WriteFile(param.OutputFile, rspData, os.ModePerm)
	}else{
		FPrintln("%s", string(rspData))
	}
	return err
}