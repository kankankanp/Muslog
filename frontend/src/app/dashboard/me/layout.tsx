import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "Muslog - マイページ",
  description: "あなたの投稿とプロフィールを管理します。",
};

export default function MeLayout({ children }: { children: React.ReactNode }) {
  return <>{children}</>;
}
