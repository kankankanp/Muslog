"use client";

import { useState } from "react";
import { Toaster } from "react-hot-toast";
import NewBlogForm from "@/app/components/elements/forms/NewBlogForm";
import SelectMusicArea from "@/app/components/elements/others/SelectMusciArea";
import { Track } from "@/app/libs/api/generated/orval/model/track";

export default function Page() {
  const [selectedTrack, setSelectedTrack] = useState<Track | null>(null);

  return (
    <main className="dark:bg-gray-900 bg-gray-100 pt-8">
      <Toaster />
      <div className="grid grid-cols-1 md:grid-cols-2 gap-8 py-8">
        <SelectMusicArea onSelect={setSelectedTrack} />
        <NewBlogForm selectedTrack={selectedTrack} />
      </div>
    </main>
  );
}
