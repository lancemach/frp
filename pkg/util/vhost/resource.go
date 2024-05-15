// Copyright 2017 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vhost

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/pkg/util/version"
)

var NotFoundPagePath = ""

const (
	NotFound = `<!DOCTYPE html>
<html>
<head>
<title>Not Found</title>
<style>
body{margin:0;padding:0;font-size:12px;line-height:22px;font-family:tahoma, arial, 'Hiragino Sans GB', "Microsoft YaHei", "黑体", sans-serif;-webkit-text-size-adjust:none;}
html,body,div,dl,dt,dd,ul,ol,li,h1,h2,h3,h4,h5,h6,pre,form,fieldset,input,textarea,p,blockquote,th,td,p{margin:0;padding:0;}
body,html,.wo-404-wrap{ width: 100%; height: 100% }
div{webkit-box-sizing: border-box; -moz-box-sizing: border-box;box-sizing: border-box;}
.wo-404-wrap{ background:url('https://images-resource-1256738796.cos.ap-chengdu.myqcloud.com/public/404.jpg') no-repeat 50% 58% ;}
.wo-404-wrap .text{padding:8% 20% 30px;text-align: center; width: 100%}
.wo-404-wrap .text p{font-size: 14px;}
.wo-404-wrap .text p a,.wo-404-wrap .text p em{color: #269DFF;margin: 0 4px;}
.wo-404-wrap .text p em{font-style: normal;color: #E40008}
.wo-404-wrap .text h3{font-size: 16px;font-weight: 400;color: #808080}
@media (max-width: 970px) {
.wo-404-wrap{ background-size:100%; }.wo-404-wrap .text{padding:18% 8%}
}
</style>
</head>
<body>
<div class="wo-404-wrap">
<div class="text">
<h3>对不起！亲，您要访问的页面连接错误或者不存在，我们正在努力修复!</h3>
</div>
</div>
</body>
</html>
`
)

func getNotFoundPageContent() []byte {
	var (
		buf []byte
		err error
	)
	if NotFoundPagePath != "" {
		buf, err = os.ReadFile(NotFoundPagePath)
		if err != nil {
			log.Warnf("read custom 404 page error: %v", err)
			buf = []byte(NotFound)
		}
	} else {
		buf = []byte(NotFound)
	}
	return buf
}

func NotFoundResponse() *http.Response {
	header := make(http.Header)
	header.Set("server", "frp/"+version.Full())
	header.Set("Content-Type", "text/html")

	content := getNotFoundPageContent()
	res := &http.Response{
		Status:        "Not Found",
		StatusCode:    404,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Header:        header,
		Body:          io.NopCloser(bytes.NewReader(content)),
		ContentLength: int64(len(content)),
	}
	return res
}
