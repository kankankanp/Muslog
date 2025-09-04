"use client";

import { faHeart } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Image from "next/image";
import { useState } from "react";
import { useGetMe } from "@/libs/api/generated/orval/auth/auth";
import { GetPosts200 } from "@/libs/api/generated/orval/model";
import { useGetPosts } from "@/libs/api/generated/orval/posts/posts"; // Import useGetPosts
import { useGetUsersIdPosts } from "@/libs/api/generated/orval/users/users";

export default function ProfilePage() {
  const {
    data: currentUser,
    isPending: userLoading,
    error: userError,
  } = useGetMe();
  console.log(currentUser);
  const {
    data: postsData,
    isPending: postsLoading,
    error: postsError,
  } = useGetUsersIdPosts<GetPosts200>(currentUser?.id || "");
  const {
    data: allPostsData,
    isPending: allPostsLoading,
    error: allPostsError,
  } = useGetPosts(); // Fetch all posts
  const [tab, setTab] = useState<"created" | "liked" | "community-history">(
    "created"
  );

  if (userLoading) {
    return (
      <div className="flex items-center justify-center min-h-screen bg-gray-100 dark:bg-gray-900">
        <div className="text-center">
          <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-500 mx-auto mb-4" />
          <p className="text-gray-600 dark:text-gray-300">
            ユーザー情報を読み込み中...
          </p>
        </div>
      </div>
    );
  }
  if (userError) {
    return (
      <div className="flex items-center justify-center min-h-screen bg-gray-100 dark:bg-gray-900">
        <p className="text-red-600 dark:text-red-400">
          ユーザーの取得に失敗しました：
          {(userError as any)?.message || "Unknown error"}
        </p>
      </div>
    );
  }
  if (!currentUser) {
    return (
      <div className="flex items-center justify-center min-h-screen bg-gray-100 dark:bg-gray-900">
        <p className="text-gray-600 dark:text-gray-300">
          ユーザーが見つかりません
        </p>
      </div>
    );
  }

  const posts = postsData?.posts ?? [];
  const likedPosts = allPostsData?.posts?.filter((post) => post.isLiked) ?? []; // Filter liked posts

  return (
    <>
      <h1 className="text-3xl font-bold border-gray-100 border-b-2 bg-white px-8 py-6">
        マイページ
      </h1>
      <section className="bg-white dark:bg-gray-800">
        <div className="max-w-6xl mx-auto px-8 lg:px-12 pt-10">
          <div className="flex gap-8">
            {/* 固定サイズの円アバター（ピル化を防ぐ） */}
            <div className="shrink-0 w-44 h-44 rounded-full bg-gray-200 dark:bg-gray-700 flex items-center justify-center">
              <span className="text-2xl text-gray-600 dark:text-gray-200"></span>
            </div>

            <div className="flex justify-center gap-1 flex-col">
              <h1 className="text-2xl font-semibold text-gray-900 dark:text-gray-100">
                {currentUser.name}
              </h1>
              <p>{currentUser.email}</p>
            </div>
          </div>
          <nav className="flex gap-10 pt-8 mx-auto">
            {[
              { k: "created", label: "作成した記事" },
              { k: "liked", label: "いいねした記事" },
              // TODO: コミュニティのアクセス履歴を今後実装する
              // { k: "community-history", label: "アクセスしたコミュニティ" },
            ].map(({ k, label }) => {
              const active = tab === (k as typeof tab);
              return (
                <button
                  key={k}
                  onClick={() => setTab(k as typeof tab)}
                  className="relative pb-3 text-gray-600 dark:text-gray-300 hover:text-gray-900 dark:hover:text-gray-100"
                >
                  {label}
                  <span
                    className={
                      "absolute left-0 right-0 -bottom-[2px] h-[2px] rounded-full " +
                      (active
                        ? "bg-gray-900 dark:bg-gray-100"
                        : "bg-transparent")
                    }
                  />
                </button>
              );
            })}
          </nav>
        </div>
      </section>
      {/* 下段：淡色キャンバス */}
      <section className="bg-indigo-50/60 dark:bg-gray-800/40 border-t border-gray-200 dark:border-gray-700 h-full">
        <div className="max-w-6xl mx-auto px-8 lg:px-12 h-full">
          <div className="pb-16">
            {tab === "created" && (
              <>
                {postsLoading ? (
                  <div className="py-16 text-center text-gray-600 dark:text-gray-300">
                    読み込み中...
                  </div>
                ) : postsError ? (
                  <div className="py-16 text-center text-red-600 dark:text-red-400">
                    記事の取得に失敗しました
                  </div>
                ) : posts.length === 0 ? (
                  // デザインどおりのプレースホルダー 4×2
                  <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-10 py-6">
                    {Array.from({ length: 8 }).map((_, i) => (
                      <div
                        key={i}
                        className="aspect-square rounded-2xl bg-white/70 dark:bg-gray-800 shadow-sm"
                      />
                    ))}
                  </div>
                ) : (
                  <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-10 py-6">
                    {posts.map((post) => (
                      <article
                        key={post.id}
                        className="bg-white dark:bg-gray-800 rounded-2xl shadow-sm hover:shadow-md transition p-4 aspect-square flex flex-col"
                      >
                        <Image
                          src="/default-image.jpg" // Placeholder header image
                          alt="Header Image"
                          width={200} // Adjust width as needed for this card size
                          height={100} // Adjust height as needed
                          className="w-full h-24 object-cover rounded-md mb-2"
                        />
                        <h3 className="text-base text-gray-900 dark:text-gray-100 line-clamp-2">
                          {post.title}
                        </h3>
                        <div className="mt-auto pt-4 flex justify-between items-center text-xs text-gray-500 dark:text-gray-400 gap-1">
                          <span className="flex items-center gap-1">
                            <FontAwesomeIcon
                              icon={faHeart}
                              className="text-gray-400"
                            />
                            {post.likesCount || 0}
                          </span>
                          <div className="flex gap-3">
                            <a
                              href={`/dashboard/post/${post.id}`}
                              className="hover:underline"
                            >
                              表示
                            </a>
                            <a
                              href={`/dashboard/post/${post.id}/edit`}
                              className="hover:underline"
                            >
                              編集
                            </a>
                          </div>
                        </div>
                      </article>
                    ))}
                  </div>
                )}
              </>
            )}

            {tab === "liked" && (
              <>
                {allPostsLoading ? (
                  <div className="py-16 text-center text-gray-600 dark:text-gray-300">
                    読み込み中...
                  </div>
                ) : allPostsError ? (
                  <div className="py-16 text-center text-red-600 dark:text-red-400">
                    記事の取得に失敗しました
                  </div>
                ) : likedPosts.length === 0 ? (
                  <div className="py-16 text-center text-gray-600 dark:text-gray-300">
                    いいねした記事はありません。
                  </div>
                ) : (
                  <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-10 py-6">
                    {likedPosts.map((post) => (
                      <article
                        key={post.id}
                        className="bg-white dark:bg-gray-800 rounded-2xl shadow-sm hover:shadow-md transition p-4 aspect-square flex flex-col"
                      >
                        <h3 className="text-base text-gray-900 dark:text-gray-100 line-clamp-2">
                          {post.title}
                        </h3>
                        <p className="mt-2 text-sm text-gray-600 dark:text-gray-300 line-clamp-3">
                          {post.description}
                        </p>
                        <div className="mt-auto pt-4 flex items-center justify-between text-xs text-gray-500 dark:text-gray-400">
                          <span>❤️ {post.likesCount || 0}</span>
                          <div className="flex gap-3">
                            <a
                              href={`/dashboard/post/${post.id}`}
                              className="hover:underline"
                            >
                              表示
                            </a>
                            <a
                              href={`/dashboard/post/${post.id}/edit`}
                              className="hover:underline"
                            >
                              編集
                            </a>
                          </div>
                        </div>
                      </article>
                    ))}
                  </div>
                )}
              </>
            )}
          </div>
        </div>
      </section>
    </>
  );
}
