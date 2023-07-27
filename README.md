# fluorite
[中文](README_zhcn.md)

A minimalist tool for generating static web pages from obsidian notes. Although there are many powerful third-party tools with the same functions, but because I don't want to use my brain to think and configure, I just want to simply share my notes with my friends (and I don't want to spend money to buy official services——), so this tool was born !

## feature
All it takes is: one executable + one css style = one website!

## How to use it

Unzip(or build) and place it in anywhere you like, make sure the theme directory and the executable file are at the same level.

Run like:

```
fluorite.exe -i MY_OBSIDIAN_NOTE_FOLDER -o TARGET_OUTPUT_FOLDER -t THEME_NAME
```

or create yaml config like:

```yaml
root: C:\library\docs\notes
theme: default
output: D:\Source\mynotes
include:
  - folder1
  - folder2
  - folder3
```

then run:

```
fluorite.exe -c THE_PATH_TO_YAML_CONFIG
```

finally you get a folder like:
- root level
  - main.css
  - index.html
  - notes
    - xxxxx.html
    - xxxxxxx.html
    - yyyyy.png
    - yyyy.jpg
    - ... ... ... ...
  
## Config File

### root
The directory of your obsidian notes folder, which is also the parent directory of the .obsidian folder.

### theme
The theme css. look at `theme/default`

The theme folder should be at the same level as the executable.

### output
The output Directory. This folder will be emptied every time it is generated, do not modify it manually.

### include
Only the defined folder prefix will be public. Note that if the attachment folder is not public, your pictures will not be displayed. If not defined, all folders are exposed.

## build
```
go build
```

## Theme
Now support custom theme. Look at `theme/default`.

So how to create a new theme?

1. Create a new folder under `theme` folder. Folder name is your theme name.
2. Create a `config.yaml` under `theme/your_theme`.
3. Complete your `config.yaml`! (just like `default/config.yaml`)

## TODO
- [ ] share button: Add a button to copy the URL and title of this page on each page.
- [ ] default css style: Now hardly any css styles are defined.
- [ ] Support for display in more formats.
