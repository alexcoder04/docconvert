
# docconvert

docconvert - more than a document converter (even if first doesn't sound like that).

## Features

 - Convert documents from/to different formats ([pandoc](https://pandoc.org/) required)
 - Extract text from images ([tesseract](https://github.com/tesseract-ocr/tesseract) required)
 - Generate random numbers or pick random items from lists

## Build and run

Install [pandoc](https://pandoc.org/installing.html) and [tesseract](https://github.com/tesseract-ocr/tesseract#installing-tesseract).

```sh
git clone https://github.com/alexcoder04/docconvert
cd docconvert
go build
./docconvert
```

## Command-line arguments

 - `-port`: port to run on, default 8080
 - `-all-hosts`: whether to listen on all hosts, not just localhost
