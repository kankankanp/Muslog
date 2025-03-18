"use client";

import Image from "next/image";
import { useState } from "react";
import toast from "react-hot-toast";

export type Track = {
  id: number;
  spotifyId: string;
  name: string;
  artistName: string;
  albumImageUrl: string;
};

type SelectMusicAreaProps = {
  onSelect: (track: Track) => void;
};

const SelectMusicArea = ({ onSelect }: SelectMusicAreaProps): JSX.Element => {
  const [query, setQuery] = useState("");
  const [tracks, setTracks] = useState<Track[]>([]);
  const [loading, setLoading] = useState(false);

  const handleSearch = async () => {
    if (!query) return toast.error("検索ワードを入力してください");

    setLoading(true);
    try {
      const res = await fetch(
        `/api/spotify/search?q=${encodeURIComponent(query)}`
      );
      const data = await res.json();

      if (res.ok) {
        setTracks(data);
      } else {
        toast.error(data.error || "検索に失敗しました");
      }
    } catch (err) {
      toast.error("通信エラーが発生しました");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="w-full max-w-md mx-auto bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
      <div className="flex gap-2 mb-4">
        <input
          type="text"
          placeholder="曲名を入力"
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          className="flex-1 px-4 py-2 border dark:border-gray-600 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 dark:bg-gray-700 dark:text-gray-200"
        />
        <button
          onClick={handleSearch}
          disabled={loading}
          className="px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 dark:hover:bg-green-700 disabled:opacity-50"
        >
          {loading ? "検索中…" : "検索"}
        </button>
      </div>

      <ul className="space-y-2">
        {tracks.map((track) => (
          <li
            key={track.id}
            onClick={() => onSelect(track)}
            className="flex items-center gap-4 p-3 border dark:border-gray-600 rounded-md cursor-pointer hover:bg-gray-100 dark:hover:bg-gray-700"
          >
            <Image
              src={track.albumImageUrl}
              width={50}
              height={50}
              alt={track.name}
              className="rounded"
            />
            <div>
              <p className="font-medium dark:text-white">{track.name}</p>
              <p className="text-sm text-gray-500 dark:text-gray-400">
                {track.artistName}
              </p>
            </div>
            {/* TODO 曲を再生できるようにする */}
            {/* <div>
              <Image src="/play_circle.png" alt="再生する" />
              <Image src="/stop_circle.png" alt="停止する" />
            </div> */}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default SelectMusicArea;
