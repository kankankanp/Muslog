"use client";

import { useEffect, useRef, useState, useCallback } from "react";

interface FullscreenWysiwygEditorProps {
  value: string;
  onChange: (value: string) => void;
  title: string;
  onTitleChange: (title: string) => void;
  headerImageUrl?: string;
}

export default function FullscreenWysiwygEditor({
  value,
  onChange,
  title,
  onTitleChange,
  headerImageUrl,
}: FullscreenWysiwygEditorProps) {
  const [zoom, setZoom] = useState(1.0);
  const editorRef = useRef<HTMLDivElement>(null);
  const [isUpdating, setIsUpdating] = useState(false);

  const handleZoom = (type: "in" | "out" | "reset") => {
    const step = 0.1;
    const minZoom = 0.5;
    const maxZoom = 2.0;

    setZoom((prev) => {
      if (type === "reset") return 1.0;
      const newZoom = type === "in" ? prev + step : prev - step;
      return Math.max(minZoom, Math.min(maxZoom, newZoom));
    });
  };

  // カーソル位置を保存する関数
  const saveCaretPosition = useCallback((element: HTMLElement) => {
    const selection = window.getSelection();
    if (!selection?.rangeCount) return null;

    const range = selection.getRangeAt(0);
    const preCaretRange = range.cloneRange();
    preCaretRange.selectNodeContents(element);
    preCaretRange.setEnd(range.endContainer, range.endOffset);
    
    return preCaretRange.toString().length;
  }, []);

  // カーソル位置を復元する関数
  const restoreCaretPosition = useCallback((element: HTMLElement, position: number) => {
    const selection = window.getSelection();
    if (!selection) return;

    let charIndex = 0;
    const walker = document.createTreeWalker(
      element,
      NodeFilter.SHOW_TEXT
    );

    let node;
    while (node = walker.nextNode()) {
      const nodeLength = node.textContent?.length || 0;
      if (charIndex + nodeLength >= position) {
        const range = document.createRange();
        range.setStart(node, position - charIndex);
        range.collapse(true);
        selection.removeAllRanges();
        selection.addRange(range);
        return;
      }
      charIndex += nodeLength;
    }

    // 位置が見つからない場合は最後に配置
    const range = document.createRange();
    range.selectNodeContents(element);
    range.collapse(false);
    selection.removeAllRanges();
    selection.addRange(range);
  }, []);

  // プレーンテキストを取得
  const getPlainText = useCallback((element: HTMLElement): string => {
    // HTMLから生のテキストを抽出（改行も保持）
    const textContent = element.innerText || element.textContent || '';
    return textContent;
  }, []);

  // Markdownテキストをスタイル付きHTMLに変換
  const renderMarkdownAsHTML = useCallback((text: string): string => {
    // エスケープ処理
    let html = text
      .replace(/&/g, '&amp;')
      .replace(/</g, '&lt;')
      .replace(/>/g, '&gt;');

    // 行ごとに処理
    const lines = html.split('\n');
    const processedLines = lines.map(line => {
      let processedLine = line;

      // ヘッダー（行の先頭から）
      if (/^### /.test(processedLine)) {
        processedLine = processedLine.replace(
          /^(### )(.*)$/,
          '<div class="header-3"><span class="md-syntax">$1</span><span class="header-content">$2</span></div>'
        );
      } else if (/^## /.test(processedLine)) {
        processedLine = processedLine.replace(
          /^(## )(.*)$/,
          '<div class="header-2"><span class="md-syntax">$1</span><span class="header-content">$2</span></div>'
        );
      } else if (/^# /.test(processedLine)) {
        processedLine = processedLine.replace(
          /^(# )(.*)$/,
          '<div class="header-1"><span class="md-syntax">$1</span><span class="header-content">$2</span></div>'
        );
      } else {
        // 通常の行の場合、インライン要素を処理
        
        // 太字
        processedLine = processedLine.replace(
          /(\*\*)([^*]+)(\*\*)/g,
          '<span class="md-syntax">$1</span><strong>$2</strong><span class="md-syntax">$3</span>'
        );
        
        // イタリック（太字と重複しないように）
        processedLine = processedLine.replace(
          /(?<!\*)(\*)([^*]+)(\*)(?!\*)/g,
          '<span class="md-syntax">$1</span><em>$2</em><span class="md-syntax">$3</span>'
        );
        
        // インラインコード
        processedLine = processedLine.replace(
          /(`)([^`]+)(`)/g,
          '<span class="md-syntax">$1</span><code>$2</code><span class="md-syntax">$3</span>'
        );
        
        // リンク
        processedLine = processedLine.replace(
          /(\[)([^\]]+)(\])(\()([^)]+)(\))/g,
          '<span class="md-syntax">$1</span><span class="link-text">$2</span><span class="md-syntax">$3$4</span><span class="link-url">$5</span><span class="md-syntax">$6</span>'
        );
        
        // 空行でない場合はdivで囲む
        if (processedLine.trim()) {
          processedLine = `<div>${processedLine}</div>`;
        } else {
          processedLine = '<div><br></div>';
        }
      }

      return processedLine;
    });

    return processedLines.join('');
  }, []);

  // 入力イベントのハンドリング
  const handleInput = useCallback(() => {
    if (!editorRef.current || isUpdating) return;

    const element = editorRef.current;
    const caretPosition = saveCaretPosition(element);
    const plainText = getPlainText(element);
    
    // 外部への変更通知
    onChange(plainText);

    // HTMLを更新（次のフレームで実行）
    setTimeout(() => {
      if (!isUpdating && editorRef.current) {
        setIsUpdating(true);
        const newHTML = renderMarkdownAsHTML(plainText);
        
        if (editorRef.current.innerHTML !== newHTML) {
          editorRef.current.innerHTML = newHTML;
          
          // カーソル位置を復元
          if (caretPosition !== null) {
            restoreCaretPosition(editorRef.current, caretPosition);
          }
        }
        
        setIsUpdating(false);
      }
    }, 0);
  }, [isUpdating, saveCaretPosition, getPlainText, onChange, renderMarkdownAsHTML, restoreCaretPosition]);

  // キーボードイベントのハンドリング
  const handleKeyDown = useCallback((e: React.KeyboardEvent<HTMLDivElement>) => {
    // Enterキーの処理
    if (e.key === 'Enter') {
      e.preventDefault();
      
      const selection = window.getSelection();
      if (!selection?.rangeCount) return;

      const range = selection.getRangeAt(0);
      
      // 新しい段落を作成
      const newDiv = document.createElement('div');
      const br = document.createElement('br');
      newDiv.appendChild(br);
      
      // 現在の位置に挿入
      range.deleteContents();
      range.insertNode(newDiv);
      
      // カーソルを新しい段落に移動
      range.setStart(newDiv, 0);
      range.collapse(true);
      selection.removeAllRanges();
      selection.addRange(range);
      
      // 入力イベントを手動でトリガー
      handleInput();
    }
  }, [handleInput]);

  // 初期値の設定と外部からの値変更への対応
  useEffect(() => {
    if (!editorRef.current || isUpdating) return;

    const currentText = getPlainText(editorRef.current);
    if (value !== currentText) {
      setIsUpdating(true);
      const newHTML = renderMarkdownAsHTML(value);
      editorRef.current.innerHTML = newHTML;
      setIsUpdating(false);
    }
  }, [value, isUpdating, getPlainText, renderMarkdownAsHTML]);

  return (
    <div className="fullscreen-wysiwyg h-screen bg-white flex flex-col">
      {/* ヘッダー部分 */}
      <div className="px-12 py-8 border-b border-gray-100 max-w-4xl mx-auto w-full">
        {/* ヘッダー画像 */}
        {headerImageUrl && (
          <div className="relative w-full h-64 mb-8 rounded-lg overflow-hidden shadow-sm">
            <img
              src={headerImageUrl}
              alt="Header Image"
              className="w-full h-full object-cover"
            />
          </div>
        )}
        
        {/* タイトル入力 */}
        <input
          type="text"
          placeholder="タイトルを入力"
          className="text-5xl font-bold mb-6 bg-transparent outline-none w-full placeholder-gray-400 leading-tight"
          value={title}
          onChange={(e) => onTitleChange(e.target.value)}
          style={{ fontSize: `${zoom * 48}px`, lineHeight: '1.2' }}
        />
        
        {/* ズームコントロール */}
        <div className="flex gap-2 justify-end opacity-70 hover:opacity-100 transition-opacity">
          <button
            className="px-3 py-1 bg-gray-100 hover:bg-gray-200 rounded text-sm transition-colors"
            onClick={() => handleZoom("out")}
            title="文字を小さく"
          >
            A-
          </button>
          <button
            className="px-3 py-1 bg-gray-100 hover:bg-gray-200 rounded text-sm transition-colors"
            onClick={() => handleZoom("in")}
            title="文字を大きく"
          >
            A+
          </button>
          <button
            className="px-3 py-1 bg-gray-100 hover:bg-gray-200 rounded text-sm transition-colors"
            onClick={() => handleZoom("reset")}
            title="サイズをリセット"
          >
            リセット
          </button>
        </div>
      </div>

      {/* エディタ部分 */}
      <div className="flex-1 overflow-auto max-w-4xl mx-auto w-full">
        <div
          ref={editorRef}
          contentEditable
          suppressContentEditableWarning
          onInput={handleInput}
          onKeyDown={handleKeyDown}
          className="min-h-full p-12 outline-none wysiwyg-editor focus:outline-none"
          style={{
            fontSize: `${zoom * 18}px`,
            lineHeight: '1.8',
            fontFamily: '-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif',
          }}
          data-placeholder="記事の本文を書いてください..."
        />
      </div>

      <style jsx>{`
        .wysiwyg-editor:empty::before {
          content: attr(data-placeholder);
          color: #9ca3af;
          font-style: italic;
        }
        
        .wysiwyg-editor {
          caret-color: #007acc;
        }
        
        .wysiwyg-editor :global(.header-1) {
          font-size: ${zoom * 32}px;
          font-weight: bold;
          color: #1a1a1a;
          margin: 20px 0 10px 0;
          line-height: 1.2;
        }
        
        .wysiwyg-editor :global(.header-1 .header-content) {
          font-size: ${zoom * 32}px;
          font-weight: bold;
          color: #1a1a1a;
        }
        
        .wysiwyg-editor :global(.header-2) {
          font-size: ${zoom * 26}px;
          font-weight: bold;
          color: #1a1a1a;
          margin: 18px 0 8px 0;
          line-height: 1.3;
        }
        
        .wysiwyg-editor :global(.header-2 .header-content) {
          font-size: ${zoom * 26}px;
          font-weight: bold;
          color: #1a1a1a;
        }
        
        .wysiwyg-editor :global(.header-3) {
          font-size: ${zoom * 22}px;
          font-weight: bold;
          color: #1a1a1a;
          margin: 16px 0 6px 0;
          line-height: 1.4;
        }
        
        .wysiwyg-editor :global(.header-3 .header-content) {
          font-size: ${zoom * 22}px;
          font-weight: bold;
          color: #1a1a1a;
        }
        
        .wysiwyg-editor :global(.md-syntax) {
          color: #9ca3af;
          font-weight: normal;
          font-style: normal;
        }
        
        .wysiwyg-editor :global(strong) {
          font-weight: bold;
          color: #1a1a1a;
        }
        
        .wysiwyg-editor :global(em) {
          font-style: italic;
          color: #374151;
        }
        
        .wysiwyg-editor :global(code) {
          background-color: #f3f4f6;
          padding: 2px 6px;
          border-radius: 4px;
          font-family: 'SF Mono', 'Monaco', 'Inconsolata', 'Roboto Mono', 'Consolas', monospace;
          font-size: ${zoom * 16}px;
          color: #ef4444;
        }
        
        .wysiwyg-editor :global(.link-text) {
          color: #2563eb;
          text-decoration: underline;
        }
        
        .wysiwyg-editor :global(.link-url) {
          color: #9ca3af;
          text-decoration: none;
        }
        
        .wysiwyg-editor :global(div) {
          margin: 0;
          padding: 0;
        }
        
        .wysiwyg-editor :global(div:empty) {
          min-height: 1.8em;
        }
      `}</style>
    </div>
  );
}