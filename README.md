# Cloudgo-IO

使用go+negroni+mux框架构造的简易http静态服务器。

## 静态服务访问

通过`http.FileServer`，能够很方便的构造一个静态服务器，代码如下

    // 服务静态文件
    mx.PathPrefix("/").Handler(
    http.StripPrefix("/", http.FileServer(http.Dir("./static")))).
    Methods(http.MethodGet)

## 简单JS访问+提交表单

在主页面提交表单，前段JS通过ajax发送一个post请求

    $.post(
        "/",
        {
            name: $("input[name='name']").val(),
            birthday: $("input[name='birth']").val()
        },
        function(data, status) {
            if (status == "success") {
                $("tbody").append(
                    "<tr>" +
                    "  <td>" + data.name + "</td>" +
                    "  <td>" + data.birthday + "</td>" +
                    "</tr>"
                );
            }
        }
    );

服务器接收到请求，返回一个JSON数据

    req.ParseForm()
    data := struct {
    	  Name     string `json:"name"`
    	  Birthday string `json:"birthday"`
    }{req.FormValue("name"), req.FormValue("birthday")}
    formatter.JSON(w, http.StatusCreated, data)

前段js接收到数据，往表格中添加一行，即可看到新增加的行。

## 处理`/unknown`

对于没有实现的页面，应当返回一个501 not implemented，代码如下


    func notImplementedHandler() http.HandlerFunc {
    	 return func(w http.ResponseWriter, req *http.Request) {
    		  http.Error(w, "501 page not implemented", http.StatusNotImplemented)
    	 }
    }
