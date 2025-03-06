"use client";

import { useState } from "react";
import { Toaster } from "react-hot-toast";
import SelectMusicArea, {
  Track,
} from "@/app/components/elements/SelectMusciArea";
import NewBlogForm from "@/app/components/elements/NewBlogForm";
import Footer from "@/app/components/layouts/Footer";
import Header from "@/app/components/layouts/Header";

const PostBlog = () => {
  const [selectedTrack, setSelectedTrack] = useState<Track | null>(null);

  return (
    <>
      <Toaster />
      {/* <Header /> */}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
        <SelectMusicArea onSelect={setSelectedTrack} />
        <NewBlogForm selectedTrack={selectedTrack} />
      </div>
      {/* <Footer /> */}
    </>
  );
};

export default PostBlog;
