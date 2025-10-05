"use client";

import Image from "next/image";
import React, { useState } from "react";
import toast from "react-hot-toast";
import { Track } from "@/libs/api/generated/orval/model/track";
import { useGetSpotifySearch } from "@/libs/api/generated/orval/spotify/spotify";

type SelectMusicAreaProps = {
  onSelect: (tracks: Track[]) => void;
  initialSelectedTracks?: Track[]; // New prop
};

const SelectMusicArea = ({
  onSelect,
  initialSelectedTracks,
}: SelectMusicAreaProps): JSX.Element => {
  // Destructure new prop
  const [query, setQuery] = useState<string>("");
  const [searchQuery, setSearchQuery] = useState<string>("");
  const [buttonLoading, setButtonLoading] = useState<boolean>(false); // New state for button loading
  const [selectedTracksInModal, setSelectedTracksInModal] = useState<Track[]>(
    initialSelectedTracks || [],
  ); // Initialize with prop

  const { data, isPending, error } = useGetSpotifySearch(
    { q: searchQuery },
    { query: { enabled: !!searchQuery } },
  );

  // Use useEffect to watch for changes in isPending from the hook
  // and update buttonLoading accordingly
  React.useEffect(() => {
    if (!isPending && buttonLoading) {
      // If API call is no longer pending and button was loading
      setButtonLoading(false); // Reset button loading
    }
  }, [isPending, buttonLoading]);

  const handleSearch = async () => {
    if (!query.trim()) {
      toast.error("検索ワードを入力してください");
      return;
    }
    setButtonLoading(true); // Set button loading to true
    setSearchQuery(query.trim());
  };

  const handleTrackToggle = (trackToToggle: Track) => {
    setSelectedTracksInModal((prevTracks) => {
      if (prevTracks.some((t) => t.spotifyId === trackToToggle.spotifyId)) {
        // Remove track if already selected
        return prevTracks.filter(
          (t) => t.spotifyId !== trackToToggle.spotifyId,
        );
      } else {
        // Add track if not selected
        return [...prevTracks, trackToToggle];
      }
    });
  };

  const tracks = data?.tracks || [];

  // エラーの監視と表示
  React.useEffect(() => {
    if (error) {
      console.error("Spotify search error:", error);
      toast.error(
        (error as any)?.response?.data?.message ||
          (error as any)?.message ||
          "検索中にエラーが発生しました",
      );
    }
  }, [error]);

  return (
    <div className="w-full max-w-md mx-auto bg-white dark:bg-gray-800 p-6 rounded-lg">
      <div className="flex gap-2 mb-4">
        <input
          type="text"
          placeholder="曲名を入力"
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          onKeyPress={(e) => e.key === "Enter" && handleSearch()}
          className="flex-1 px-4 py-2 border dark:border-gray-600 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 dark:bg-gray-700 dark:text-gray-200"
        />
        <button
          onClick={handleSearch}
          disabled={buttonLoading}
          className="px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 dark:hover:bg-green-700 disabled:opacity-50"
        >
          {buttonLoading ? "検索中…" : "検索"}
        </button>
      </div>

      <ul className="space-y-2">
        {tracks.map((track) => (
          <li
            key={track.spotifyId}
            onClick={() => handleTrackToggle(track)} // Modified onClick
            className={`flex items-center gap-4 p-3 border dark:border-gray-600 rounded-md cursor-pointer hover:bg-gray-100 dark:hover:bg-gray-700 ${
              selectedTracksInModal.some((t) => t.spotifyId === track.spotifyId)
                ? "bg-blue-100 dark:bg-blue-900 border-blue-500" // Highlight selected
                : ""
            }`}
          >
            <Image
              src={track.albumImageUrl || "/default-image.jpg"}
              width={50}
              height={50}
              alt={track.name || ""}
              className="rounded"
            />
            <div>
              <p className="font-medium dark:text-white">{track.name}</p>
              <p className="text-sm text-gray-500 dark:text-gray-400">
                {track.artistName}
              </p>
            </div>
          </li>
        ))}
      </ul>

      {selectedTracksInModal.length > 0 && (
        <div className="mt-4 pt-4 border-t dark:border-gray-600">
          <h3 className="font-semibold mb-2 dark:text-white">選択中の曲:</h3>
          <div className="flex flex-wrap gap-2 max-h-40 overflow-y-auto">
            {selectedTracksInModal.map((track) => (
              <div
                key={track.spotifyId}
                className="flex items-center bg-gray-100 dark:bg-gray-700 rounded-full px-3 py-1 text-sm dark:text-gray-200"
              >
                {track.name} - {track.artistName}
                <button
                  onClick={(e) => {
                    e.stopPropagation(); // Prevent li onClick from firing
                    handleTrackToggle(track);
                  }}
                  className="ml-2 text-gray-500 dark:text-gray-400 hover:text-gray-700 dark:hover:text-gray-200"
                >
                  &times;
                </button>
              </div>
            ))}
          </div>
          <div className="flex justify-end mt-4">
            <button
              onClick={() => onSelect(selectedTracksInModal)} // Call onSelect with array
              className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 dark:hover:bg-blue-700"
            >
              完了
            </button>
          </div>
        </div>
      )}
    </div>
  );
};

export default SelectMusicArea;
