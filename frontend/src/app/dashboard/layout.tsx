'use client';

import dynamic from 'next/dynamic';
import Skeleton from '@/components/layouts/Skeleton';

// export const metadata: Metadata = {
//   title: "Muslog - ダッシュボード",
//   description: "Muslogのダッシュボードです。全記事を閲覧できます。",
// };

const DashboardContainer = dynamic(
  () => import('@/components/layouts/DashboardContainer'),
  {
    ssr: false,
    loading: () => <Skeleton />,
  }
);

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return <DashboardContainer>{children}</DashboardContainer>;
}
