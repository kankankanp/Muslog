import type { Metadata } from "next";
import DashboardContainer from "@/components/layouts/DashboardContainer";

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
    <DashboardContainer>
      {children}
    </DashboardContainer>
  );
}