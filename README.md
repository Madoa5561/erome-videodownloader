# erome ダウンローダー

**erome** の動画やアルバムを簡単にダウンロードできるスクリプトです。  
`src` ディレクトリ内には **Python** と **Go** の2種類の実装があり、どちらも同じ動作をします。

[English README here](https://github.com/Madoa5561/erome-videodownloader/blob/main/README_en.md)

---

## 特徴

- URLを指定するだけで自動的に動画を取得  
- 単一ファイルまたはアルバム全体のダウンロードに対応  
- Python版・Go版の両方を提供  
- シンプルで軽量な設計  

---

## ディレクトリ構成

erome-downloader/
├── src/
│ ├── python/ # Python版スクリプト
│ └── go/ # Go版スクリプト
└── README.md


---

## 動作環境

- Python 3.x 以上  
- Go 1.20 以上  
- Windows / macOS / Linux 対応  

---

## 使い方

### Python版

```bash
cd src/python
pip install -r requirements.txt
python3 main.py
```

### Go版
```bash
cd src/go
go run main.go
```

> [!Note] 備考
> ダウンロード先ディレクトリはカレントフォルダ内に自動作成されます。
> 非公開アルバムの取得には対応していません。
> eromeの利用規約に従い、個人利用の範囲でご使用ください。

ライセンス

MIT License

[LICENSE URL](https://github.com/Madoa5561/erome-videodownloader/blob/main/LICENSE)

自由に改変・再配布可能です。



