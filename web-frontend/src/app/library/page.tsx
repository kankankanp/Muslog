"use client";

import {
  Key,
  ReactElement,
  JSXElementConstructor,
  ReactNode,
  ReactPortal,
  AwaitedReactNode,
} from "react";
import { useAllBlogs } from "@/app/libs/blog";

export default function TestPage() {
  const { data: blogs, isLoading, error } = useAllBlogs();

  if (isLoading) return <p>読み込み中...</p>;
  if (error) return <p>データの取得に失敗しました。</p>;
  console.log(error);

  return (
    <div className="dark:bg-gray-900 bg-gray-100 min-h-screen py-8 px-4 max-w-6xl mx-auto">
      <h1 className="text-3xl font-bold text-gray-800 dark:text-white mb-6">
        全ブログ一覧
      </h1>

      {blogs && blogs.length > 0 ? (
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          {blogs.map(
            (blog: {
              id: Key | null | undefined;
              title:
                | string
                | number
                | bigint
                | boolean
                | ReactElement<any, string | JSXElementConstructor<any>>
                | Iterable<ReactNode>
                | ReactPortal
                | Promise<AwaitedReactNode>
                | null
                | undefined;
              createdAt: string | number | Date;
              content:
                | string
                | number
                | bigint
                | boolean
                | ReactElement<any, string | JSXElementConstructor<any>>
                | Iterable<ReactNode>
                | ReactPortal
                | Promise<AwaitedReactNode>
                | null
                | undefined;
            }) => (
              <div
                key={blog.id}
                className="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6"
              >
                <h2 className="text-xl font-semibold text-gray-800 dark:text-white mb-2">
                  {blog.title}
                </h2>
                <p className="text-sm text-gray-500 dark:text-gray-400 mb-4">
                  {new Date(blog.createdAt).toLocaleDateString()}
                </p>
                <p className="text-gray-700 dark:text-gray-300 line-clamp-3">
                  {blog.content}
                </p>
              </div>
            )
          )}
        </div>
      ) : (
        <p className="text-gray-600 dark:text-gray-400">
          ブログ記事はまだありません。
        </p>
      )}
    </div>
  );
}
