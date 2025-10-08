"use client";

import { faHeart } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { useQueryClient } from "@tanstack/react-query";
import Image from "next/image";
import { useMemo, useState, useRef } from "react";
import BandRecruitmentCard from "@/components/bandRecruitment/BandRecruitmentCard";
import { useGetMe } from "@/libs/api/generated/orval/auth/auth";
import {
  useGetBandRecruitments,
  useGetBandRecruitmentsAppliedMe,
} from "@/libs/api/generated/orval/band-recruitments/band-recruitments";
import { useGetUsersMeLikedPosts } from "@/libs/api/generated/orval/likes/likes";
import {
  GetPosts200,
  GetUsersMeLikedPosts200,
} from "@/libs/api/generated/orval/model";
import {
  useGetUsersIdPosts,
  usePostUsersUserIdProfileImage,
} from "@/libs/api/generated/orval/users/users";
import { useGetCommunities } from "@/libs/api/generated/orval/communities/communities";

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
  const {
    data: likedPostsData,
    isPending: likedPostsLoading,
    error: likedPostsError,
  } = useGetUsersMeLikedPosts<GetUsersMeLikedPosts200>();
  const {
    data: appliedRecruitmentsData,
    isPending: appliedLoading,
    error: appliedError,
  } = useGetBandRecruitmentsAppliedMe();
  const {
    data: communitiesData,
    isPending: communitiesLoading,
    error: communitiesError,
  } = useGetCommunities({
    query: {
      enabled: Boolean(currentUser?.id),
    },
  });
  const {
    data: bandRecruitmentsData,
    isPending: createdBandsLoading,
    error: createdBandsError,
  } = useGetBandRecruitments(
    { page: 1, perPage: 50 },
    {
      query: {
        enabled: Boolean(currentUser?.id),
      },
    },
  );
  const [tab, setTab] = useState<
    | "created"
    | "liked"
    | "created-communities"
    | "created-bands"
    | "applied"
    | "community-history"
  >("created");

  const fileInputRef = useRef<HTMLInputElement>(null);
  const queryClient = useQueryClient();
  const { mutate: uploadProfileImage } = usePostUsersUserIdProfileImage();

  const handleFileChange = async (
    event: React.ChangeEvent<HTMLInputElement>,
  ) => {
    const file = event.target.files?.[0];
    if (!file) return;

    if (!currentUser?.id) {
      alert("ユーザーIDが見つかりません。");
      return;
    }

    const formData = new FormData();
    formData.append("image", file);

    uploadProfileImage(
      { userId: currentUser.id, data: { image: file } },
      {
        onSuccess: () => {
          alert("プロフィール画像を更新しました！");
          queryClient.invalidateQueries({ queryKey: ["getUserMe"] });
        },
        onError: (error) => {
          console.error("プロフィール画像の更新に失敗しました:", error);
          alert("プロフィール画像の更新に失敗しました。");
        },
      },
    );
  };

  const posts = postsData?.posts ?? [];
  const likedPosts = likedPostsData?.posts ?? [];
  const createdCommunities = useMemo(() => {
    if (!currentUser?.id) return [];
    return (
      communitiesData?.communities?.filter(
        (community) => community.creatorId === currentUser.id,
      ) ?? []
    );
  }, [communitiesData?.communities, currentUser?.id]);
  const createdRecruitments = useMemo(() => {
    if (!currentUser?.id) return [];
    return (
      bandRecruitmentsData?.recruitments?.filter(
        (recruitment) => recruitment.userId === currentUser.id,
      ) ?? []
    );
  }, [bandRecruitmentsData?.recruitments, currentUser?.id]);

  const formatDate = (value?: string) => {
    if (!value) return "";
    const date = new Date(value);
    if (Number.isNaN(date.getTime())) return value;
    return date.toLocaleDateString();
  };

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

  return (
    <>
      <h1 className="text-3xl font-bold border-gray-100 border-b-2 bg-white px-8 py-6">
        マイページ
      </h1>
      <section className="bg-white dark:bg-gray-800">
        <div className="max-w-6xl mx-auto px-8 lg:px-12 pt-10">
          <div className="flex gap-8 max-md:flex-col max-md:items-center">
            {/* 固定サイズの円アバター（ピル化を防ぐ） */}
            <div
              className="shrink-0 w-44 h-44 rounded-full bg-gray-200 dark:bg-gray-700 flex items-center justify-center relative overflow-hidden cursor-pointer"
              onClick={() => fileInputRef.current?.click()}
            >
              {currentUser.profileImageUrl ? (
                <Image
                  src={currentUser.profileImageUrl}
                  alt="Profile Picture"
                  layout="fill"
                  objectFit="cover"
                />
              ) : (
                <span className="text-2xl text-gray-600 dark:text-gray-200">
                  No Image
                </span>
              )}
              <input
                type="file"
                ref={fileInputRef}
                onChange={handleFileChange}
                accept="image/*"
                className="hidden"
              />
            </div>

            <div className="flex justify-center gap-1 flex-col text-center">
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
              { k: "created-communities", label: "作成したコミュニティ" },
              { k: "created-bands", label: "作成したバンド" },
              { k: "applied", label: "応募済みのバンド" },
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
                    {posts.map((post) => {
                      const headerImageSrc =
                        post.headerImageUrl && post.headerImageUrl.trim() !== ""
                          ? post.headerImageUrl
                          : "/default-image.jpg";
                      return (
                        <article
                          key={post.id}
                          className="bg-white dark:bg-gray-800 rounded-2xl shadow-sm hover:shadow-md transition p-4 aspect-square flex flex-col"
                        >
                          <Image
                            src={headerImageSrc}
                            alt="Header Image"
                            width={200}
                            height={100}
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
                      );
                    })}
                  </div>
                )}
              </>
            )}

            {tab === "liked" && (
              <>
                {likedPostsLoading ? (
                  <div className="py-16 text-center text-gray-600 dark:text-gray-300">
                    読み込み中...
                  </div>
                ) : likedPostsError ? (
                  <div className="py-16 text-center text-red-600 dark:text-red-400">
                    いいねした記事の取得に失敗しました
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
                          </div>
                        </div>
                      </article>
                    ))}
                  </div>
                )}
              </>
            )}

            {tab === "created-communities" && (
              <>
                {communitiesLoading ? (
                  <div className="py-16 text-center text-gray-600 dark:text-gray-300">
                    読み込み中...
                  </div>
                ) : communitiesError ? (
                  <div className="py-16 text-center text-red-600 dark:text-red-400">
                    作成したコミュニティの取得に失敗しました
                  </div>
                ) : createdCommunities.length === 0 ? (
                  <div className="py-16 text-center text-gray-600 dark:text-gray-300">
                    作成したコミュニティはありません。
                  </div>
                ) : (
                  <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8 py-6">
                    {createdCommunities.map((community) => (
                      <article
                        key={community.id}
                        className="flex flex-col gap-4 rounded-2xl bg-white p-6 shadow-sm transition hover:shadow-md dark:bg-gray-800"
                      >
                        <h3 className="text-lg font-semibold text-gray-900 dark:text-gray-100 line-clamp-2">
                          {community.name}
                        </h3>
                        <p className="text-sm text-gray-600 dark:text-gray-300 line-clamp-3">
                          {community.description}
                        </p>
                        <div className="mt-auto flex items-center justify-between text-sm text-gray-500 dark:text-gray-400">
                          <span>作成日: {formatDate(community.createdAt)}</span>
                          <a
                            href={`/dashboard/community/${community.id}`}
                            className="text-indigo-600 hover:underline"
                          >
                            詳細
                          </a>
                        </div>
                      </article>
                    ))}
                  </div>
                )}
              </>
            )}

            {tab === "created-bands" && (
              <>
                {createdBandsLoading ? (
                  <div className="py-16 text-center text-gray-600 dark:text-gray-300">
                    読み込み中...
                  </div>
                ) : createdBandsError ? (
                  <div className="py-16 text-center text-red-600 dark:text-red-400">
                    作成したバンドの取得に失敗しました
                  </div>
                ) : createdRecruitments.length === 0 ? (
                  <div className="py-16 text-center text-gray-600 dark:text-gray-300">
                    作成したバンドはありません。
                  </div>
                ) : (
                  <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8 py-6">
                    {createdRecruitments.map((recruitment) => (
                      <BandRecruitmentCard
                        key={recruitment.id}
                        recruitment={recruitment}
                      />
                    ))}
                  </div>
                )}
              </>
            )}

            {tab === "applied" && (
              <>
                {appliedLoading ? (
                  <div className="py-16 text-center text-gray-600 dark:text-gray-300">
                    読み込み中...
                  </div>
                ) : appliedError ? (
                  <div className="py-16 text-center text-red-600 dark:text-red-400">
                    応募済みバンドの取得に失敗しました
                  </div>
                ) : (appliedRecruitmentsData?.recruitments?.length || 0) === 0 ? (
                  <div className="py-16 text-center text-gray-600 dark:text-gray-300">
                    応募済みのバンドはありません。
                  </div>
                ) : (
                  <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8 py-6">
                    {appliedRecruitmentsData?.recruitments?.map((recruitment) => (
                      <BandRecruitmentCard
                        key={recruitment.id}
                        recruitment={recruitment}
                      />
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
