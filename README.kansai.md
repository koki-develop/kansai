<p align="center">
<img src="./assets/logo_light.svg#gh-light-mode-only" />
<img src="./assets/logo_dark.svg#gh-dark-mode-only" />
</p>

<p align="center">
Powered by Gemini API.
</p>

<p align="center">
<a href="https://github.com/koki-develop/kansai/releases/latest"><img src="https://img.shields.io/github/v/release/koki-develop/kansai" alt="GitHub release (latest by date)"></a>
<a href="https://github.com/koki-develop/kansai/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/koki-develop/kansai/ci.yml?logo=github" alt="GitHub Workflow Status"></a>
<a href="https://codeclimate.com/github/koki-develop/kansai/maintainability"><img src="https://img.shields.io/codeclimate/maintainability/koki-develop/kansai?style=flat&amp;logo=codeclimate" alt="Maintainability"></a>
<a href="https://goreportcard.com/report/github.com/koki-develop/kansai"><img src="https://goreportcard.com/badge/github.com/koki-develop/kansai" alt="Go Report Card"></a>
<a href="./LICENSE"><img src="https://img.shields.io/github/license/koki-develop/kansai" alt="LICENSE"></a>
</p>

<p align="center">
kansAI はテキストを関西弁に変換する CLI ツールや。
</p>

<p align="center">
<a href="./README.md">English</a> | 関西弁
</p>

## 目次

- [目次](#目次)
- [インストール](#インストール)
  - [Homebrew Tap](#homebrew-tap)
  - [`go install`](#go-install)
  - [Releases](#releases)
- [使い方](#使い方)
  - [API キーを取得](#api-キーを取得)
  - [設定](#設定)
  - [関西弁に変換](#関西弁に変換)
- [ライセンス](#ライセンス)

## インストール

### Homebrew Tap

```console
$ brew install koki-develop/tap/kansai
```

### `go install`

```console
$ go install github.com/koki-develop/kansai@latest
```

### Releases

[リリースページ](https://github.com/koki-develop/kansai/releases/latest)からバイナリをダウンロードしてや。

## 使い方

```console
$ kansai --help
CLI tool for converting text to Kansai dialect.

Usage:
  kansai [flags]

Flags:
  -k, --api-key string   API Key for the Gemini API
      --configure        configure API key
  -h, --help             help for kansai
  -v, --version          version for kansai
```

### API キーを取得

まず [Google AI Studio](https://makersuite.google.com/app/apikey) から Gemini API の API キーを取得してや。

### 設定

`kansai` を `--configure` オプション付きで実行することで API キーを設定することができるで。

```console
$ kansai --configure
```

他の方法としては、 API キーを `KANSAI_API_KEY` っちゅう環境変数に設定するか、 `kansai` を動かす時に `--api-key` っちゅうオプションで指定すんのもありやで。

```sh
$ export KANSAI_API_KEY='YOUR_API_KEY'
# or
$ echo '...' | kansai --api-key 'YOUR_API_KEY'
```

### 関西弁に変換

標準入力から変換するテキストを渡すだけや。

```console
$ echo 'こんにちは。僕の名前は koki です。好きな食べ物はラーメンです。' | kansai
```

```
こんちゃ。おれ、名前はこうきっちゅうねん。好きなもんはラーメンやで。
```

## ライセンス

[MIT](./LICENSE)
