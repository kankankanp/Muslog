"use client";

import { useState } from "react";
import toast from "react-hot-toast";

export type Track = {
  id: string;
  name: string;
  artists: { name: string }[];
  album: { images: { url: string }[] };
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
    <div className="w-full max-w-md mx-auto bg-white p-6 rounded-lg shadow-md">
      <div className="flex gap-2 mb-4">
        <input
          type="text"
          placeholder="曲名を入力"
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          className="flex-1 px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-green-500"
        />
        <button
          onClick={handleSearch}
          disabled={loading}
          className="px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 disabled:opacity-50"
        >
          {loading ? "検索中…" : "検索"}
        </button>
      </div>

      <ul className="space-y-2">
        {tracks.map((track) => (
          <li
            key={track.id}
            onClick={() => onSelect(track)}
            className="flex items-center gap-4 p-3 border rounded-md cursor-pointer hover:bg-gray-100"
          >
            <img
              src={track.album.images[0]?.url}
              width={50}
              height={50}
              alt={track.name}
              className="rounded"
            />
            <div>
              <p className="font-medium">{track.name}</p>
              <p className="text-sm text-gray-500">{track.artists[0].name}</p>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default SelectMusicArea;
