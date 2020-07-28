package grpc

//定义一个服务结构体
type HelloService struct {
}

/**
*大写的Hello表式 必须要大写，因为此方法要公开，不能小写
*定义远程方法Hello
 */
func (p *HelloService) Hello(request string, response *string) error {
	*response = "hello:" + request
	return nil
}
