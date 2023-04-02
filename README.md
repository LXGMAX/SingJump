# 唱跳

一个golang编写的电脑控制本地播放器服务端

## 接口

| 路径 | 功能 | 返回 |
|-------------|----------|---------|
| `/play/pause`| 暂停/播放 | `pause` |
| `/play/next` | 下一首 | `next` |
| `/play/previous` | 上一首 | `previous` |
| `/play/volup` | 音量加 | `volume`(int) |
| `/play/voldown` | 音量减 | `volume`(int) |
| `/play/like` | 歌曲红心 | `like` |
| `/play/lyric` | 歌词控件开关 | `lyric` |
| `/play/mute` | 静音 | `volume`(int) |
| `/sys/getvol` | 获取音量 | `volume`(int) |