"use client";

import { useGetPosts } from "@/libs/api/generated/orval/posts/posts";

export default function Page() {
  const { data: posts, isPending, error } = useGetPosts();
  return (
    <div className="dark:bg-gray-900 bg-gray-100 min-h-screen">
      <h1>全記事</h1>
      <div className="py-8 px-4 max-w-6xl mx-auto">
        <div className="mt-8">
          {posts?.posts?.map((post) => (
            <div
              key={post.id}
              className="bg-white dark:bg-gray-800 rounded-lg shadow-md hover:shadow-lg transition-shadow duration-300 p-6 mb-4"
            >
              <h2 className="text-xl font-semibold text-gray-900 dark:text-gray-100">
                {post.title}
              </h2>
              <p className="text-gray-600 dark:text-gray-300">
                {post.description}
              </p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
