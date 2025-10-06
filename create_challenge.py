#!/usr/bin/env python3
"""
Go & Rust 学習ジャーニー - 課題プロジェクト自動生成ツール

使用方法:
    python create_challenge.py <課題番号> <課題名>

例:
    python create_challenge.py 002 "単位換算CLI"
"""

import os
import sys
import re
from pathlib import Path

def sanitize_name(name, challenge_num=None):
    """課題名をファイル名として安全な形式に変換"""
    # 課題番号に基づく安全な名前のマッピング
    safe_names = {
        1: "calculator",
        2: "unit-converter",
        3: "bmi-calculator", 
        4: "password-generator",
        5: "json-formatter",
        6: "csv-to-json",
        7: "url-encoder-decoder",
        8: "base64-encoder-decoder",
        9: "uuid-generator",
        10: "file-hash-calculator",
        11: "file-counter",
        12: "todo-list",
        13: "memo",
        14: "vocabulary",
        15: "markdown-to-html",
        16: "log-monitor",
        17: "process-list",
        18: "time-sync",
        19: "weather-api",
        20: "ip-address",
        21: "http-status-checker",
        22: "web-scraper",
        23: "rss-reader",
        24: "file-copy",
        25: "file-diff",
        26: "zip-compress",
        27: "extract",
        28: "sqlite-memo",
        29: "keyword-search",
        30: "word-chain",
        31: "number-guessing",
        32: "rock-paper-scissors",
        33: "blackjack",
        34: "hangman",
        35: "stopwatch",
        36: "timer",
        37: "cron-scheduler"
    }
    
    # 課題番号から安全な名前を取得
    if challenge_num and challenge_num in safe_names:
        return safe_names[challenge_num]
    
    # フォールバック: 汎用的な名前を生成
    name = name.replace("CLI", "").replace(" ", "-").lower()
    # 日本語文字を除去して英数字とハイフンのみにする
    name = re.sub(r'[^\w\-]', '', name)
    # 空の場合はデフォルト名を使用
    if not name:
        name = "challenge"
    return name

def create_go_project(challenge_num, challenge_name, safe_name):
    """Goプロジェクトを作成"""
    project_name = f"go-challenge-{challenge_num:03d}-{safe_name}"
    project_path = Path("go") / project_name
    
    # ディレクトリ構造を作成
    directories = [
        project_path,
        project_path / "src",
        project_path / "test",
        project_path / "docs"
    ]
    
    for directory in directories:
        directory.mkdir(parents=True, exist_ok=True)
    
    # go.modファイルを作成
    go_mod_content = f"""module {project_name}

go 1.25.1
"""
    (project_path / "go.mod").write_text(go_mod_content, encoding='utf-8')
    
    # README.mdを作成
    readme_content = f"""# {challenge_name} (Go版)

## 概要
{challenge_name}のGo実装です。

## 実行方法
```bash
go run src/main.go
```

## テスト実行
```bash
go test ./test/...
```

## 課題番号
{challenge_num:03d}
"""
    (project_path / "README.md").write_text(readme_content, encoding='utf-8')
    
    # main.goファイルを作成
    main_go_content = f"""package main

import "fmt"

func main() {{
    fmt.Println("Hello, {challenge_name}!")
}}
"""
    (project_path / "src" / "main.go").write_text(main_go_content, encoding='utf-8')
    
    # テストファイルを作成
    test_go_content = f"""package main

import "testing"

func Test{challenge_name.replace(" ", "").replace("CLI", "")}(t *testing.T) {{
    // TODO: テストを実装
    t.Log("テストを実装してください")
}}
"""
    (project_path / "test" / f"{safe_name}_test.go").write_text(test_go_content, encoding='utf-8')
    
    print(f"Goプロジェクト作成完了: {project_path}")
    return project_path

def create_rust_project(challenge_num, challenge_name, safe_name):
    """Rustプロジェクトを作成"""
    project_name = f"rust-challenge-{challenge_num:03d}-{safe_name}"
    project_path = Path("rust") / project_name
    
    # ディレクトリ構造を作成
    directories = [
        project_path,
        project_path / "src",
        project_path / "test",
        project_path / "docs"
    ]
    
    for directory in directories:
        directory.mkdir(parents=True, exist_ok=True)
    
    # Cargo.tomlファイルを作成
    cargo_toml_content = f"""[package]
name = "{project_name}"
version = "0.1.0"
edition = "2021"

[dependencies]
"""
    (project_path / "Cargo.toml").write_text(cargo_toml_content, encoding='utf-8')
    
    # README.mdを作成
    readme_content = f"""# {challenge_name} (Rust版)

## 概要
{challenge_name}のRust実装です。

## 実行方法
```bash
cargo run
```

## テスト実行
```bash
cargo test
```

## 課題番号
{challenge_num:03d}
"""
    (project_path / "README.md").write_text(readme_content, encoding='utf-8')
    
    # main.rsファイルを作成
    main_rs_content = f"""fn main() {{
    println!("Hello, {challenge_name}!");
}}
"""
    (project_path / "src" / "main.rs").write_text(main_rs_content, encoding='utf-8')
    
    # テストファイルを作成
    test_rs_content = f"""#[cfg(test)]
mod tests {{
    use super::*;

    #[test]
    fn test_{safe_name}() {{
        // TODO: テストを実装
        assert!(true, "テストを実装してください");
    }}
}}
"""
    (project_path / "test" / f"{safe_name}_test.rs").write_text(test_rs_content, encoding='utf-8')
    
    print(f"Rustプロジェクト作成完了: {project_path}")
    return project_path

def update_main_readme(challenge_num, challenge_name):
    """メインのREADME.mdを更新"""
    readme_path = Path("README.md")
    if not readme_path.exists():
        print("メインのREADME.mdが見つかりません")
        return
    
    content = readme_path.read_text(encoding='utf-8')
    
    # 課題番号に対応するチェックボックスを探して更新
    # この部分は手動で更新する必要があります
    print(f"メインのREADME.mdで課題 {challenge_num:03d}: {challenge_name} を完了済みにマークしてください")

def main():
    if len(sys.argv) != 3:
        print("使用方法: python create_challenge.py <課題番号> <課題名>")
        print("例: python create_challenge.py 002 '単位換算CLI'")
        sys.exit(1)
    
    try:
        challenge_num = int(sys.argv[1])
        challenge_name = sys.argv[2]
    except ValueError:
        print("課題番号は数値で入力してください")
        sys.exit(1)
    
    safe_name = sanitize_name(challenge_name, challenge_num)
    
    print(f"課題 {challenge_num:03d}: {challenge_name} のプロジェクトを作成中...")
    
    # Goプロジェクトを作成
    go_path = create_go_project(challenge_num, challenge_name, safe_name)
    
    # Rustプロジェクトを作成
    rust_path = create_rust_project(challenge_num, challenge_name, safe_name)
    
    # メインREADMEの更新を案内
    update_main_readme(challenge_num, challenge_name)
    
    print(f"\n課題 {challenge_num:03d} のプロジェクト作成完了!")
    print(f"Go: {go_path}")
    print(f"Rust: {rust_path}")
    print(f"\n次のステップ:")
    print(f"1. 各プロジェクトで実装を開始")
    print(f"2. メインのREADME.mdで進捗を更新")
    print(f"3. テストを実装して動作確認")

if __name__ == "__main__":
    main()
