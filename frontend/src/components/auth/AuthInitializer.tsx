"use client";

import { useEffect } from "react";
import { useDispatch } from "react-redux";
import { initializeAuth, expireAuth } from "@/libs/store/authSlice";

export default function AuthInitializer({
  children,
}: {
  children: React.ReactNode;
}) {
  const dispatch = useDispatch();

  useEffect(() => {
    const initAuth = () => {
      if (typeof window === "undefined") {
        dispatch(initializeAuth({ user: null }));
        return;
      }

      try {
        const savedAuth = localStorage.getItem("auth");
        if (!savedAuth) {
          dispatch(initializeAuth({ user: null }));
          return;
        }

        const authData = JSON.parse(savedAuth);
        const now = Date.now();

        // セキュリティチェック1: 保存から一定時間経過している場合は無効化
        const maxAge = 7 * 24 * 60 * 60 * 1000; // 7日間
        if (authData.timestamp && now - authData.timestamp > maxAge) {
          console.log("認証情報が古すぎるため無効化");
          localStorage.removeItem("auth");
          dispatch(initializeAuth({ user: null }));
          return;
        }

        // セキュリティチェック2: トークンの有効期限チェック
        if (authData.tokenExpiry && now > authData.tokenExpiry) {
          console.log("トークンの有効期限切れ");
          localStorage.removeItem("auth");
          dispatch(expireAuth());
          return;
        }

        // セキュリティチェック3: 必要なデータの存在確認
        if (!authData.user || !authData.isAuthenticated) {
          localStorage.removeItem("auth");
          dispatch(initializeAuth({ user: null }));
          return;
        }

        // 有効な認証情報として復元
        dispatch(
          initializeAuth({
            user: authData.user,
            tokenExpiry: authData.tokenExpiry,
          }),
        );
      } catch (error) {
        console.error("認証情報の復元エラー:", error);
        localStorage.removeItem("auth");
        dispatch(initializeAuth({ user: null }));
      }
    };

    initAuth();
  }, [dispatch]);

  return <>{children}</>;
}
