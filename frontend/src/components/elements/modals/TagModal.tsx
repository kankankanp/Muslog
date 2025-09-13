"use client";

import React, { useEffect, useMemo, useState } from "react";
import Modal from "react-modal";
import type { Tag } from "@/libs/api/generated/orval/model/tag";
import { useGetTags, usePostTags } from "@/libs/api/generated/orval/tags/tags";

type TagModalProps = {
  isOpen: boolean;
  onClose: () => void;
  onSelectTags: (tags: string[]) => void;
  initialSelectedTags: string[];
};

const TagModal = ({
  isOpen,
  onClose,
  onSelectTags,
  initialSelectedTags,
}: TagModalProps): JSX.Element => {
  const [query, setQuery] = useState<string>("");
  const [selectedTagsInModal, setSelectedTagsInModal] = useState<string[]>(
    initialSelectedTags || []
  );

  // Fetch all tags
  const { data: tagsData, isPending, error, refetch } = useGetTags();
  const allTags: string[] = useMemo(() => {
    const list = (tagsData?.tags as Tag[] | undefined) || [];
    return list
      .map((t) => t.name)
      .filter((n): n is string => typeof n === "string");
  }, [tagsData]);

  // Create tag mutation
  const { mutate: createTag, isPending: isCreating } = usePostTags();

  // Keep selected state in sync when modal opens with different initial values
  useEffect(() => {
    setSelectedTagsInModal(initialSelectedTags || []);
  }, [initialSelectedTags, isOpen]);

  const handleTagToggle = (tagToToggle: string) => {
    setSelectedTagsInModal((prevTags) => {
      if (prevTags.includes(tagToToggle)) {
        return prevTags.filter((tag) => tag !== tagToToggle);
      } else {
        return [...prevTags, tagToToggle];
      }
    });
  };

  const handleAddCustomTag = () => {
    const name = query.trim();
    if (!name) return;

    // 既存にあるなら選択だけ追加
    if (allTags.includes(name)) {
      if (!selectedTagsInModal.includes(name)) {
        setSelectedTagsInModal((prev) => [...prev, name]);
      }
      setQuery("");
      return;
    }

    // 新規作成
    createTag(
      { data: { name } },
      {
        onSuccess: async () => {
          // 再取得してリストに反映
          await refetch();
          setSelectedTagsInModal((prev) =>
            prev.includes(name) ? prev : [...prev, name]
          );
          console.log("bbbb");
          setQuery("");
        },
        onError: () => {
          alert("タグの作成に失敗しました");
        },
      }
    );
  };

  return (
    <Modal
      isOpen={isOpen}
      onRequestClose={onClose}
      contentLabel="タグ選択/作成"
      className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 bg-white p-6 rounded-lg shadow-lg max-w-md w-full outline-none overflow-auto"
      overlayClassName="fixed inset-0 bg-black bg-opacity-75 z-50"
    >
      <div className="p-6 bg-white rounded-lg shadow-lg max-w-md mx-auto my-20">
        <h2 className="text-2xl font-bold mb-4">タグを選択または作成</h2>
        <form
          className="flex gap-2 mb-4"
          onSubmit={(e) => {
            e.preventDefault();
            e.stopPropagation();
            handleAddCustomTag();
          }}
        >
          <input
            type="text"
            placeholder="新しいタグを作成"
            className="w-full border rounded p-2"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                e.preventDefault();
                e.stopPropagation();
                handleAddCustomTag();
              }
            }}
          />
          <button
            type="submit"
            className="px-4 py-2 bg-blue-500 text-white rounded"
            disabled={isCreating}
          >
            追加
          </button>
        </form>

        <div className="mb-4">
          <h3 className="font-semibold mb-2">既存のタグ</h3>
          {isPending ? (
            <div className="text-sm text-gray-500">読み込み中...</div>
          ) : error ? (
            <div className="text-sm text-red-600">タグの取得に失敗しました</div>
          ) : (
            <div className="flex flex-wrap gap-2">
              {allTags
                .filter((t) =>
                  query.trim()
                    ? t.toLowerCase().includes(query.trim().toLowerCase())
                    : true
                )
                .map((tag) => (
                  <span
                    key={tag}
                    onClick={() => handleTagToggle(tag)}
                    className={`px-3 py-1 rounded-full text-sm cursor-pointer ${
                      selectedTagsInModal.includes(tag)
                        ? "bg-blue-500 text-white"
                        : "bg-gray-200"
                    }`}
                  >
                    {tag}
                  </span>
                ))}
            </div>
          )}
        </div>

        {selectedTagsInModal.length > 0 && (
          <div className="mt-4 pt-4 border-t dark:border-gray-600">
            <h3 className="font-semibold mb-2">選択中のタグ:</h3>
            <div className="flex flex-wrap gap-2">
              {selectedTagsInModal.map((tag) => (
                <div
                  key={tag}
                  className="flex items-center bg-gray-100 rounded-full px-3 py-1 text-sm"
                >
                  {tag}
                  <button
                    onClick={() => handleTagToggle(tag)}
                    className="ml-2 text-gray-500 hover:text-gray-700"
                  >
                    &times;
                  </button>
                </div>
              ))}
            </div>
            <div className="flex justify-end mt-4">
              <button
                onClick={() => onSelectTags(selectedTagsInModal)}
                className="px-4 py-2 bg-blue-500 text-white rounded disabled:opacity-60"
                disabled={isCreating}
              >
                {isCreating ? "保存中..." : "完了"}
              </button>
            </div>
          </div>
        )}

        <div className="flex justify-end gap-2 mt-4">
          <button className="px-4 py-2 bg-gray-300 rounded" onClick={onClose}>
            閉じる
          </button>
        </div>
      </div>
    </Modal>
  );
};

export default TagModal;
