## API 文档
### 1. "获取go.mod信息"

1. route definition

- Url: /baize/v1/local/module
- Method: GET
- Request: `CommonDirReq`
- Response: `GetModuleInfoResp`

2. request definition



```golang
type CommonDirReq struct {
	Dir string `form:"dir"`
}
```


3. response definition



```golang
type GetModuleInfoResp struct {
	ModulePath string `json:"modulePath"`
	ModuleName string `json:"moduleName"`
	ModuleVersion string `json:"moduleVersion"`
}
```

### 2. "获取包列表"

1. route definition

- Url: /baize/v1/local/packages
- Method: GET
- Request: `CommonDirReq`
- Response: `GetPackagesResp`

2. request definition



```golang
type CommonDirReq struct {
	Dir string `form:"dir"`
}
```


3. response definition



```golang
type GetPackagesResp struct {
	OtherPkgList []*OtherPkgItem `json:"otherPkgList"`
	OtherPkgCount uint `json:"otherPkgCount"`
	ProjectPkgList []string `json:"projectPkgList"`
	ProjectPkgCount uint `json:"projectPkgCount"`
}
```

### 3. "获取go文件列表"

1. route definition

- Url: /baize/v1/local/file
- Method: GET
- Request: `CommonDirReq`
- Response: `GetGoFilesResp`

2. request definition



```golang
type CommonDirReq struct {
	Dir string `form:"dir"`
}
```


3. response definition



```golang
type GetGoFilesResp struct {
	GoFileList []*GoFileItem `json:"goFileList"`
	GoFileCount uint `json:"goFileCount"`
}
```

### 4. "生成有向依赖图"

1. route definition

- Url: /baize/v1/local/graph
- Method: GET
- Request: `CommonDirReq`
- Response: `GenerateGraphResp`

2. request definition



```golang
type CommonDirReq struct {
	Dir string `form:"dir"`
}
```


3. response definition



```golang
type GenerateGraphResp struct {
	Nodes []*GraphNode `json:"nodes"`
	Edges []*GraphEdge `json:"edges"`
}
```

### 5. "ping"

1. route definition

- Url: /baize/v1/ping
- Method: GET
- Request: `-`
- Response: `-`

2. request definition



3. response definition


