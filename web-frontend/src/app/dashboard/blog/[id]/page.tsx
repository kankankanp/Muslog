"use client";

import { useParams } from "next/navigation";
import { useEffect, useState } from "react";
import { BlogsService, LikesService } from "@/app/libs/api/generated";
import BlogCard from "@/app/components/elements/cards/BlogCard";
import { useUser } from "@/app/libs/hooks/useUser";

export default function Page() {
  const params = useParams();
  const { id } = params as { id: string };
  const [post, setPost] = useState<any>(null);
  const [isLiked, setIsLiked] = useState<boolean>(false);
  const { user } = useUser();

  useEffect(() => {
    const fetchPostAndLikeStatus = async () => {
      try {
        const postResponse = await BlogsService.getBlogs1(Number(id));
        setPost(postResponse.post);

        if (user) {
          const likeStatusResponse = await LikesService.getPostsLiked(
            Number(id)
          );
          setIsLiked(likeStatusResponse.isLiked ?? false);
        }
      } catch (error) {
        console.error("Failed to fetch post or like status:", error);
      }
    };

    fetchPostAndLikeStatus();
  }, [id, user]);

  const handleLikeClick = async () => {
    if (!user) {
      alert("ログインしてください");
      return;
    }

    try {
      if (isLiked) {
        await LikesService.deletePostsUnlike(Number(id));
        setPost((prevPost: any) => ({
          ...prevPost,
          likesCount: prevPost.likesCount - 1,
        }));
      } else {
        await LikesService.postPostsLike(Number(id));
        setPost((prevPost: any) => ({
          ...prevPost,
          likesCount: prevPost.likesCount + 1,
        }));
      }
      setIsLiked(!isLiked);
    } catch (error) {
      console.error("Failed to update like status:", error);
    }
  };

  if (!post) {
    return (
      <div className="dark:bg-gray-900 bg-gray-100 h-screen flex items-center justify-center">
        <h1 className="text-2xl font-bold text-red-500">
          記事が見つかりません
        </h1>
      </div>
    );
  }

  return (
    <div className="dark:bg-gray-900 bg-gray-100">
      <main>
        <BlogCard
          isDetailPage={true}
          posts={[post]}
          onLikeClick={handleLikeClick}
          isLiked={isLiked}
        />
      </main>
    </div>
  );
}
