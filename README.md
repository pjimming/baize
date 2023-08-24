# 白泽 - golang项目包依赖可视化工具😼

## 前言
白泽可以直观的展示golang项目的包依赖之间的关系，采用了可视化的方式展示依赖关系之间的拓扑图。

- [点击跳转至Demo](https://baize.pjmcode.top/#/)
- [点击下载可执行二进制](https://github.com/pjimming/baize/releases/download/v1.0.0/baize)

## 技术栈
- 后端：
  - 编程语言：[golang](https://golang.google.cn/)
  - 框架：[go-zero](https://go-zero.dev/)
- 前端：
  - JS框架：[Vue.js](https://cn.vuejs.org/)
  - UI库：
    - [bootstrap5](https://getbootstrap.com/)
    - [element-plus](https://element-plus.gitee.io/zh-CN/)
  - 图可视化引擎：[AntV G6](https://g6.antv.antgroup.com/)
- 数据库：不需要
- 打包工具：
  - 前端静态文件：[statik](https://github.com/rakyll/statik)
  - 配置文件：[go-bindata](https://github.com/go-bindata/go-bindata)

## 项目架构
### 后端架构
#### 通用组件
- [x] HttpResponse
- [x] CodeError

#### API接口
- [x] Ping
- [x] GetModuleInfo
- [x] GetPackages
- [x] GetGoFiles
- [x] GenerateGraph
  
### 前端架构
#### 组件
- [x] NavBar
- [x] ContentBase
- [x] Input
- [x] ListInfo
- [x] ProjectInfo
- [x] Error
- [x] Graph
- [x] 404
- [x] AboutBaize

#### 视图
- [x] 首页
- [x] 关于白泽

### 基础工具
#### Makefile
- [x] 打包gh-pages
- [x] 编译二进制并打包vue静态文件