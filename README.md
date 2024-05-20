# voicevoxcore.go

[![codecov](https://codecov.io/gh/sh1ma/voicevoxcore.go/branch/main/graph/badge.svg?token=5111EKB6SF)](https://codecov.io/gh/sh1ma/voicevoxcore.go)
[![CI: Test](https://github.com/sh1ma/voicevoxcore.go/actions/workflows/test.yaml/badge.svg)](https://github.com/sh1ma/voicevoxcore.go/actions/workflows/test.yaml)
[![golangci-lint](https://github.com/sh1ma/voicevoxcore.go/actions/workflows/lint.yaml/badge.svg)](https://github.com/sh1ma/voicevoxcore.go/actions/workflows/lint.yaml)

voicevoxcore.go は[voicevox_core](https://github.com/VOICEVOX/voicevox_core)を Go 言語で使えるラッパーです。
FFI を用いて、voicevox_core の C API を呼んでいます。

## インストール

```
go get https://github.com/sh1ma/voicevoxcore.go
```

## 例

以下は本ライブラリを使用して Text to Speech を行う例です。

```go

//go:build ignore

package main

import (
	"fmt"
	"os"

	voicevoxcorego "github.com/sh1ma/voicevoxcore.go"
)

func main() {
	text := "ずんだもんなのだ"

	core := voicevoxcorego.New()
	initializeOptions := voicevoxcorego.NewVoicevoxInitializeOptions(0, 0, false, "./open_jtalk_dic_utf_8-1.11")
	core.Initialize(initializeOptions)

	core.LoadModel(3)

	ttsOptions := voicevoxcorego.NewVoicevoxTtsOptions(false, true)
	result, err := core.Tts(text, 1, ttsOptions)
	if err != nil {
		fmt.Println(err)
	}
	f, _ := os.Create("out.wav")
	_, err = f.Write(result)
	if err != nil {
		fmt.Println(err)
	}
}
```

## おすすめの環境構築方法 (Linux / MacOS)

本ライブラリを使用するには openJTalk の辞書ファイルと voicevox_core の動的ライブラリ、ヘッダファイル、そしてモデルファイルが必要になります。
以下でそれらをダウンロードし、本ライブラリから使えるようにするための簡単なセットアップの手順を説明します。

### 1. voicevox_core のダウンロード

[voicevox_core の releases](https://github.com/VOICEVOX/voicevox_core/releases)から自分の OS、アーキテクチャに合ったダウンローダをダウンロードし、実行してください。実行するとカレントディレクトリに`voicevox_core_*`のディレクトリが配置されます。直下には以下のものが入っています

- openJTalk の辞書ファイルが入った`open_jtalk_dic_*/`ディレクトリ
- voicevox_core の動的ライブラリ
  - 拡張子は.dll, .dylib, .so のいずれかです
- 圧縮されたモデルファイルの入った`model/`ディレクトリ

### 2. voicevox_core を配置する

`voicevox_core`を任意のパスに移動(おすすめは`~/.local`のなか)します

### 3. シンボリックリンクを張る

`voicevox_core`内にある 2 つのファイルのシンボリックリンクを作ります。
以下のようなコマンドを実行します。(プラットフォームによってコマンドが異なる場合があります)。動的ライブラリが`.dylib`の場合を例に挙げます。

**注意: ln に渡すパスは相対パスではなく絶対パスにしてください**

```sh
# [VOICEVOX_CORE_DIR] を`voicevox_core`の絶対パスにします

# 動的ライブラリのシンボリックリンクを`/usr/local/lib`に配置します
ln -s [VOICEVOX_CORE_DIR]/libvoicevox_core.dylib /usr/local/lib

# 動的ライブラリのシンボリックリンクを`/usr/local/include`に配置します
ln -s [VOICEVOX_CORE_DIR]/voicevox_core.h /usr/local/include
```

以上の手順で本ライブラリが使えるようになっているはずです。
