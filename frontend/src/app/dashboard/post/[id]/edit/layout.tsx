import type { Metadata } from "next";
import type {
  GetPosts200,
  GetPostsId200,
} from "@/libs/api/generated/orval/model";
import { serverInstance } from "@/libs/api/server-instance";

// サーバーサイド用のgetPosts関数
const getPostsServer = (signal?: AbortSignal) => {
  return serverInstance<GetPosts200>({ url: `/posts`, method: "GET", signal });
};

export async function generateStaticParams() {
  try {
    const res = await getPostsServer();
    const posts = res.posts ?? [];

    if (posts.length === 0) {
      return [];
    }

    return posts.map((post) => ({
      id: post.id.toString(),
    }));
  } catch (error) {
    console.error("Failed to fetch posts for generateStaticParams:", error);
    return [];
  }
}

// サーバーサイド用のgetPostsId関数
const getPostsIdServer = (id: number, signal?: AbortSignal) => {
  return serverInstance<GetPostsId200>({
    url: `/posts/${id}`,
    method: "GET",
    signal,
  });
};

export async function generateMetadata({
  params,
}: {
  params: Promise<{ id: string }>;
}): Promise<Metadata> {
  const { id: idString } = await params;
  const id = Number(idString);
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
    title: "Muslog - 記事が見つかりません",
    description: "指定された記事は見つかりませんでした。",
  };
}

type EditPostLayoutParams = {
  params: Promise<{ id: string }>;
};

export default function EditPostLayout({
  children,
  params,
}: React.PropsWithChildren<EditPostLayoutParams>) {
  return <>{children}</>;
}
