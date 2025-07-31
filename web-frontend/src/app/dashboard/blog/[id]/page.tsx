"use client";

import { useParams } from "next/navigation";
import { useSelector } from "react-redux";
import { useGetBlogsId } from "@/app/libs/api/generated/orval/blogs/blogs";
import { useGetPostsPostIDLiked, usePostPostsPostIDLike, useDeletePostsPostIDUnlike } from "@/app/libs/api/generated/orval/likes/likes";
import { RootState } from "@/app/libs/store/store";
import BlogCard from "@/app/components/elements/cards/BlogCard";

export default function Page() {
  const params = useParams();
  const { id } = params as { id: string };
  const user = useSelector((state: RootState) => state.auth.user);

  const { data: post, isPending: isPostPending, error: postError } = useGetBlogsId(Number(id));
  const { data: likeStatusData, isPending: isLikeStatusPending, error: likeStatusError } = useGetPostsPostIDLiked(Number(id), { query: { enabled: !!user } });

  const { mutate: likePost } = usePostPostsPostIDLike();
  const { mutate: unlikePost } = useDeletePostsPostIDUnlike();

  const handleLikeClick = async () => {
    if (!user) {
      alert("ログインしてください");
      return;
    }

    try {
      if (likeStatusData?.isLiked) {
        unlikePost({ postID: Number(id) });
      } else {
        likePost({ postID: Number(id) });
      }
    } catch (error: any) {
      console.error("Failed to update like status:", error);
    }
  };

  if (isPostPending || isLikeStatusPending) return <div>Loading...</div>;
  if (postError || likeStatusError || !post) return <div>Error loading post.</div>;

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
