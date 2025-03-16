"use client";

import { useState } from "react";
import { Toaster } from "react-hot-toast";
import SelectMusicArea, {
  Track,
} from "@/app/components/elements/SelectMusciArea";
import NewBlogForm from "@/app/components/elements/NewBlogForm";

const PostBlog = () => {
  const [selectedTrack, setSelectedTrack] = useState<Track | null>(null);

  return (
    <>
      <Toaster />
      <div className="grid grid-cols-1 md:grid-cols-2 gap-8 py-8 bg-gray-100">
        <SelectMusicArea onSelect={setSelectedTrack} />
        <NewBlogForm selectedTrack={selectedTrack} />
      </div>
    </>
  );
};

export default PostBlog;
