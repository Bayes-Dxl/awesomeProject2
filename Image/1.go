package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// 下载图片信息
func downLoad(base string, url string) error {
	pic := base
	idx := strings.LastIndex(url, "/")
	if idx < 0 {
		pic += "/" + url
	} else {
		pic += url[idx+1:]
	}
	v, err := http.Get(url)
	if err != nil {
		fmt.Printf("Http get [%v] failed! %v", url, err)
		return err
	}
	defer v.Body.Close()
	content, err := ioutil.ReadAll(v.Body)
	if err != nil {
		fmt.Printf("Read http response failed! %v", err)
		return err
	}
	err = ioutil.WriteFile(pic+".jpg", content, 0666)
	if err != nil {
		fmt.Printf("Save to file failed! %v", err)
		return err
	}
	return nil
}

func main() {
	savePath := "./"
	//url := "https://avatar.csdnimg.cn/2/9/0/1_togolife.jpg"
	url1 := "https://hbimg.huaban.com/32f065b3afb3fb36b75a5cbc90051b1050e1e6b6e199-Ml6q9F_fw320"
	err := downLoad(savePath, url1)
	if err != nil {
		fmt.Println("Download pic file failed!", err)
	} else {
		fmt.Println("Download file success.")
	}
	return
}

//imgPath := "D:\\GitHub\\Goland\\awesomeProject2\\Image"
//imgUrl := "https://hbimg.huaban.com/32f065b3afb3fb36b75a5cbc90051b1050e1e6b6e199-Ml6q9F_fw320"
//
//
