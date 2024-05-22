package middleware

import (
	"bytes"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

//全局中间件

//功能:记录所有请求响应信息到自己的小本本上

// rest.Middleware 本质: Middleware func(next http.HandlerFunc) http.HandlerFunc
// type HandlerFunc func(ResponseWriter, *Request)

// bodyCopy 满足http.ResponseWriter接口类型
type bodyCopy struct {
	http.ResponseWriter               //结构体嵌入建构提
	body                *bytes.Buffer //我的小本本，记录响应体内容
}

// NewBodyCopy 初始化方法
func NewBodyCopy(w http.ResponseWriter) *bodyCopy {
	return &bodyCopy{
		ResponseWriter: w,
		body:           bytes.NewBuffer([]byte{}),
	}
}

func (bc bodyCopy) Write(b []byte) (int, error) {
	//1. 先在我的小本本上记录响应内容
	bc.body.Write(b)
	// 2.再往合同谈判响应里面写内容
	return bc.ResponseWriter.Write(b)
}

// CopyResp
func CopyResp(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//处理请求前

		//初始化得到一个自定义的 http.ResponseWriter：
		bc := NewBodyCopy(w)
		next(bc, r) //实际的路由处理handler函数
		//处理请求后
		logx.Infof("-->req:%v resp:%v\n", r.URL.RawPath, bc.body.String())

		//fmt.Printf("-->req:%v resp:%v\n", r.URL.RawPath, bc.body.String())
	}

}

//调用其它服务的中间件(开关 例子)
// func IsOpen(ok bool) rest.Middleware {
// 	return func(next http.HandlerFunc) http.HandlerFunc {
// 		return func(w http.ResponseWriter, r *http.Request) {
// 			//实现逻辑
// 			if !ok{
// 				fmt.Println("! ok")
// 			}
// 			next(w, r)
// 		}

// 	}
// }
