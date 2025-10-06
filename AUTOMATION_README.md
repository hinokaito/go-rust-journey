# 自動化ツール使用ガイド

このプロジェクトには、課題管理を自動化するためのPythonツールが含まれています。

## ツール一覧

### 1. `create_challenge.py` - 単一課題作成ツール

新しい課題のプロジェクトを作成します。

```bash
# 使用方法
python create_challenge.py <課題番号> <課題名>

# 例
python create_challenge.py 002 "単位換算CLI"
```

**作成されるもの:**
- `go/go-challenge-002-unit-converter/` (Goプロジェクト)
- `rust/rust-challenge-002-unit-converter/` (Rustプロジェクト)

各プロジェクトには以下が含まれます:
- 適切な設定ファイル (`go.mod`, `Cargo.toml`)
- 基本的なソースファイル (`main.go`, `main.rs`)
- テストファイル
- README.md
- ドキュメント用ディレクトリ

### 2. `create_all_challenges.py` - 全課題一括作成ツール

全200課題のプロジェクトを一括で作成します。

```bash
python create_all_challenges.py
```

⚠️ **注意**: このツールは大量のファイルを作成するため、実行前に確認が求められます。

### 3. `progress_manager.py` - 進捗管理ツール

学習の進捗を追跡・管理します。

```bash
# 進捗状況を表示
python progress_manager.py status

# 統計情報を表示
python progress_manager.py stats

# 課題完了をマーク
python progress_manager.py complete <課題番号> <言語>
```

## 使用例

### 新しい課題を開始する場合

```bash
# 課題2: 単位換算CLIを作成
python create_challenge.py 2 "単位換算CLI"

# 作成されたプロジェクトで作業開始
cd go/go-challenge-002-unit-converter
# または
cd rust/rust-challenge-002-unit-converter
```

### 進捗を確認する場合

```bash
# 現在の進捗状況を確認
python progress_manager.py status

# 詳細な統計情報を表示
python progress_manager.py stats
```

### 課題を完了した場合

```bash
# Go実装を完了としてマーク
python progress_manager.py complete 1 go

# Rust実装を完了としてマーク
python progress_manager.py complete 1 rust
```

## プロジェクト構造

自動化ツールが作成する標準的なプロジェクト構造:

```
go/go-challenge-XXX-name/
├── docs/           # ドキュメント
├── src/           # ソースコード
│   └── main.go    # メインファイル
├── test/          # テストファイル
├── go.mod         # Goモジュール設定
└── README.md      # プロジェクト説明

rust/rust-challenge-XXX-name/
├── docs/          # ドキュメント
├── src/           # ソースコード
│   └── main.rs    # メインファイル
├── test/          # テストファイル
├── Cargo.toml     # Rustプロジェクト設定
└── README.md      # プロジェクト説明
```

## カスタマイズ

### 課題名の変換ルール

課題名は以下のルールでファイル名に変換されます:
- 日本語 → ローマ字（簡易版）
- スペース → ハイフン
- 特殊文字 → 除去
- 小文字に変換

例: `"単位換算CLI"` → `"unit-converter"`

### テンプレートのカスタマイズ

各ツールのテンプレート部分を編集することで、生成されるファイルの内容をカスタマイズできます。

## トラブルシューティング

### よくある問題

1. **課題番号が重複している**
   - 既存のプロジェクトを確認してから実行してください

2. **権限エラー**
   - ファイル作成権限があることを確認してください

3. **Pythonの依存関係**
   - Python 3.6以上が必要です

### ログとデバッグ

各ツールは詳細な出力を提供します。エラーが発生した場合は、出力メッセージを確認してください。

## 貢献

自動化ツールの改善提案やバグ報告は歓迎します。新しい機能や改善案があれば、IssueまたはPull Requestでお知らせください。

