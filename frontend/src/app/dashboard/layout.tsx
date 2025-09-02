import type { Metadata } from "next";
import Header from "@/components/layouts/Header";
import Sidebar from "@/components/layouts/Sidebar";

export const metadata: Metadata = {
  title: "Muslog - ダッシュボード",
  description: "Muslogのダッシュボードです。全記事を閲覧できます。",
};

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <>
      <Header />
      <div className="flex h-screen bg-gray-100 dark:bg-gray-900">
        <Sidebar />
        <main className="flex-1 ml-64 p-8 overflow-y-auto">{children}</main>
      </div>
    </>
  );
}
