"use client";

import React, { useState } from "react";
import ReactMarkdown from "react-markdown";

export default function AddPostPage() {
  const [title, setTitle] = useState("");
  const [markdown, setMarkdown] = useState("");
  const [viewMode, setViewMode] = useState<'editor' | 'preview' | 'split'>('split');

  return (
    <>
      <h1 className="text-3xl font-bold border-gray-100 border-b-2 bg-white px-6 py-6">
        記事を作成する
      </h1>
      <div className="flex justify-center gap-4 mb-4 mt-4">
        <button
          className={`px-4 py-2 rounded ${viewMode === 'editor' ? 'bg-blue-500 text-white' : 'bg-gray-200'}`}
          onClick={() => setViewMode('editor')}
        >
          エディタ
        </button>
        <button
          className={`px-4 py-2 rounded ${viewMode === 'preview' ? 'bg-blue-500 text-white' : 'bg-gray-200'}`}
          onClick={() => setViewMode('preview')}
        >
          プレビュー
        </button>
        <button
          className={`px-4 py-2 rounded ${viewMode === 'split' ? 'bg-blue-500 text-white' : 'bg-gray-200'}`}
          onClick={() => setViewMode('split')}
        >
          分割
        </button>
      </div>
      <div className="flex h-screen bg-white">
        {/* 右側：プレビュー */}
        <div className={`p-8 overflow-y-auto ${viewMode === 'editor' ? 'hidden' : 'flex-1'} ${viewMode === 'split' ? 'w-1/2' : ''}`}>
          <div className="mb-6">
            <h2 className="text-3xl font-bold text-gray-400">
              {title || "記事タイトル"}
            </h2>
          </div>
          <div className="prose prose-lg max-w-none">
            <ReactMarkdown>
              {markdown || "プレビューがここに表示されます。"}
            </ReactMarkdown>
          </div>
        </div>
        {/* 左側：エディタ */}
        <div className={`p-8 flex flex-col gap-4 border-r border-gray-200 ${viewMode === 'preview' ? 'hidden' : 'flex-1'} ${viewMode === 'split' ? 'w-1/2' : ''}`}>
          <input
            type="text"
            placeholder="記事タイトル"
            className="text-3xl font-bold mb-4 bg-transparent outline-none"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
          />
          <button
            className="flex items-center gap-2 px-4 py-2 bg-gray-100 rounded w-fit mb-4"
            // 画像追加のロジックは後で
          >
            <span className="text-xl">＋</span> 画像を追加
          </button>
          <textarea
            className="flex-1 w-full border rounded p-4 resize-none bg-gray-50"
            placeholder="本文をマークダウンで入力してください"
            value={markdown}
            onChange={(e) => setMarkdown(e.target.value)}
          />
        </div>
      </div>
    </>
  );
}
