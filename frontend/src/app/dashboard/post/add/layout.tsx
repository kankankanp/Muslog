import type { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'Muslog - 新しい記事を作成',
  description: '新しい音楽ブログ記事を作成します。',
};

export default function AddPostLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return <>{children}</>;
}
