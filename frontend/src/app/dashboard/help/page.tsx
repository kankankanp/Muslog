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

  const [openIndex, setOpenIndex] = useState<number | null>(null);

  // TODO：indexの型を修正する
  const toggleFAQ = (index: any) => {
    setOpenIndex(openIndex === index ? null : index);
  };

  return (
    <>
      <h1 className="text-3xl font-bold border-gray-100 border-b-2 bg-white px-6 py-6">
        ヘルプ
      </h1>
      <div className="min-h-screen py-10 px-4">
        <div className="container mx-auto max-w-3xl space-y-8">
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
                  <div
                    className={`overflow-hidden transition-all duration-200 ease-in-out ${
                      openIndex === index ? "max-h-screen py-4" : "max-h-0"
                    } text-gray-600 bg-gray-50 px-4`}
                  >
                    <p>{faq.answer}</p>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default HelpPage;
