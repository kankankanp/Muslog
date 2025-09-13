"use client";

import { Tag, Music } from "lucide-react";
import Image from "next/image";
import { useParams, useRouter } from "next/navigation";
import React, { useState, useEffect, useRef } from "react";
import ReactMarkdown from "react-markdown";
import ImageUploadModal from "@/components/elements/modals/ImageUploadModal"; // New
import SpotifySearchModal from "@/components/elements/modals/SpotifySearchModal";
import TagModal from "@/components/elements/modals/TagModal";
import { useGetMe } from "@/libs/api/generated/orval/auth/auth";
import { usePostImagesUpload } from "@/libs/api/generated/orval/images/images";
import { PostPostsBody } from "@/libs/api/generated/orval/model";
import { Track } from "@/libs/api/generated/orval/model/track";
import {
  useGetPostsId,
  usePutPostsId,
  useDeletePostsId,
  usePostPostsPostIdHeaderImage, // New
} from "@/libs/api/generated/orval/posts/posts";

export default function EditPostPage() {
  const params = useParams();
  const { id } = params as { id: string };

  const [title, setTitle] = useState("");
  const [markdown, setMarkdown] = useState("");
  const [viewMode, setViewMode] = useState<"editor" | "preview" | "split">(
    "split"
  );
  const [previewZoom, setPreviewZoom] = useState(1.0); // Default to 1.0 for font-size scaling
  const [editorZoom, setEditorZoom] = useState(1.0);
  const [editorWidth, setEditorWidth] = useState(50); // Initial width for editor in split view
  const [previewWidth, setPreviewWidth] = useState(50); // Initial width for preview in split view

  const [isTagModalOpen, setIsTagModalOpen] = useState(false);
  const [isSpotifyModalOpen, setIsSpotifyModalOpen] = useState(false);
  const [finalSelectedTracks, setFinalSelectedTracks] = useState<Track[]>([]); // New state for final selected tracks
  const [finalSelectedTags, setFinalSelectedTags] = useState<string[]>([]); // New state for final selected tags

  const [isHeaderImageModalOpen, setIsHeaderImageModalOpen] = useState(false); // New
  const [headerImageUrl, setHeaderImageUrl] = useState<string | undefined>(
    undefined
  ); // New
  const [currentUploadType, setCurrentUploadType] = useState<
    "header" | "in-post" | null
  >(null); // New

  const containerRef = useRef<HTMLDivElement>(null);

  const router = useRouter();
  const {
    data: userData,
    isLoading: isUserLoading,
    isError: isUserError,
    error: userError,
  } = useGetMe();
  const {
    data: postData,
    isLoading: isPostLoading,
    error: postError,
  } = useGetPostsId(Number(id));
  const { mutate: updatePost } = usePutPostsId();
  const { mutate: deletePost } = useDeletePostsId();
  const { mutate: uploadHeaderImage } = usePostPostsPostIdHeaderImage(); // New
  const { mutate: uploadGenericImage } = usePostImagesUpload(); // New

  useEffect(() => {
    if (postData?.post) {
      setTitle(postData.post.title);
      setMarkdown(postData.post.description);
      setFinalSelectedTracks(postData.post.tracks || []);
      setFinalSelectedTags(
        postData?.post?.tags
          ?.map((tag) => tag.name)
          .filter((name): name is string => typeof name === "string") || []
      );
      setHeaderImageUrl(postData?.post?.headerImageUrl ?? undefined); // Initialize header image URL
    }
  }, [postData]);

  useEffect(() => {
    // Modal.setAppElement is now handled by the individual modal components or a higher-level component.
  }, []);

  const handleZoom = (
    area: "editor" | "preview",
    type: "in" | "out" | "reset"
  ) => {
    const step = 0.1;
    const minZoom = 0.5;
    const maxZoom = 2.0;

    if (area === "preview") {
      setPreviewZoom((prev) => {
        if (type === "reset") return 1.0;
        const newZoom = type === "in" ? prev + step : prev - step;
        return Math.max(minZoom, Math.min(maxZoom, newZoom));
      });
    } else {
      // editor
      setEditorZoom((prev) => {
        if (type === "reset") return 1.0;
        const newZoom = type === "in" ? prev + step : prev - step;
        return Math.max(minZoom, Math.min(maxZoom, newZoom));
      });
    }
  };

  const handleHeaderImageUpload = (file: File) => {
    uploadHeaderImage(
      { postId: Number(id), data: { image: file } },
      {
        onSuccess: (response) => {
          setHeaderImageUrl(response.imageUrl);
          alert("ヘッダー画像を更新しました！");
        },
        onError: (error) => {
          console.error("ヘッダー画像の更新に失敗しました:", error);
          alert("ヘッダー画像の更新に失敗しました。");
        },
      }
    );
  };

  const handleInPostImageUpload = (file: File) => {
    const formData = new FormData();
    formData.append("image", file);

    uploadGenericImage(
      { data: { image: file } },
      {
        onSuccess: (response) => {
          // Insert the image URL into the markdown content
          setMarkdown(
            (prevMarkdown) =>
              `${prevMarkdown}\n![image](${response.imageUrl})\n`
          );
          alert("画像を投稿内に挿入しました！");
        },
        onError: (error) => {
          console.error("投稿内画像のアップロードに失敗しました:", error);
          alert("投稿内画像のアップロードに失敗しました。");
        },
      }
    );
  };

  const handleImageUpload = (file: File) => {
    if (currentUploadType === "header") {
      handleHeaderImageUpload(file);
    } else if (currentUploadType === "in-post") {
      handleInPostImageUpload(file);
    }
    setIsHeaderImageModalOpen(false); // Close modal after upload
  };

  const handleSubmit = () => {
    if (!userData?.id) {
      alert("ユーザー情報が取得できませんでした。ログインしてください。");
      return;
    }

    const userId = userData.id;

    const postDataToUpdate: PostPostsBody = {
      title,
      description: markdown,
      userId: userId,
      tracks: finalSelectedTracks,
      tags: finalSelectedTags,
      headerImageUrl: headerImageUrl, // New: Include header image URL
    };

    updatePost(
      { id: Number(id), data: postDataToUpdate },
      {
        onSuccess: () => {
          alert("記事を更新しました！");
          router.push(`/dashboard/post/${id}`);
        },
        onError: (err) => {
          console.error("Failed to update article:", err);
          alert("記事の更新に失敗しました。");
        },
      }
    );
  };

  const handleDelete = async () => {
    if (!confirm("本当にこの投稿を削除しますか？")) return;

    deletePost(
      { id: Number(id) },
      {
        onSuccess: () => {
          alert("投稿を削除しました。");
          router.push("/dashboard");
        },
        onError: (error: any) => {
          console.error("Delete error:", error);
          alert("削除に失敗しました。");
        },
      }
    );
  };

  const handleTrackSelect = (tracks: Track[]) => {
    setFinalSelectedTracks(tracks);
    setIsSpotifyModalOpen(false);
  };

  const handleRemoveFinalTrack = (trackToRemove: Track) => {
    setFinalSelectedTracks((prevTracks) =>
      prevTracks.filter((track) => track.spotifyId !== trackToRemove.spotifyId)
    );
  };

  const handleTagSelect = (tags: string[]) => {
    setFinalSelectedTags(tags);
    setIsTagModalOpen(false);
  };

  const handleRemoveFinalTag = (tagToRemove: string) => {
    setFinalSelectedTags((prevTags) =>
      prevTags.filter((tag) => tag !== tagToRemove)
    );
  };

  if (isPostLoading) {
    return (
      <div className="dark:bg-gray-900 bg-gray-100 min-h-screen flex items-center justify-center">
        <p className="text-gray-600 dark:text-gray-300">投稿を読み込み中...</p>
      </div>
    );
  }

  if (postError || !postData?.post) {
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

  return (
    <>
      <h1 className="text-3xl font-bold border-gray-100 border-b-2 bg-white px-6 py-6">
        記事を編集する
      </h1>
      <div ref={containerRef}>
        <div className="flex justify-center gap-4 mb-4 mt-4">
          <button
            className={`px-4 py-2 rounded ${viewMode === "editor" ? "bg-blue-500 text-white" : "bg-gray-200"}`}
            onClick={() => setViewMode("editor")}
          >
            エディタ
          </button>
          <button
            className={`px-4 py-2 rounded ${viewMode === "preview" ? "bg-blue-500 text-white" : "bg-gray-200"}`}
            onClick={() => setViewMode("preview")}
          >
            プレビュー
          </button>
          <button
            className={`px-4 py-2 rounded max-md:hidden ${viewMode === "split" ? "bg-blue-500 text-white" : "bg-gray-200"}`}
            onClick={() => setViewMode("split")}
          >
            分割
          </button>
        </div>
        <div className="flex h-screen bg-white">
          {/* 右側：プレビュー */}
          <div
            className={`p-8 overflow-y-auto ${viewMode === "editor" ? "hidden" : "flex-1"} ${viewMode === "split" ? "w-1/2" : ""}`}
          >
            <div className="flex gap-2 mb-2 justify-end">
              <button
                className="px-2 py-1 bg-gray-200 rounded"
                onClick={() => handleZoom("preview", "out")}
              >
                -
              </button>
              <button
                className="px-2 py-1 bg-gray-200 rounded"
                onClick={() => handleZoom("preview", "in")}
              >
                +
              </button>
              <button
                className="px-2 py-1 bg-gray-200 rounded"
                onClick={() => handleZoom("preview", "reset")}
              >
                リセット
              </button>
            </div>
            <div className="mb-6">
              {headerImageUrl && (
                <div className="relative w-full h-48 mb-4 rounded-md overflow-hidden">
                  <Image
                    src={headerImageUrl}
                    alt="Header Image"
                    layout="fill"
                    objectFit="cover"
                  />
                </div>
              )}
              <h2 className="text-3xl font-bold text-gray-400 mt-6">
                {title || "記事タイトル"}
              </h2>
            </div>
            <div
              className="prose prose-lg max-w-none w-full"
              style={{ fontSize: `${previewZoom * 16}px` }}
            >
              <ReactMarkdown>
                {markdown || "プレビューがここに表示されます。"}
              </ReactMarkdown>
            </div>
          </div>
          <div
            className={`p-8 flex flex-col gap-4 border-r border-gray-200 ${viewMode === "preview" ? "hidden" : "flex-1"} ${viewMode === "split" ? "md:w-1/2" : ""}`}
          >
            <div className="flex gap-2 mb-2 justify-end">
              <button
                className="px-2 py-1 bg-gray-200 rounded"
                onClick={() => handleZoom("editor", "out")}
              >
                -
              </button>
              <button
                className="px-2 py-1 bg-gray-200 rounded"
                onClick={() => handleZoom("editor", "in")}
              >
                +
              </button>
              <button
                className="px-2 py-1 bg-gray-200 rounded"
                onClick={() => handleZoom("editor", "reset")}
              >
                リセット
              </button>
            </div>
            <input
              type="text"
              placeholder="記事タイトル"
              className="text-3xl font-bold mb-4 bg-transparent outline-none"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
            />
            <div className="flex gap-2 max-md:flex-col max-md:gap-0">
              <button
                className="flex items-center gap-2 px-4 py-2 bg-gray-100 rounded w-fit mb-4"
                onClick={() => {
                  setCurrentUploadType("header");
                  setIsHeaderImageModalOpen(true);
                }}
              >
                <span className="text-xl">＋</span> ヘッダー画像
              </button>
              <button
                className="flex items-center gap-2 px-4 py-2 bg-gray-100 rounded w-fit mb-4"
                onClick={() => {
                  setCurrentUploadType("in-post");
                  setIsHeaderImageModalOpen(true);
                }}
              >
                <span className="text-xl">＋</span> 投稿内画像
              </button>
              <button
                className="flex items-center gap-2 px-4 py-2 bg-gray-100 rounded w-fit mb-4"
                onClick={() => setIsTagModalOpen(true)}
              >
                <Tag className="h-5 w-5" /> タグ
              </button>
              <button
                className="flex items-center gap-2 px-4 py-2 bg-gray-100 rounded w-fit mb-4"
                onClick={() => setIsSpotifyModalOpen(true)}
              >
                <Music className="h-5 w-5" /> 曲
              </button>
            </div>

            {finalSelectedTracks.length > 0 && (
              <div className="flex gap-2 mb-4 overflow-x-auto pb-2">
                {" "}
                {finalSelectedTracks.map((track) => (
                  <div
                    key={track.spotifyId}
                    className="flex items-center bg-gray-100 rounded-full px-3 py-1 text-sm flex-shrink-0" // Added flex-shrink-0
                  >
                    <Image
                      src={track.albumImageUrl || "/default-image.jpg"}
                      width={20}
                      height={20}
                      alt={track.name || ""}
                      className="rounded-full mr-2"
                    />
                    {track.name} - {track.artistName}
                    <button
                      onClick={() => handleRemoveFinalTrack(track)}
                      className="ml-2 text-gray-500 hover:text-gray-700"
                    >
                      &times;
                    </button>
                  </div>
                ))}
              </div>
            )}

            {finalSelectedTags.length > 0 && (
              <div className="flex gap-2 mb-4 overflow-x-auto pb-2">
                {" "}
                {finalSelectedTags.map((tag) => (
                  <div
                    key={tag}
                    className="flex items-center bg-gray-100 rounded-full px-3 py-1 text-sm flex-shrink-0" // Added flex-shrink-0
                  >
                    {tag}
                    <button
                      onClick={() => handleRemoveFinalTag(tag)}
                      className="ml-2 text-gray-500 hover:text-gray-700"
                    >
                      &times;
                    </button>
                  </div>
                ))}
              </div>
            )}

            <textarea
              className="flex-1 w-full border rounded p-4 resize-none bg-gray-50"
              placeholder="本文をマークダウンで入力してください"
              value={markdown}
              onChange={(e) => setMarkdown(e.target.value)}
              style={{ fontSize: `${editorZoom * 16}px` }}
            />
            <div className="flex justify-end mt-4">
              <button
                className="px-6 py-3 bg-blue-600 text-white rounded-lg text-lg font-semibold hover:bg-blue-700 transition-colors"
                onClick={handleSubmit}
              >
                記事を更新する
              </button>
              <button
                className="px-6 py-3 bg-red-600 text-white rounded-lg text-lg font-semibold hover:bg-red-700 transition-colors ml-4"
                onClick={handleDelete}
              >
                記事を削除する
              </button>
            </div>
          </div>

          <TagModal
            isOpen={isTagModalOpen}
            onClose={() => setIsTagModalOpen(false)}
            onSelectTags={handleTagSelect}
            initialSelectedTags={finalSelectedTags}
          />

          <SpotifySearchModal
            isOpen={isSpotifyModalOpen}
            onClose={() => setIsSpotifyModalOpen(false)}
            onSelectTracks={handleTrackSelect}
            initialSelectedTracks={finalSelectedTracks}
          />

          <ImageUploadModal
            isOpen={isHeaderImageModalOpen}
            onClose={() => setIsHeaderImageModalOpen(false)}
            onImageUpload={handleImageUpload}
            currentImageUrl={
              currentUploadType === "header" ? headerImageUrl : undefined
            }
          />
        </div>
      </div>
    </>
  );
}
