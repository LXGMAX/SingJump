# 唱跳

一个golang编写的电脑控制本地播放器服务端

## 接口

| 路径 | 请求类型 | 功能 | 返回 | 错误 |
|-------------|-------|----------|---------|---|
| `/play/pause` | GET | 暂停/播放 | `pause` |
| `/play/next` | GET | 下一首 | `next` |
| `/play/previous` | GET | 上一首 | `previous` |
| `/play/volup` | GET | 音量加 | `volume`(int) |
| `/play/voldown` | GET | 音量减 | `volume`(int) |
| `/play/like` | GET | 歌曲红心 | `like` |
| `/play/lyric` | GET | 歌词控件开关 | `lyric` |
| `/play/mute` | GET | 静音 | `volume`(int) |
| `/sys/getvol` | GET | 获取音量 | `volume`(int) |
| `/sys/setvol` | POST | 设置音量 | `volume`(int) | `error`(string) |

## 用法

执行程序会显示出本机IP和端口，用GET访问`ip:port/路径`即可

eg:
1. GET `{{ip}}:{{port}}/sys/getvol`
得到
```json
{
    "volume": 10
}
```

2. POST `{{ip}}:{{port}}/sys/setvol`
CONTENT 
```json
{
    "setSysVolume":20
}
```
得到
```json
{
    "volume": 20
}
```
超出音量范围时
```json
{
    "error": "invalid volume range"
}
```

## 手机客户端

[RapMusic](https://github.com/LXGMAX/RapMusic)
