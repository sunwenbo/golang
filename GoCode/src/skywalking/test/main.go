type Req struct {
    A       int
    Header  string        // 添加此字段，用于传递context信息
}

func (p *Req) Set(key, value string) error {
    p.Header = fmt.Sprintf("%s:%s", key, value)
    return nil
}

r, err = reporter.NewGRPCReporter("192.168.204.130:11800")
if err != nil {
    logs.Info("[New GRPC Reporter Error]: [%v]", err)
    return
}
 
// 这个程序中所有的span都会跟服务名叫RTS_Test的服务关联起来
tracer, err = go2sky.NewTracer("RTS_Test", go2sky.WithReporter(r), go2sky.WithInstance("RTS_Test_1"))
if err != nil {
    logs.Info("[New Tracer Error]: [%v]", err)
    return
}
tracer.WaitUntilRegister()

func OnSnapshot() {
    // client := GetClinet()
    // 表示收到客户端请求，因为只追踪后台服务之间的链路，所以这里不需要提取context信息
    span2, ctx, err := tracer.CreateEntrySpan(context.Background(), "/API/Snapshot", func() (string, error){
        return "", nil
    })
    if err != nil {
        logs.Info("[Create Exit Span Error]: [%v]", err)
        return
    }
    span2.SetComponent(5200)
	
    // 表示rpc调用的span，这里需要向上游服务注入context信息，即参数中的header
    req := Req{3, ""}
    span1, err := tracer.CreateExitSpan(ctx, "/Service/OnSnapshot", "RTS_Server", func(header string) error{
        return req.Set(propagation.Header, header)
    })
    if err != nil {
        logs.Info("[Create Exit Span Error]: [%v]", err)
        return
    }
    span1.SetComponent(5200)    // Golang程序使用范围是[5000, 6000)，还要在skywalking中配置，config目录下的component-libraries.yml文件
 
    var res Res
    // rpc调用
    err = conn.Call("Req.Snapshot", req, &res)
    if err != nil {
        logs.Info("[RPC Call Snapshot Error]: [%v]", err)
        return
    } else {
        logs.Info("[RPC Call Snapshot Success]: [%s]", res)
    }
 
    span1.End()
    span2.End()    // 一定要确保span被结束
 
    // s1 := ReportedSpan(span1)
    // s2 := ReportedSpan(span2)
    // spans := []go2sky.ReportedSpan{s1, s2}
    // r.Send(spans)
}

type ReqBody struct {
    A       int
    Header  string
}
func (p *ReqBody) Get(key string) string {
    subs := strings.Split(p.Header, ":")
    if len(subs) != 2 || subs[0] != key {
	    return ""
    }
 
    return subs[1]
}
r, err = reporter.NewGRPCReporter("192.168.204.130:11800")
if err != nil {
    logs.Info("[New GRPC Reporter Error]: [%v]\n", err)
    return
}
 
tracer, err = go2sky.NewTracer("Service_Test", go2sky.WithReporter(r), go2sky.WithInstance("Service_Test_1"))
if err != nil {
    logs.Info("[New Tracer Error]: [%v]\n", err)
    return
}
tracer.WaitUntilRegister()
func (p *Req)Snapshot(req ReqBody, res *Res) error {
    // 表示收到 rpc 客户端的请求，这里需要提取context信息
    span1, ctx, err := tracer.CreateEntrySpan(context.Background(), "/Service/OnSnapshot/QueringSnapshot", func() (string, error){
	return req.Get(propagation.Header), nil
    })
    if err != nil {
	    logs.Info("[Create Exit Span Error]: [%v]\n", err)
	    return err
    }
    span1.SetComponent(5200)
    // span1.SetPeer("Service_Test")
 
    // 表示去请求了一次数据库
    span2, err := tracer.CreateExitSpan(ctx, "/database/QuerySnapshot", "APIService", func(header string) error {
	    return nil
    })
    span2.SetComponent(5200)
 
    time.Sleep(time.Millisecond * 6)
    *res = "Return Snapshot Info"
 
    span2.End()
    span1.End()
 
    // s1 := ReportedSpan(span1)
    // s2 := ReportedSpan(span2)
    // spans := []go2sky.ReportedSpan{s1, s2}
    // r.Send(spans)
 
    return nil
}
