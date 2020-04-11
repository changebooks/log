# log
日志
==

<pre>
data := make(map[string]string)
data[log.ProfileDirectory] = "./log"
data[log.ProfileChannel] = "changebooks"
data[log.ProfileCallerShort] = "true"

profile, err := log.NewProfile(data)
if err != nil {
    fmt.Println(err)
    return
}

stream, err := log.NewStream(profile)
if err != nil {
    fmt.Println(err)
    return
}

logger, err := log.NewLogger(stream, "sample", 0)
if err != nil {
    fmt.Println(err)
    return
}

idRegister := &log.IdRegister{}
idRegister.SetTraceId("trace-id-10001")
idRegister.SetBizId("biz-id-20002")

logger.I("abc", "123", "test", idRegister)
logger.E("def", "123", "test", errors.New("456"), "", idRegister)

// 进程正常关闭前
errs := logger.Close()
fmt.Println(errs)
</pre>

<pre>
log.I("abc", "123", "456")
log.E("def", "123", "456")
</pre>
