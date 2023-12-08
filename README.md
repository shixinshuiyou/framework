# framework

| 仓库包        | 主要功能                                           | 其他  |
|------------|------------------------------------------------|-----|
| export     | 导出文件功能，支持导出excel 文件                            | 已完成 |
| kafka      | 支持kafka 的消费和生产,支持批量生产，批量消费待支持                  |     |
| netx       | 网络包，支持 s3 相关操作，http 相关操作，邮件相关操作等               | 建设中 |
| profile    | 性能检测工具                                         |     |
| threading  | 协程组件包，来源 go-zero                               | 已完成 |
| timex      | 支持日期维度，或者标准时间维度的类型                             | 已完成 |
| signal     | 信号处理模块,将服务的所有信号统一处理                            | 已完成 |
| redisx     | redis 查询                                       |     |
| log        | 基于 logrus 开发的日志组件，支持上传 es，支持gorm 日志采集等         | 已完成 |
| db         | db 库，提供了MySQL，Mongo 的链接封装，并提供了MySQL的通用Model    |     |
| config     | 支持 ini/yaml 等配置文件加载                            |     |
| collection | 集合操作：支持先进先出队列，环形队列，滚动窗口，集合，时间片轮转。来源 go-zero 仓库 |     |
| web        | 基于gin的web服务                                    | 已完成 |
| trace      | 基于jaeger的链路追踪服务                                | 已完成 |
| util/api   | 基于配置自动注入外部接口（支持GET/POST）请求                     | 已完成 |
