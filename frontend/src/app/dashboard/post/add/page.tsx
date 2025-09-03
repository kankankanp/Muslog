"use client";

import { useState } from "react";
import { Toaster } from "react-hot-toast";
import NewBlogForm from "@/components/elements/forms/NewBlogForm";
import SelectMusicArea from "@/components/elements/others/SelectMusciArea";
import { Track } from "@/libs/api/generated/orval/model/track";

export default function Page() {
  const [selectedTrack, setSelectedTrack] = useState<Track | null>(null);

  return (
    <>
      <Toaster />
      <h1 className="text-3xl font-bold border-gray-100 border-b-2 bg-white px-6 py-6">
        記事を作成する
      </h1>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-8 py-8">
        <SelectMusicArea onSelect={setSelectedTrack} />
        <NewBlogForm selectedTrack={selectedTrack} />
      </div>
    </>
  );
}
