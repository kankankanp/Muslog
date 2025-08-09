import type { Metadata } from 'next';
import { serverInstance } from '@/libs/api/server-instance';
import type { GetPostsId200 } from '@/libs/api/generated/orval/model';

// サーバーサイド用のgetPostsId関数
const getPostsIdServer = (id: number, signal?: AbortSignal) => {
  return serverInstance<GetPostsId200>({
    url: `/posts/${id}`,
    method: "GET",
    signal,
  });
};

export async function generateMetadata({ params }: { params: { id: string } }): Promise<Metadata> {
  const id = Number(params.id);
  try {
    const response = await getPostsIdServer(id);
    const post = response.post;
    if (post) {
      return {
        title: `Muslog - ${post.title} を編集`,
        description: `記事「${post.title}」を編集します。`,
      };
    }
  } catch (error) {
    console.error("Failed to fetch post for metadata:", error);
  }

  return {
    title: 'Muslog - 記事が見つかりません',
    description: '指定された記事は見つかりませんでした。',
  };
}

export default function EditPostLayout({ children }: { children: React.ReactNode }) {
  return <>{children}</>;
}
