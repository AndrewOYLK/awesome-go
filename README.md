## 简述

此仓库用于记录日常开发中，自行设计的算法代码与编程技巧代码，方便日后快速查阅与回忆。每个库都编写了相应的测试用例，只需要到指定库目录下，执行 `go test -v` 即可看到运行效果

## 部分库介绍

### pkg/patchyaml

此递归代码用于合并两个yaml文件的内容，并最终输出合并后的结果，示例如下：

> 备注，sigs.k8s.io/yaml库也可实现此功能

source.yml文件：

```yaml
name: andrew
info:
  age: 28
  address: GZ
  goods:
  - python
  - golang
  - java
```

patch.yml文件：

```yaml
name: xiaoming
info:
  age: 18
  address: DongGuang
```

在pkg/patchyaml目录下执行go test -v，即得到如下输出文件：

finnal.yml文件：

```yaml
info:
  address: DongGuang
  age: 18
  goods:
  - python
  - golang
  - java
name: xiaoming
```