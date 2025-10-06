#!/usr/bin/env python3
"""
Go & Rust 学習ジャーニー - 進捗管理ツール

使用方法:
    python progress_manager.py status    # 進捗状況を表示
    python progress_manager.py complete <課題番号> <言語>  # 課題完了をマーク
    python progress_manager.py stats     # 統計情報を表示
"""

import os
import sys
import re
from pathlib import Path
from datetime import datetime

def get_challenge_status():
    """全課題の状況を取得"""
    go_dir = Path("go")
    rust_dir = Path("rust")
    
    status = {
        'go': {'total': 0, 'completed': 0, 'in_progress': 0},
        'rust': {'total': 0, 'completed': 0, 'in_progress': 0}
    }
    
    # Goプロジェクトの状況をチェック
    if go_dir.exists():
        for project_dir in go_dir.iterdir():
            if project_dir.is_dir() and project_dir.name.startswith('go-challenge-'):
                status['go']['total'] += 1
                
                # main.goの存在と内容で進捗を判定
                main_go = project_dir / "src" / "main.go"
                if main_go.exists():
                    content = main_go.read_text(encoding='utf-8')
                    if "Hello, world!" not in content and len(content.strip()) > 50:
                        status['go']['completed'] += 1
                    else:
                        status['go']['in_progress'] += 1
    
    # Rustプロジェクトの状況をチェック
    if rust_dir.exists():
        for project_dir in rust_dir.iterdir():
            if project_dir.is_dir() and project_dir.name.startswith('rust-challenge-'):
                status['rust']['total'] += 1
                
                # main.rsの存在と内容で進捗を判定
                main_rs = project_dir / "src" / "main.rs"
                if main_rs.exists():
                    content = main_rs.read_text(encoding='utf-8')
                    if "Hello, world!" not in content and len(content.strip()) > 50:
                        status['rust']['completed'] += 1
                    else:
                        status['rust']['in_progress'] += 1
    
    return status

def show_status():
    """進捗状況を表示"""
    status = get_challenge_status()
    
    print("Go & Rust 学習ジャーニー - 進捗状況")
    print("=" * 50)
    
    for lang in ['go', 'rust']:
        lang_name = "Go" if lang == 'go' else "Rust"
        total = status[lang]['total']
        completed = status[lang]['completed']
        in_progress = status[lang]['in_progress']
        not_started = total - completed - in_progress
        
        print(f"\n{lang_name}言語:")
        print(f"  総プロジェクト数: {total}")
        print(f"  完了: {completed}")
        print(f"  進行中: {in_progress}")
        print(f"  未開始: {not_started}")
        
        if total > 0:
            completion_rate = (completed / total) * 100
            print(f"  完了率: {completion_rate:.1f}%")
            
            # プログレスバー
            bar_length = 30
            filled = int(bar_length * completion_rate / 100)
            bar = "=" * filled + "-" * (bar_length - filled)
            print(f"  [{bar}] {completion_rate:.1f}%")

def show_stats():
    """統計情報を表示"""
    status = get_challenge_status()
    
    print("統計情報")
    print("=" * 30)
    
    total_projects = status['go']['total'] + status['rust']['total']
    total_completed = status['go']['completed'] + status['rust']['completed']
    
    print(f"総プロジェクト数: {total_projects}")
    print(f"完了プロジェクト数: {total_completed}")
    print(f"全体完了率: {(total_completed/total_projects*100):.1f}%" if total_projects > 0 else "全体完了率: 0%")
    
    # 言語別完了率
    for lang in ['go', 'rust']:
        lang_name = "Go" if lang == 'go' else "Rust"
        total = status[lang]['total']
        completed = status[lang]['completed']
        if total > 0:
            rate = (completed / total) * 100
            print(f"{lang_name}完了率: {rate:.1f}%")

def mark_complete(challenge_num, language):
    """課題完了をマーク"""
    if language not in ['go', 'rust']:
        print("言語は 'go' または 'rust' を指定してください")
        return
    
    lang_name = "Go" if language == 'go' else "Rust"
    project_pattern = f"{language}-challenge-{challenge_num:03d}-"
    
    # 該当プロジェクトを検索
    lang_dir = Path(language)
    if not lang_dir.exists():
        print(f"{lang_name}ディレクトリが見つかりません")
        return
    
    found_project = None
    for project_dir in lang_dir.iterdir():
        if project_dir.is_dir() and project_dir.name.startswith(project_pattern):
            found_project = project_dir
            break
    
    if not found_project:
        print(f"課題 {challenge_num:03d} の{lang_name}プロジェクトが見つかりません")
        return
    
    # 完了マークファイルを作成
    completion_file = found_project / ".completed"
    completion_file.write_text(f"Completed at: {datetime.now().isoformat()}", encoding='utf-8')
    
    print(f"課題 {challenge_num:03d} の{lang_name}実装を完了としてマークしました")
    print(f"プロジェクト: {found_project}")

def main():
    if len(sys.argv) < 2:
        print("使用方法:")
        print("  python progress_manager.py status")
        print("  python progress_manager.py complete <課題番号> <言語>")
        print("  python progress_manager.py stats")
        sys.exit(1)
    
    command = sys.argv[1]
    
    if command == "status":
        show_status()
    elif command == "stats":
        show_stats()
    elif command == "complete":
        if len(sys.argv) != 4:
            print("使用方法: python progress_manager.py complete <課題番号> <言語>")
            print("例: python progress_manager.py complete 1 go")
            sys.exit(1)
        
        try:
            challenge_num = int(sys.argv[2])
            language = sys.argv[3]
            mark_complete(challenge_num, language)
        except ValueError:
            print("課題番号は数値で入力してください")
            sys.exit(1)
    else:
        print(f"不明なコマンド: {command}")
        sys.exit(1)

if __name__ == "__main__":
    main()
