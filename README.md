# readme-bot
Automating README.md updates via GitHub Actions on pull request(or push) creation.  
プルリクエスト作成またはPushをトリガーにしてGithub ActionsからキックされるREADME.md自動更新Botを作りたい。

---

## What is this?

`readme-bot` is a small CLI tool written in Go that:

1. Reads a GitHub pull request diff
2. Sends the diff + repository rules to an LLM
3. Gets a one-line summary suitable for the latest change section in `README.md`
4. Writes that summary back into `README.md` (and later into `CHANGELOG.md`)

Go 製の小さな CLI ツールで、

1. GitHub のプルリクエスト差分を取得し  
2. Diff と RULES（ルール）を LLM に渡して要約してもらい  
3. `README.md` の「最新の変更」セクションに追記する（将来的には `CHANGELOG.md` も）

---

## Features / できること

- ✅ GitHub Pull Request の diff を API 経由で取得（`go-github` 利用）
- ✅ LLM（OpenAI API）に diff + RULES を投げて JSON 形式でレスポンスを受け取る
- ✅ `README.md` の特定セクション（例: `## latest change`）を書き換え or 追記
- ✅ リポジトリ固有のルール（`RULES.md`）を埋め込み + 環境変数で上書き可能
- 🛠 GitHub Actions からの自動実行に対応予定（現在一部実験中）
- 📝 将来的に `CHANGELOG.md` 自動更新にも対応予定

---

## Requirements / 必要要件

- Go `1.24.x` 以上
- GitHub リポジトリ
- OpenAI API キー

GitHub Actions 上で動かす場合は、以下が使われます：

- `GITHUB_TOKEN`（GitHub Actions が自動で注入してくれるトークン）
- `OPENAI_API_KEY`（Actions のリポジトリシークレットから渡す）

---

## Setup / セットアップ

### Clone

```bash
git clone https://github.com/<your-account>/readme-bot.git
cd readme-bot
```

## Environment variables / 環境変数
ローカル実行 :  
```bash
export GITHUB_TOKEN="your github token"
export OPENAI_API_KEY="your openai api key"
# 任意: カスタム RULES.md を使いたい場合
export RULES_PATH="./docs/RULES.custom.md"
```

```powershell
$env:GITHUB_TOKEN = "your github token"
$env:OPENAI_API_KEY = "your openai api key"
$env:RULES_PATH = ".\docs\RULES.custom.md"
```

## Usage / Local
```bash
go run ./cmd/readme-bot \
  -owner <github-owner> \
  -repo <repository-name> \
  -number <pull-request-number>
```

## Usage / GitHub Actions
Workflow sample:  
```yaml
name: readme-bot

on:
  workflow_dispatch:

jobs:
  run-readme-bot:
    runs-on: ubuntu-latest

    steps:
      - name: Check out this repository
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: "1.24.0"

      - name: Run README Bot
        run: go run ./cmd/readme-bot -owner ${{ github.repository_owner }} -repo ${{ github.event.repository.name }} -number 7
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          OPENAI_API_KEY: ${{ secrets.OPENAI_API_KEY }}
```

## Roadmap
- CHANGELOG.md 自動更新機能
- GitHub Actions の pull_request イベントに正式対応
- LLM バックエンドの切り替え（OpenAI 以外の API 対応）
- PR コメントとして要約を自動投稿するモード
- 設定ファイル（YAML）によるルール／動作のカスタマイズ

## latest change
README Botの機能を強化し、レビュー機能を追加しました。
