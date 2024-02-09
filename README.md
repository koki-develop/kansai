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
kansAI is a CLI tool for converting text to Kansai dialect.
</p>

<p align="center">
English | <a href="./README.kansai.md">関西弁</a>
</p>

## Contents

- [Contents](#contents)
- [Installation](#installation)
  - [Homebrew Tap](#homebrew-tap)
  - [`go install`](#go-install)
  - [Releases](#releases)
- [Usage](#usage)
  - [Get API key](#get-api-key)
  - [Configure](#configure)
  - [Convert text to Kansai dialect](#convert-text-to-kansai-dialect)
- [LICENSE](#license)

## Installation

### Homebrew Tap

```console
$ brew install koki-develop/tap/kansai
```

### `go install`

```console
$ go install github.com/koki-develop/kansai@latest
```

### Releases

Download the binary from the [releases page](https://github.com/koki-develop/kansai/releases/latest).

## Usage

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

### Get API key

First, get the API key for the Gemini API from [Google AI Studio](https://makersuite.google.com/app/apikey).

### Configure

You can configure the API key by running kansai with the `--configure` option.

```console
$ kansai --configure
```

Alternatively, you can set the API key in the `KANSAI_API_KEY` environment variable or specify it with the `--api-key` option when running `kansai`.

```sh
$ export KANSAI_API_KEY='YOUR_API_KEY'
# or
$ echo '...' | kansai --api-key 'YOUR_API_KEY'
```

### Convert text to Kansai dialect

Just pass the text to be converted from standard input.

```console
$ echo 'こんにちは。僕の名前は koki です。好きな食べ物はラーメンです。' | kansai
```

```
こんちゃ。おれ、名前はこうきっちゅうねん。好きなもんはラーメンやで。
```

## LICENSE

[MIT](./LICENSE)
