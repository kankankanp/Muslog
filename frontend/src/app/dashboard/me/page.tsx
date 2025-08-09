"use client";

import { useGetMe } from "@/libs/api/generated/orval/auth/auth";
import { GetPosts200 } from "@/libs/api/generated/orval/model";
import { useGetUsersIdPosts } from "@/libs/api/generated/orval/users/users";

export default function ProfilePage() {
  const {
    data: currentUser,
    isPending: userLoading,
    error: userError,
  } = useGetMe();

  const {
    data: postsData,
    isPending: postsLoading,
    error: postsError,
  } = useGetUsersIdPosts<GetPosts200>(currentUser?.id || "");

  if (userLoading) {
    return (
      <div className="flex items-center justify-center min-h-screen bg-gray-100 dark:bg-gray-900">
        <div className="text-center">
          <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-500 mx-auto mb-4"></div>
          <p className="text-gray-600 dark:text-gray-300">
            Loading user info...
          </p>
        </div>
      </div>
    );
  }

  if (userError) {
    return (
      <div className="flex items-center justify-center min-h-screen bg-gray-100 dark:bg-gray-900">
        <div className="text-center">
          <p className="text-red-600 dark:text-red-400">
            Error loading user: {(userError as any)?.message || "Unknown error"}
          </p>
        </div>
      </div>
    );
  }

  if (!currentUser) {
    return (
      <div className="flex items-center justify-center min-h-screen bg-gray-100 dark:bg-gray-900">
        <div className="text-center">
          <p className="text-gray-600 dark:text-gray-300">User not found</p>
        </div>
      </div>
    );
  }

  const posts = postsData?.posts;

  return (
    <main className="dark:bg-gray-900 bg-gray-100 min-h-screen pt-8">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="bg-white dark:bg-gray-800 shadow-md rounded-lg p-6 mb-8">
          <div className="flex items-center space-x-4">
            <div className="w-16 h-16 bg-blue-500 rounded-full flex items-center justify-center">
              <span className="text-white text-xl font-bold">
                {currentUser.name}
              </span>
            </div>
            <div>
              <h1 className="text-2xl font-bold text-gray-900 dark:text-gray-100">
                {currentUser.name}
              </h1>
              <p className="text-gray-600 dark:text-gray-300">
                {currentUser.email}
              </p>
              <p className="text-sm text-gray-500 dark:text-gray-400">
                Registered:{" "}
                {currentUser.createdAt
                  ? new Date(currentUser.createdAt).toLocaleDateString()
                  : ""}
              </p>
            </div>
          </div>
        </div>

        <div className="mb-6">
          <h2 className="text-xl font-semibold text-gray-900 dark:text-gray-100">
            My Postsaaa
          </h2>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {posts?.map((post) => (
            <div
              key={post.id}
              className="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6"
            >
              <h3 className="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-2">
                {post.title}
              </h3>
              <p className="text-gray-600 dark:text-gray-300 mb-4 line-clamp-3">
                {post.description}
              </p>
              <div className="flex items-center text-gray-500 dark:text-gray-400 text-sm mb-4">
                <span>❤️ {post.likesCount || 0} Likes</span>
              </div>
              {post.tags && post.tags.length > 0 && (
                <div className="flex flex-wrap gap-2 mb-4">
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
              <div className="flex gap-2">
                <a
                  href={`/dashboard/post/${post.id}/edit`}
                  className="inline-block bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-md text-sm font-medium transition duration-200"
                >
                  編集
                </a>
                <a
                  href={`/dashboard/post/${post.id}`}
                  className="inline-block bg-gray-500 hover:bg-gray-600 text-white px-4 py-2 rounded-md text-sm font-medium transition duration-200"
                >
                  表示
                </a>
              </div>
            </div>
          ))}
        </div>
      </div>
    </main>
  );
}
