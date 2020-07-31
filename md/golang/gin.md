# 将 request body 绑定到不同的结构体中
```go
func SomeHandler(c *gin.Context) {
  objA := formA{}
  objB := formB{}
  // 读取 c.Request.Body 并将结果存入上下文。
  if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
    c.String(http.StatusOK, `the body should be formA`)
  // 这时, 复用存储在上下文中的 body。
  } else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
    c.String(http.StatusOK, `the body should be formB JSON`)
  // 可以接受其他格式
  } else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
    c.String(http.StatusOK, `the body should be formB XML`)
  } else {
    ...
  }
}
```
