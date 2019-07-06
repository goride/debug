# goride/debug

> Use the browser to golang display debugging information

> 使用浏览器显示golang 调试信息 

<img width="40" src="https://avatars0.githubusercontent.com/u/52598299?s=200&v=4" >   

<img width="1274" alt="1" src="https://user-images.githubusercontent.com/3949015/60758759-912bbf00-a04d-11e9-902c-148ac38ec27d.png" />

## install

`go get github.com/goride/debug`


## use 

use `Debug(v interfave{})` and open chrome [http://127.0.0.1:9999/](http://127.0.0.1:9999/)

```go
type Author struct {
    Name string
    Age int
}
nimo := Author{
    Name: "nimo",
    Age: 27,
}
Debug(nimo)
```