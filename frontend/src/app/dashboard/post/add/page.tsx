"use client";

import { Tag, Music } from "lucide-react";
import Image from "next/image";
import React, { useState, useEffect, useRef } from "react";
import ReactMarkdown from "react-markdown";
import SpotifySearchModal from "@/components/elements/modals/SpotifySearchModal";
import TagModal from "@/components/elements/modals/TagModal";
import { Track } from "@/libs/api/generated/orval/model/track";

export default function AddPostPage() {
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

  const containerRef = useRef<HTMLDivElement>(null);

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

  const handleSubmit = () => {
    console.log("Post submitted:", {
      title,
      markdown,
      selectedTrack: finalSelectedTracks,
    }); // Updated to use finalSelectedTracks
    // TODO: Implement actual post submission logic
    alert("記事を投稿しました！ (実際にはまだ投稿されていません)");
  };

  const handleTrackSelect = (tracks: Track[]) => {
    // Modified to accept array of tracks
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

  return (
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
          className={`px-4 py-2 rounded ${viewMode === "split" ? "bg-blue-500 text-white" : "bg-gray-200"}`}
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
        {/* 左側：エディタ */}
        <div
          className={`p-8 flex flex-col gap-4 border-r border-gray-200 ${viewMode === "preview" ? "hidden" : "flex-1"} ${viewMode === "split" ? "w-1/2" : ""}`}
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
          <div className="flex gap-2">
            <button
              className="flex items-center gap-2 px-4 py-2 bg-gray-100 rounded w-fit mb-4"
              // 画像追加のロジックは後で
            >
              <span className="text-xl">＋</span> 画像を追加
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
              {/* Added overflow-x-auto and pb-2 for scrollbar */}
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
            <div className="flex gap-2 mb-4 overflow-x-auto pb-2"> {/* Added overflow-x-auto and pb-2 for scrollbar */}
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
          <button
            className="mt-4 px-6 py-3 bg-blue-600 text-white rounded-lg text-lg font-semibold hover:bg-blue-700 transition-colors"
            onClick={handleSubmit}
          >
            記事を投稿する
          </button>
        </div>
      </div>

      <TagModal
        isOpen={isTagModalOpen}
        onClose={() => setIsTagModalOpen(false)}
        onSelectTags={handleTagSelect} // Assuming handleTagSelect will be created
        initialSelectedTags={finalSelectedTags} // Assuming finalSelectedTags will be created
      />

      <SpotifySearchModal
        isOpen={isSpotifyModalOpen}
        onClose={() => setIsSpotifyModalOpen(false)}
        onSelectTracks={handleTrackSelect}
        initialSelectedTracks={finalSelectedTracks}
      />
    </div>
  );
}
