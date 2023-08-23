# 白泽 - golang项目包依赖可视化工具😼

## 前言
白泽可以直观的展示golang项目的包依赖之间的关系，采用了可视化的方式展示依赖关系之间的拓扑图。

[Demo](https://baize.pjmcode.top/#/)

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
- [ ] AboutBaize

#### 视图
- [ ] 首页
- [ ] 关于白泽

### 基础工具
#### Makefile
- [x] 打包gh-pages
- [ ] 编译二进制并打包vue静态文件