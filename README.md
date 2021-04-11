
## 服务需要微服务，但代码需要高内聚。

文件夹作用规范

```
├─assets                                1.资源目录
│  ├─apis                                   1.1接口定义目录             
│  │  ├─appm                                    1.1.1项目1
│  │  └─userm                                   1.1.2项目2
│  ├─etcs                                   1.2配置目录
│  │  ├─appm                                    1.2.1项目1
│  │  └─userm                                   1.2.2项目2    
│  └─goctl                                  1.3goctl生成模板目录
│      └─api                                    1.3.x
├─common                                2.公共目录
│  ├─configz                                2.1统一配置处理
│  ├─errorz                                 2.2统一错误处理
│  ├─middlez                                2.3统一中间件
│  ├─respz                                  2.4统一返回值处理
│  └─toolsz                                 2.5统一工具
└─service                               3.服务目录
    ├─appm                                  3.1项目1目录
    │  ├─crm                                    3.1.1启动目录
    │  │  ├─api                                     1.1.1.1api入口
    │  │  │  └─internal
    │  │  │      ├─config
    │  │  │      ├─handler
    │  │  │      │  └─v1
    │  │  │      │      └─index
    │  │  │      ├─logic
    │  │  │      │  └─v1
    │  │  │      │      └─index
    │  │  │      ├─svc
    │  │  │      └─types
    │  │  ├─rmq                                     1.1.1.2队列脚本
    │  │  └─rpc                                     1.1.1.3rpc入口
    │  ├─model                                  3.1.2模型目录
    │  └─server                                 3.1.3服务目录
    └─userm                                 3.1项目2目录
```

#### ！！！注意

1. assets和脚本分开,因为:api需要交给前端,配置需要交给运维,goctl需要交给开发
2. etcs/common.yaml主要是为了保存全局唯一的配置,比如rpc配置就很适合.
3. app.api和app-api.api关系也是公共和部分的关系，如果是公共的放入app.api中
4. errorz/errors为了统一的错误处理,保证错误码在全局项目中唯一
5. respz/resphandler主要是统一错误处理,和统一的返回值


