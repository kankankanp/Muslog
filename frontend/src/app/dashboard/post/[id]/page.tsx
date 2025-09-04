"use client";

import { useQueryClient } from "@tanstack/react-query";
import { useParams } from "next/navigation";
import { useSelector } from "react-redux";
import { usePostPostsPostIDLike, useDeletePostsPostIDUnlike } from "@/libs/api/generated/orval/likes/likes";
import { useGetPostsId } from "@/libs/api/generated/orval/posts/posts";
import { RootState } from "@/libs/store/store";
import Image from "next/image";
import ReactMarkdown from "react-markdown";

export default function Page() {
  const params = useParams();
  const { id } = params as { id: string };
  const user = useSelector((state: RootState) => state.auth.user);
  const queryClient = useQueryClient();

  const {
    data: postData,
    isLoading,
    error,
  } = useGetPostsId(Number(id), {
    query: { enabled: !!id },
  });

  const { mutate: likePost } = usePostPostsPostIDLike({
    mutation: {
      onSuccess: () => {
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
        queryClient.invalidateQueries({ queryKey: [`/posts/${id}`] });
      },
      onError: (error) => {
        console.error("Failed to unlike post:", error);
        alert("いいね解除に失敗しました");
      },
    },
  });

  const post = postData?.post;

  if (isLoading) {
    return (
      <div className="dark:bg-gray-900 bg-gray-100 min-h-screen flex items-center justify-center">
        <p className="text-gray-600 dark:text-gray-300">読み込み中...</p>
      </div>
    );
  }

  if (error || !post) {
    return (
      <div className="dark:bg-gray-900 bg-gray-100 min-h-screen flex items-center justify-center">
        <div className="p-6 bg-white dark:bg-gray-800 rounded-lg shadow-md text-center">
          <h2 className="text-xl font-semibold text-gray-900 dark:text-gray-100 mb-4">
            投稿が見つかりません
          </h2>
          <p className="text-gray-600 dark:text-gray-300">
            この投稿は存在しないか、削除された可能性があります。
          </p>
        </div>
      </div>
    );
  }

  const handleLikeToggle = () => {
    if (!user) {
      alert("ログインしていいねしてください");
      return;
    }
    if (post.isLiked) {
      unlikePost({ postID: post.id });
    } else {
      likePost({ postID: post.id });
    }
  };

  return (
    <div className="dark:bg-gray-900 bg-gray-100 min-h-screen p-6">
      <main className="container mx-auto">
        <div className="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
          <Image
            src="/default-image.jpg" // Placeholder header image
            alt="Header Image"
            width={800} // Adjust width as needed
            height={400} // Adjust height as needed
            className="w-full h-64 object-cover rounded-md mb-6"
          />
          <h1 className="text-3xl font-bold text-gray-900 dark:text-gray-100 mb-4">
            {post.title}
          </h1>

          {post.tracks && post.tracks.length > 0 && (
            <div className="mb-6">
              <h2 className="text-xl font-semibold text-gray-900 dark:text-gray-100 mb-3">
                関連トラック
              </h2>
              <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
                {post.tracks.map((track, index) => (
                  <div
                    key={index}
                    className="bg-gray-50 dark:bg-gray-700 rounded-lg p-4 flex items-center space-x-4"
                  >
                    {track.albumImageUrl && (
                      <img
                        src={track.albumImageUrl}
                        alt={track.name}
                        className="w-16 h-16 rounded-md object-cover"
                      />
                    )}
                    <div>
                      <p className="font-medium text-gray-900 dark:text-gray-100">
                        {track.name}
                      </p>
                      <p className="text-sm text-gray-600 dark:text-gray-300">
                        {track.artistName}
                      </p>
                    </div>
                  </div>
                ))}
              </div>
            </div>
          )}

          {post.tags && post.tags.length > 0 && (
            <div className="flex flex-wrap gap-2 mb-6">
              {post.tags.map((tag) => (
                <span
                  key={tag.id}
                  className="bg-blue-100 text-blue-800 text-xs font-medium px-2.5 py-0.5 rounded dark:bg-blue-900 dark:text-blue-300"
                >
                  {tag.name}
                </span>
              ))}
            </div>
          )}

          <div className="prose prose-lg max-w-none w-full mt-6">
            <ReactMarkdown>{post.description}</ReactMarkdown>
          </div>

          <div className="flex items-center justify-between border-t border-gray-200 dark:border-gray-700 pt-4">
            <div className="flex items-center text-gray-500 dark:text-gray-400 text-sm">
              <button
                onClick={handleLikeToggle}
                className={`flex items-center space-x-1 focus:outline-none ${
                  post.isLiked ? "text-red-500" : "text-gray-400"
                }`}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-5 w-5"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                >
                  <path
                    fillRule="evenodd"
                    d="M3.172 5.172a4 4 0 015.656 0L10 6.343l1.172-1.171a4 4 0 115.656 5.656L10 17.657l-6.828-6.829a4 4 0 010-5.656z"
                    clipRule="evenodd"
                  />
                </svg>
                <span>{post.likesCount || 0} Likes</span>
              </button>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
}
