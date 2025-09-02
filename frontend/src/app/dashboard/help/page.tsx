// app/help/page.tsx または pages/help.tsx
"use client";

import { Search, ChevronDown, ChevronUp } from "lucide-react";
import React, { useState } from "react";

const HelpPage = () => {
  const faqs = [
    {
      question: "パスワードを忘れてしまったら？",
      answer:
        "パスワードをリセットするには、登録されたメールアドレスを入力してください。",
    },
    {
      question: "アカウントを削除するには？",
      answer: "設定ページからアカウント削除を選択できます。",
    },
    {
      question: "利用料金はかかりますか？",
      answer:
        "基本的な機能は無料で利用できますが、一部のプレミアム機能は有料です。",
    },
    {
      question: "利用規約はどこで確認できますか？",
      answer: "利用規約はフッターにリンクがあります。",
    },
  ];

  const [openIndex, setOpenIndex] = useState(null);

  const toggleFAQ = (index) => {
    setOpenIndex(openIndex === index ? null : index);
  };

  return (
    <div className="bg-gray-100 min-h-screen py-10 px-4">
      <div className="container mx-auto max-w-3xl space-y-8">
        {/* サービスガイドセクション */}
        <div className="bg-blue-700 text-white p-6 rounded-xl shadow-lg text-center">
          <h2 className="text-xl font-bold">サービスガイドはこちら</h2>
          <p className="mt-1 text-sm text-blue-100">
            サービス利用のためのユーザーガイド
          </p>
        </div>

        {/* 検索バーセクション */}
        <div className="bg-white p-6 rounded-xl shadow-lg text-center">
          <h3 className="text-lg font-semibold mb-4">お困りですか？</h3>
          <div className="relative mx-auto max-w-lg">
            <input
              type="text"
              placeholder=""
              className="w-full pl-10 pr-4 py-2 rounded-full border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
            <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-500" />
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
                {openIndex === index && (
                  <div className="py-4 text-gray-600 bg-gray-50 px-4">
                    <p>{faq.answer}</p>
                  </div>
                )}
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};

export default HelpPage;
