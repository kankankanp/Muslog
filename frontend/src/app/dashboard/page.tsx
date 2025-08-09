"use client";

import { useQueryClient } from "@tanstack/react-query";
import { useSelector } from "react-redux";
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

  if (isPending) return <div>Loading...</div>;
  if (error || !posts) return <div>Error loading posts.</div>;

  return (
    <div className="dark:bg-gray-900 bg-gray-100 min-h-screen">
      <h1 className="text-2xl font-bold text-center py-8 text-gray-900 dark:text-gray-100">
        全記事
      </h1>
      <BlogCard posts={posts.posts || []} onLikeToggle={handleLikeToggle} />
    </div>
  );
}
