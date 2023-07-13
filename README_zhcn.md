# fluorite

一个极简的obsidian笔记生成静态网页的工具。虽然有很多强大的功能相同的第三方工具，但由于我不想动脑思考和配置，只是想简单的把笔记分享给朋友（同时也不想花钱买官方的服务——），因此本工具便诞生了！

## 特色
只需要：一个可执行文件 + 单个 css 样式 = 一个网站！

## 如何使用它

解压缩（或构建）并将其放在您喜欢的任何位置，确保主题目录和可执行文件位于同一级别。

运行如下：

```
fluorite.exe -i MY_OBSIDIAN_NOTE_FOLDER -o TARGET_OUTPUT_FOLDER -t THEME_NAME
```

或创建 yaml 配置，例如：

```yaml
root: C:\library\docs\notes
theme: default
output: D:\Source\mynotes
include:
  - folder1
  - folder2
  - folder3
```

然后运行：

```
fluorite.exe -c THE_PATH_TO_YAML_CONFIG
```

最后你会得到一个像这样的文件夹：
- root level
  - main.css
  - index.html
  - notes
    - xxxxx.html
    - xxxxxxx.html
    - yyyyy.png
    - yyyy.jpg
    - ... ... ... ...
  
## 配置文件

### root
obsidian笔记文件夹的目录，也是 .obsidian 文件夹的父目录。

### theme
主题CSS。 查看`theme/default/main.css`

主题文件夹应与可执行文件位于同一级别。

### output
输出目录。 该文件夹每次生成都会被清空，请勿手动修改。

### include
只有定义的文件夹前缀才是公开的。 请注意，如果附件文件夹不公开，则不会显示您的图片。 如果未定义，则所有文件夹都会公开。

## build
```
go build
```

## TODO
- [ ] 分享按钮：在每个页面添加一个按钮，用于复制该页面的 URL 和标题。
- [ ] 默认 css 样式：现在几乎没有定义任何 css 样式。
- [ ] 支持更多格式显示。