# fluorite

# How to use it

Unzip and place it where you like, make sure the theme directory and the executable file are at the same level.

Run like:

```
fluorite.exe -i MY_OBSIDIAN_NOTE_FOLDER -o TARGET_OUTPUT_FOLDER -t THEME_NAME
```

or create yaml config like:

```yaml
root: C:\library\docs\notes
theme: default
output: D:\Source\mynotes
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