import type { Metadata } from "next";
import type {
  GetPosts200,
  GetPostsId200,
} from "@/libs/api/generated/orval/model";
import { serverInstance } from "@/libs/api/server-instance";

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
        title: post.title,
        description: post.description,
      };
    }
  } catch (error) {
    console.error("Failed to fetch post for metadata:", error);
  }

  return {
    title: "Muslog - 記事の詳細ページ",
    description: "記事の詳細ページ",
  };
}

type PostLayoutParams = {
  params: Promise<{ id: string }>;
};

export default function PostLayout({
  children,
  params,
}: React.PropsWithChildren<PostLayoutParams>) {
  return <>{children}</>;
}
