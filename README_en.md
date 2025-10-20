## English README

[日本語版はこちら](https://github.com/Madoa5561/erome-videodownloader/blob/main/README.md)

erome Downloader is a simple script to download videos or albums from erome.
Both Python and Go implementations are available under the src directory, performing the same behavior.

Features

Download single videos or full albums by URL

Supports both Python and Go

Works on Windows, macOS, and Linux

Lightweight and easy to use

Directory Structure
erome-downloader/
├── src/
│   ├── python/    # Python version
│   └── go/        # Go version
└── README.md

## Requirements

Python 3.x

Go 1.20+

Usage
### Python Version
```bash
cd src/python
python main.py
```

### Go Version
```bash
cd src/go
go run main.go
```

> [!Note]
> The download folder is automatically created in the current directory.
> 
> Private albums are not supported.
> 
> Please use this script for personal purposes only, following erome’s Terms of Service.

License

MIT License — free to use, modify, and distribute.
