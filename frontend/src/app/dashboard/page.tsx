"use client";

import { useQueryClient } from "@tanstack/react-query";
import { Search } from "lucide-react";
import { useSelector } from "react-redux";
import Loading from "../loading";
import BlogCard from "@/components/elements/cards/BlogCard";
import {
  usePostPostsPostIDLike,
  useDeletePostsPostIDUnlike,
} from "@/libs/api/generated/orval/likes/likes";
import { useGetPosts } from "@/libs/api/generated/orval/posts/posts";
import { RootState } from "@/libs/store/store";

export default function Page() {
  const { data: posts, isPending, error } = useGetPosts();
  const user = useSelector((state: RootState) => state.auth.user);
  const queryClient = useQueryClient();

  const { mutate: likePost } = usePostPostsPostIDLike({
    mutation: {
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["/posts"] });
      },
      onError: (error) => {
        console.error("Failed to like post:", error);
        alert("いいねに失敗しました");
      },
    },
  });

  const { mutate: unlikePost } = useDeletePostsPostIDUnlike({
    mutation: {
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["/posts"] });
      },
      onError: (error) => {
        console.error("Failed to unlike post:", error);
        alert("いいね解除に失敗しました");
      },
    },
  });

  const handleLikeToggle = (postId: number, isCurrentlyLiked: boolean) => {
    if (!user) {
      alert("ログインしてください");
      return;
    }

    if (isCurrentlyLiked) {
      unlikePost({ postID: postId });
    } else {
      likePost({ postID: postId });
    }
  };

  if (isPending) return <Loading />;
  if (error || !posts) return <div>Error loading posts.</div>;

  return (
    <div className="dark:bg-gray-900 min-h-screen">
      <h1 className="border-gray-100 border-b-2 bg-white px-6 py-6 flex items-center justify-between">
        <div className="text-gray-800 text-3xl font-bold">ホーム</div>
        <div className="relative flex-grow mx-4 max-w-lg">
          <input
            type="text"
            placeholder="検索"
            className="w-full pl-4 pr-10 py-2 rounded-full text-gray-200 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-purple-500 border-gray-200 border-2"
          />
          <Search className="absolute right-3 top-1/2 -translate-y-1/2 h-5 w-5 text-gray-400" />
        </div>
      </h1>
      <BlogCard posts={posts.posts || []} onLikeToggle={handleLikeToggle} />
    </div>
  );
}
