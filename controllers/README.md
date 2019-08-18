## 1. 获取参数
### 1.0 打印请求所有的信息 
``` 
this.Ctx.Request
```

### 1.1 获取头信息
```
this.Ctx.Request.Header
this.Ctx.Request.Host
this.Ctx.Request.Method
```

### 1.2 获取 url 上的参数【/:id】 
```
this.Ctx.Input.Param(":id")
```

### 1.3 获取 url 上的参数【?号后面的部分】
```
this.GetString("fields")
this.GetInt64("limit")
this.GetInt64("offset")
```

### 1.4 获取 form 表单数据
```
this.Ctx.Input.RequestBody
```