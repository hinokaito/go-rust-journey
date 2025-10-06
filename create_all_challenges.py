#!/usr/bin/env python3
"""
Go & Rust 学習ジャーニー - 全課題プロジェクト一括生成ツール

使用方法:
    python create_all_challenges.py
"""

import os
import sys
from pathlib import Path
from create_challenge import create_go_project, create_rust_project, sanitize_name

# 課題リスト（README.mdから抽出）
CHALLENGES = [
    "四則演算CLI",
    "単位換算CLI", 
    "BMI計算CLI",
    "パスワード生成CLI",
    "JSON整形CLI",
    "CSV→JSON変換CLI",
    "URLエンコード/デコードCLI",
    "Base64エンコード/デコードCLI",
    "UUID生成CLI",
    "ファイルハッシュ計算CLI",
    "フォルダ内ファイル数カウントCLI",
    "ToDoリストCLI",
    "メモ帳CLI",
    "英単語帳CLI",
    "Markdown→HTML変換CLI",
    "ログファイル監視CLI",
    "プロセス一覧取得CLI",
    "時刻同期CLI",
    "天気API取得CLI",
    "IPアドレス取得CLI",
    "HTTPステータスチェッカーCLI",
    "WebスクレイパーCLI",
    "RSSリーダーCLI",
    "ファイルコピーCLI",
    "ファイル差分比較CLI",
    "圧縮(zip)CLI",
    "解凍CLI",
    "SQLiteメモ帳CLI",
    "キーワード検索CLI",
    "しりとりCLI",
    "数当てゲームCLI",
    "じゃんけんゲームCLI",
    "ブラックジャックCLI",
    "ハングマンCLI",
    "ストップウォッチCLI",
    "タイマーCLI",
    "cron風スケジューラーCLI",
    "GitHubリポ一覧取得CLI",
    "APIレスポンス速度計測CLI",
    "並列ダウンロードCLI",
    "並列アップロードCLI",
    "goroutine数測定CLI",
    "チャットサーバー(TCP)",
    "チャットクライアント(TCP)",
    "ファイル転送サーバー",
    "ファイル転送クライアント",
    "HTTPサーバー(静的HTML)",
    "HTTPサーバー(JSON API)",
    "簡易REST APIサーバー",
    "メール送信CLI(SMTP)",
    "メール受信CLI(IMAP)",
    "RSSフィード生成CLI",
    "URL短縮サービスCLI",
    "QRコード生成CLI",
    "QRコード読み取りCLI",
    "HTMLスクリーンショットCLI",
    "ローカルDNSキャッシュCLI",
    "画像リサイズCLI",
    "画像フォーマット変換CLI",
    "PDF→画像変換CLI",
    "PDF結合CLI",
    "PDF分割CLI",
    "パケットキャプチャCLI",
    "ログローテーションCLI",
    "アクセス解析CLI",
    "HTMLテンプレートサーバー",
    "JSON Web Token認証サーバー",
    "OAuthクライアントCLI",
    "OAuthサーバー",
    "WebSocketチャットサーバー",
    "WebSocketチャットクライアント",
    "プロキシサーバー",
    "URLフィルタリングサーバー",
    "TCPポートスキャナCLI",
    "UDPポートスキャナCLI",
    "ファイル暗号化CLI",
    "ファイル復号CLI",
    "キーストアCLI",
    "HMAC生成CLI",
    "パスワード強度判定CLI",
    "コンテナ風サンドボックスCLI",
    "ミニDocker風CLI",
    "キュー処理ワーカーCLI",
    "Pub/SubシステムCLI",
    "Kafka風メッセージングCLI",
    "MySQL接続CLI",
    "PostgreSQL接続CLI",
    "RedisキャッシュCLI",
    "タスクキューCLI",
    "並列ソートCLI",
    "マージソートCLI",
    "クイックソートCLI",
    "グラフ探索CLI",
    "DijkstraアルゴリズムCLI",
    "BFS/DFSアルゴリズムCLI",
    "JSON SchemaバリデータCLI",
    "ConfigファイルローダーCLI",
    "環境変数管理CLI",
    "CLIフレームワーク作成",
    "独自言語インタプリタCLI",
    "タスク管理REST API",
    "Slack風CLIチャット",
    "GitHub Issueクローン",
    "Pastebin風API",
    "JSON-Linter CLI",
    "ログ監視デーモン",
    "Discord Bot (Go製)",
    "WebSocketベース白板アプリ",
    "URLモニタリングサーバー",
    "Jenkins風ジョブ実行サーバー",
    "Docker StatsビューアCLI",
    "mini-S3風オブジェクトストレージ",
    "CLIベースGitクローン",
    "SSHファイルブラウザ",
    "OpenAI APIクライアント",
    "MarkdownプレビューHTTPサーバー",
    "REST APIテストランナー",
    "HTMLテンプレートエンジン自作",
    "ショートリンクWebアプリ(Gin)",
    "ブックマーク同期サーバー",
    "CLIインタラクティブシェル",
    "ファイルウォッチャー＆自動バックアップ",
    "音声→文字変換CLI",
    "CLI辞書クライアント(API接続)",
    "HTTPキャッシュプロキシ",
    "JSON RPCサーバー",
    "WebSocketメッセージブローカー",
    "Discord風Webチャット(Gin+WS)",
    "CLIニュースアグリゲータ",
    "Go言語コード整形器",
    "テキスト差分ビジュアライザCLI",
    "並列ファイルアップローダ",
    "自動翻訳CLI",
    "画像→ASCIIアート変換CLI",
    "CLI音楽プレイヤー(MP3再生)",
    "cron式スケジューラライブラリ",
    "ファイル改ざん検知デーモン",
    "CLIテキスト整形ツール群(cat/grep風)",
    "マルチスレッドWebクローラ",
    "Gzip対応HTTPサーバー",
    "OAuth2対応認証サーバー",
    "JWT認証付きREST API",
    "MySQL接続プール実装",
    "Redisキャッシュラッパー",
    "GraphQLサーバー実装",
    "WebAssembly実行環境",
    "Goプラグインローダー",
    "ローカルHTTP負荷試験ツール",
    "JSON Schemaジェネレーター",
    "XML→JSON変換CLI",
    "mini-Dockerランタイム",
    "Goスクリプトエンジン実装",
    "TCPロードバランサ",
    "UNIXソケット通信サーバー",
    "CLIブックマーク検索エンジン",
    "RSS→メール配信サーバー",
    "SMTP受信サーバー",
    "Webhook集約サーバー",
    "OAuthクライアントダッシュボード",
    "ファイル暗号化REST API",
    "画像サムネイル自動生成API",
    "リアルタイムメトリクス可視化サーバー",
    "mini-Prometheus風メトリクス収集ツール",
    "ディスク容量監視CLI",
    "Go Module可視化CLI",
    "CLI依存関係解析ツール",
    "WASM実験用HTTPサーバー",
    "TCPパケット解析CLI",
    "CLIログ整形＋カラー出力ツール",
    "シンプルなCI/CDツール",
    "コンフィグ管理サーバー(YAML中心)",
    "YAML/JSON設定統合ライブラリ",
    "CLI環境変数テンプレート展開ツール",
    "CLIコードスニペット管理ツール",
    "CLIで進捗バー付きダウンローダ",
    "マルチスレッドZIP圧縮CLI",
    "Webカメラ映像ストリーミングサーバー",
    "mini-Grafana風ダッシュボード",
    "CLI PDFメタデータ編集ツール",
    "CLIで画像透かし追加ツール",
    "CSVデータビジュアライザ(Go+Fyne)",
    "GUI版ToDoアプリ(Fyne)",
    "ローカル音楽管理GUI(Fyne)",
    "CLIノート全文検索エンジン",
    "Markdown管理＋Git自動コミットツール",
    "ローカルサーバーログ可視化ツール",
    "CLIスレッドプール実装",
    "ファイル監査システム(API連携)",
    "プロジェクト依存関係ツリー描画CLI",
    "mini-Gitリポジトリ実装",
    "mini-Redis実装",
    "mini-HTTPサーバー from scratch",
    "TLSハンドシェイク学習サーバー",
    "DNSクエリリゾルバ",
    "DNSサーバー実装",
    "CLIトレースルートツール",
    "CLI pingツール",
    "CLIネットワークスピード測定",
    "Go製ローカルWeb IDE",
    "Goコンパイラ内部構造模倣プロジェクト"
]

