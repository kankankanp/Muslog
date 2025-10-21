'use client';

import { expireAuth } from '../store/authSlice';
import { store } from '../store/store';

// API エラーハンドリング用の関数
export const handleApiError = (error: any) => {
  // 401 Unauthorized または 403 Forbidden の場合は認証無効化
  if (error?.response?.status === 401 || error?.response?.status === 403) {
    console.log('認証エラーを検知: トークンを無効化します');
    store.dispatch(expireAuth());

    // ログイン画面にリダイレクト
    if (typeof window !== 'undefined') {
      const currentPath = window.location.pathname;
      if (!currentPath.includes('/login-or-signup')) {
        window.location.href = `/login-or-signup?returnUrl=${encodeURIComponent(currentPath)}`;
      }
    }
  }

  return Promise.reject(error);
};

// カスタムインスタンスにインターセプターを追加する関数
export const setupApiInterceptors = (instance: any) => {
  // レスポンスインターセプター
  instance.interceptors.response.use(
    (response: any) => response,
    handleApiError
  );

  return instance;
};
