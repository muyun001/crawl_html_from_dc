## 提供两个接口：发送任务到下载中心，从下载中心获取结果

#### SpiderRequest - config


|参数 | 作用 | 备注|
| --- | --- | ---|
| store_type | 存储类型（默认不修改） | 5 redis 操作 ； 默认5， 新版本一律使用5 |
| redirect | 是否跳转 | 0： 默认不跳转 1：跳转 |
| priority | 抓取优先级，只针对当前用户 |1（最不优先）2（默认）3（最优先） |
| post_data | post请求需要传输的数据 | |
| conf_district_id | 指定具体地域执行抓取 | 参考数据库district表 1 北京 2 上海 3 苏州 4 广东深圳 5 广东江门 （由于机器业务问题若有使用 具体咨询） |
| single | 是否一台adsl单次只执行一个任务（默认为0） | 0 默认 1 一次执行一次， 针对时效性任务 |
| expire_time | 任务过期时间（默认两小时） | |
| param | 其他参数 | a) filter : 1 过滤html， 0 默认不过滤 config={"param": {"filter": 1}} ； b) 设置截图大小 移动端截图大小 {"param": {"capture_width": 414, "capture_height": 736}} c) param method 可以手动设置get 或post 请求 |


#### SpiderRequest - urls

| 设置 | 作用 | 备注 |
| -------- | ------------------------ | ---- |
| url | 链接 | |
| type | 1(返回html， 默认)， 2(返回html和header) 4 截图 | |
| unique_key | 任务唯一标识，默认根据url去重 | 一般调用get_unique_key()生成 |



#### SpiderRequest - header 请求头设置

| 设置 | 作用 | 备注 |
| ---------- | ---- | ---- |
| User-Agent | | |
| Cookie | | |


#### 其他

| 设置 | 作用 | 备注 |
| ------- | ------------- | ----------- |
| get_unique_key() | 获取唯一标识字符串 | |
| add_black_ip(ip, type, expire_time=6) | 添加ip到抓取黑名单 | config : {param: {'task_type': 1}} |


#### 抓取结果 - 字典

| key | 含义 |  unique_md5
| ------------ | ------------------- |
| result | 抓取结果 |
| code | http状态码，默认为0 -> 200 |
| redirect_url | 跳转url |
| status | 抓取状态 0（未抓取）1（正在抓取）2（抓取成功）3（抓取失败） |
| header | 响应头 |
|inter_pro | 抓取ip |

#### BaseSpider 公共参数配置及调优

| 设置 | 作用 | 备注 |
| ---------- | ---- | ---- |
| url_repeat | 配置重复url是否抓取 如果发送url重复报“url send failure unique_md5” | True 重复的重发 False 重复的不重发(版本5.1.2默认False) |
| pc_user_agents | pc端ua | 可自行配置 |
| mb_user_agents | 移动端ua | 可自行配置 |
| url_repeat | True默认重复的url重发 False 不重发 | 可自行配置 |


#### 注意点
1、spider = DemoSpider(remote=True)
spider.run(1, 1, 1, 1, record_log=True)
remote True 线上(默认) False本地下载
record_log True 打印队列大小（调试用） False 默认不打印
2、各个队列说明，调试阶段 record_log 设置为True
sended_queue 过大，在等待下载中心下载（或者取速度慢）
response_queue 过大，拿到页面处理速度慢
store_queue 过大，存储速度慢
3、注意更新到下载中心最新版本
4、http状态码详解 http://tool.oschina.net/commons?type=5