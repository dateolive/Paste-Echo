```bash
├── app 
│   ├── controller    # 控制器
│   ├── httpio       
│   │   ├── in        # 请求结构体定义  
│   │   └── out       # 响应结构体定义  
│   ├── middleware    # 自定义中间件， 错误处理程序 
│   ├── model         # 持久层   
│   ├── repo          # db层封装
│   └── router        # 路由定义
├── cmd
│   └── server        # 启动程序
├── config
├── model            # model 模块化  
│   ├── webserver     # http服务
│   └── wsserver      # websocket服务  
└── pkg               # 自定义的扩展包 
    ├── atomicbool    # 原子bool
    ├── modules       # 模块包
    └── validator     # 验证包
 ```