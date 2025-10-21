// app/help/page.tsx または pages/help.tsx
'use client';

import { ChevronDown, ChevronUp } from 'lucide-react';
import React, { useState } from 'react';
import ReactMarkdown from 'react-markdown';

const HelpPage = () => {
  const [openIndex, setOpenIndex] = useState<number | null>(null);

  // TODO：indexの型を修正する
  const toggleFAQ = (index: any) => {
    setOpenIndex(openIndex === index ? null : index);
  };
  const faqs = [
    {
      question: '利用料金はかかりますか？',
      answer: '全ての機能が無料で利用できます。',
    },
    {
      question: 'ゲストログインとは何ですか？',
      answer:
        'ゲストログインは、登録やログインをせずに一時的にサービスを利用できる機能です。ユーザー登録しなくても簡単に利用を開始できます。',
    },
    {
      question: 'タグの使い方を教えてください。',
      answer:
        '投稿やコミュニティに関連するキーワードをタグとして設定できます。タグを付けることで、他のユーザーが同じテーマの投稿を検索・発見しやすくなります。投稿作成や編集画面でタグを入力してください。',
    },
    {
      question: 'サポートへ問い合わせるには？',
      answer:
        'こちらのGoogleフォーム(https://forms.gle/xxxxxxxxxxxxxxxx)からお問合わせください。',
    },
  ];

  const [isGuideOpen, setIsGuideOpen] = useState(false);

  const toggleGuide = () => {
    setIsGuideOpen(!isGuideOpen);
  };

  return (
    <>
      <h1 className="text-3xl font-bold border-gray-100 border-b-2 bg-white px-6 py-6">
        ヘルプ
      </h1>
      <div className="py-10 px-4">
        <div className="bg-white p-6 rounded-xl shadow-lg mb-8">
          <button
            className="flex justify-between items-center w-full py-2 text-left font-medium text-gray-800 hover:bg-gray-50"
            onClick={toggleGuide}
          >
            <h3 className="text-lg font-bold text-center">ご利用ガイド</h3>
            {isGuideOpen ? (
              <ChevronUp className="h-5 w-5 text-gray-500" />
            ) : (
              <ChevronDown className="h-5 w-5 text-gray-500" />
            )}
          </button>
          <div
            className={`transition-all duration-200 ease-in-out ${
              isGuideOpen
                ? 'py-4 max-h-[80vh] overflow-auto'
                : 'max-h-0 overflow-hidden'
            } text-gray-600`}
          >
            <div className="prose max-w-none">
              <ReactMarkdown>
                {`
# Muslog 利用ガイド

Muslogへようこそ！このガイドでは、Muslogの主な機能と使い方について説明します。

## 1. 記事の作成と投稿

### 記事の作成
ダッシュボードの「記事を作成する」ボタンから新しい記事を作成できます。
*   **タイトル:** 記事のタイトルを入力します。
*   **本文:** マークダウン形式で記事の本文を入力します。プレビュー機能でリアルタイムに表示を確認できます。
*   **ヘッダー画像:** 記事のトップに表示される画像をアップロードできます。
*   **投稿内画像:** 記事の本文中に画像を挿入できます。
*   **タグ:** 記事に関連するタグを追加できます。
*   **曲:** Spotifyから記事に関連する曲を検索し、追加できます。

### 記事の投稿
記事の作成が完了したら、「記事を投稿する」ボタンをクリックして公開します。

## 2. 記事の閲覧と検索

### ホーム画面
ホーム画面では、最新の投稿を閲覧できます。
上部の検索バーを使って、投稿のタイトルや本文に含まれるキーワードで記事を検索できます。

### 記事の詳細
各記事をクリックすると、詳細ページで記事の全文、関連する曲、タグなどを確認できます。

## 3. コミュニティ機能

### コミュニティの閲覧
コミュニティ画面では、様々なテーマのコミュニティを閲覧できます。
検索バーを使って、コミュニティの名前や説明文で検索できます。

### コミュニティの作成
「コミュニティを作成する」ボタンから、新しいコミュニティを作成できます。

## 4. マイページ

マイページでは、あなたのプロフィール情報、作成した記事、いいねした記事などを確認できます。
プロフィール画像を更新することも可能です。

## 5. その他の機能

*   **いいね機能:** 気に入った記事には「いいね」をすることができます。
*   **レスポンシブデザイン:** スマートフォンやタブレットなど、様々なデバイスで快適に利用できます。

ご不明な点がございましたら、「よくある質問」をご参照いただくか、お問い合わせください。
              `}
              </ReactMarkdown>
            </div>
          </div>
        </div>

        {/* よくある質問セクション */}
        <div className="bg-white p-6 rounded-xl shadow-lg">
          <h3 className="text-lg font-bold mb-4 text-center">よくある質問</h3>
          <div>
            {faqs.map((faq, index) => (
              <div key={index} className="border-b border-gray-200">
                <button
                  className="flex justify-between items-center w-full py-4 text-left font-medium text-gray-800 hover:bg-gray-50"
                  onClick={() => toggleFAQ(index)}
                >
                  <span>Q. {faq.question}</span>
                  {openIndex === index ? (
                    <ChevronUp className="h-5 w-5 text-gray-500" />
                  ) : (
                    <ChevronDown className="h-5 w-5 text-gray-500" />
                  )}
                </button>
                <div
                  className={`overflow-hidden transition-all duration-200 ease-in-out ${
                    openIndex === index ? 'max-h-screen py-4' : 'max-h-0'
                  } text-gray-600 bg-gray-50 px-4`}
                >
                  <p>{faq.answer}</p>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </>
  );
};

export default HelpPage;
