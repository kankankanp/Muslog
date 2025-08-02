"use client";

import { useQueryClient } from "@tanstack/react-query";
import { useParams } from "next/navigation";
import { useSelector } from "react-redux";
import BlogCard from "@/components/elements/cards/BlogCard";
import {
  useGetPostsPostIDLiked,
  usePostPostsPostIDLike,
  useDeletePostsPostIDUnlike,
} from "@/libs/api/generated/orval/likes/likes";
import { useGetPostsId } from "@/libs/api/generated/orval/posts/posts";
import { RootState } from "@/libs/store/store";

export default function Page() {
  const params = useParams();
  const { id } = params as { id: string };
  const user = useSelector((state: RootState) => state.auth.user);
  const queryClient = useQueryClient();

  const {
    data: post,
    isPending: isPostPending,
    error: postError,
  } = useGetPostsId(Number(id));
  const {
    data: likeStatusData,
    isPending: isLikeStatusPending,
    error: likeStatusError,
  } = useGetPostsPostIDLiked(Number(id), { query: { enabled: !!user } });

  const { mutate: likePost } = usePostPostsPostIDLike({
    mutation: {
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: [`/posts/${id}/liked`] });
        queryClient.invalidateQueries({ queryKey: [`/posts/${id}`] });
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
        queryClient.invalidateQueries({ queryKey: [`/posts/${id}/liked`] });
        queryClient.invalidateQueries({ queryKey: [`/posts/${id}`] });
      },
      onError: (error) => {
        console.error("Failed to unlike post:", error);
        alert("いいね解除に失敗しました");
      },
    },
  });

  const handleLikeClick = async () => {
    if (!user) {
      alert("ログインしてください");
      return;
    }

    if (likeStatusData?.isLiked) {
      unlikePost({ postID: Number(id) });
    } else {
      likePost({ postID: Number(id) });
    }
  };

  if (isPostPending || isLikeStatusPending) return <div>Loading...</div>;
  if (postError || likeStatusError || !post)
    return <div>Error loading post.</div>;

  return (
    <div className="dark:bg-gray-900 bg-gray-100">
      <main>
        <BlogCard
          isDetailPage={true}
          posts={post.post ? [post.post] : []}
          onLikeClick={handleLikeClick}
          isLiked={likeStatusData?.isLiked ?? false}
        />
      </main>
    </div>
  );
}
