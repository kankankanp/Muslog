// components/AuthContainer.tsx
"use client";

import { useState } from "react";
import LoginForm from "@/components/elements/forms/LoginForm";
import SignupForm from "@/components/elements/forms/SignupForm";

const Page = () => {
  const [isLoginForm, setIsLoginForm] = useState<boolean>(true);

  return (
    <div className="flex w-full min-h-screen bg-white shadow-xl overflow-hidden">
      {/* 左側のデザイン（768px以下は非表示） */}
      <div className="w-2/5 bg-indigo-700 text-white flex-col justify-center items-center p-8 hidden md:flex">
        <div className="w-full text-center">
          <h1 className="text-3xl font-bold mb-2">
            さあ、<span className="inline-block">音楽を始めよう。</span>
          </h1>
        </div>
      </div>

      {/* 右側のフォームコンテナ */}
      <div className="w-full md:w-3/5 p-12 bg-white flex flex-col justify-center items-center">
        <div className="border-gray-200 border-2 rounded-md p-[100px]">
          <div className="flex w-full max-w-sm mb-8 rounded-lg border border-gray-300 overflow-hidden">
            <button
              className={`flex-1 py-3 text-center transition-colors duration-200 ${
                isLoginForm
                  ? "bg-gray-800 text-white font-bold"
                  : "bg-white text-gray-700 hover:bg-gray-100"
              }`}
              onClick={() => setIsLoginForm(true)}
            >
              ログイン
            </button>
            <button
              className={`flex-1 py-3 text-center transition-colors duration-200 ${
                !isLoginForm
                  ? "bg-gray-800 text-white font-bold"
                  : "bg-white text-gray-700 hover:bg-gray-100"
              }`}
              onClick={() => setIsLoginForm(false)}
            >
              新規登録
            </button>
          </div>
          {isLoginForm ? <LoginForm /> : <SignupForm />}
        </div>
      </div>
    </div>
  );
};

export default Page;