def create_all_projects():
    """全課題のプロジェクトを作成"""
    print(f"🚀 全{len(CHALLENGES)}課題のプロジェクトを作成開始...")
    
    created_count = 0
    skipped_count = 0
    
    for i, challenge_name in enumerate(CHALLENGES, 1):
        challenge_num = i
        safe_name = sanitize_name(challenge_name)
        
        print(f"\n📝 課題 {challenge_num:03d}: {challenge_name}")
        
        try:
            # Goプロジェクトを作成
            go_path = Path("go") / f"go-challenge-{challenge_num:03d}-{safe_name}"
            if not go_path.exists():
                create_go_project(challenge_num, challenge_name, safe_name)
                created_count += 1
            else:
                print(f"⏭️  Goプロジェクトは既に存在: {go_path}")
                skipped_count += 1
            
            # Rustプロジェクトを作成
            rust_path = Path("rust") / f"rust-challenge-{challenge_num:03d}-{safe_name}"
            if not rust_path.exists():
                create_rust_project(challenge_num, challenge_name, safe_name)
                created_count += 1
            else:
                print(f"⏭️  Rustプロジェクトは既に存在: {rust_path}")
                skipped_count += 1
                
        except Exception as e:
            print(f"❌ エラー: {e}")
            continue
    
    print(f"\n🎉 完了!")
    print(f"✅ 新規作成: {created_count}プロジェクト")
    print(f"⏭️  スキップ: {skipped_count}プロジェクト")
    print(f"📊 合計: {len(CHALLENGES)}課題 × 2言語 = {len(CHALLENGES) * 2}プロジェクト")

def main():
    print("Go & Rust 学習ジャーニー - 全課題プロジェクト一括生成")
    print("=" * 60)
    
    # 確認
    response = input(f"全{len(CHALLENGES)}課題のプロジェクトを作成しますか？ (y/N): ")
    if response.lower() != 'y':
        print("キャンセルしました")
        return
    
    create_all_projects()

if __name__ == "__main__":
    main()

