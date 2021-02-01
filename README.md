## Telegram Bot

Telegram Bot 是一言在 Telegram 提供的 Bot 组件。

## 鸣谢

项目的诞生与发展离不开 **萌创团队** 以及 **一言项目组** 的支持，更离不开  [JetBrains](https://www.jetbrains.com/?from=hitokoto-osc) 为开源项目免费提供具有强生产力的 IDE 等相关授权。
[<img src=".github/jetbrains-variant-3.png" width="200"/>](https://www.jetbrains.com/?from=hitokoto-osc)

## 许可证

项目代码遵循 **GNU General Public License v3.0** 许可。
此外，项目不会通过任何途径 **签发** 或 **授权** 商用行为（commercial use）。

## 开发

现在，让我们简单介绍下怎样参与咱们的开发。

### 依赖

-   框架（重要的外部依赖）
    -   配置：viper
    -   日志：logrus
    -   flag 解析：pflag
    -   CI/CD：Github Action（后期前端内容的继承也将通过此服务）

### 初始开发环境

1.  自行安装 Go SDK
2.  克隆项目至 `$GOROOT/src/github.com/hitokoto-osc/Moe`
3.  安装 node 依赖。

```shell
$ yarn
```

### 编译

```shell
$ make build
```

### 测试

```shell
$ make test
```
