"use client";

import { Music } from "lucide-react";
import { Carousel } from "react-responsive-carousel";
import "react-responsive-carousel/lib/styles/carousel.min.css";

interface Track {
  id: string;
  title: string;
  artist: string;
  albumArt: string;
}

interface RecentTracksProps {
  tracks: Track[];
}

const RecentTracks: React.FC<RecentTracksProps> = ({ tracks }) => {
  console.log(
    "RecentTracks component received tracks:",
    JSON.stringify(tracks, null, 2)
  );

  if (!tracks || tracks.length === 0) {
    return null;
  }

  return (
    <div className="w-full max-w-4xl mx-auto mt-8 p-6 bg-white dark:bg-gray-800 rounded-lg shadow-lg">
      <h2 className="text-2xl font-bold mb-6 text-gray-800 dark:text-white flex items-center gap-2">
        <Music className="w-6 h-6" />
        最近聴いた曲
      </h2>
      <Carousel
        showArrows={true}
        showStatus={false}
        showThumbs={false}
        infiniteLoop={true}
        autoPlay={true}
        interval={5000}
        className="recent-tracks-carousel"
      >
        {tracks.map((track) => (
          <div key={track.id} className="p-4">
            <div className="flex items-center space-x-4">
              <img
                src={track.albumArt}
                alt={track.title}
                className="w-24 h-24 rounded-lg shadow-md"
              />
              <div className="text-left">
                <h3 className="text-lg font-semibold text-gray-800 dark:text-white">
                  {track.title}
                </h3>
                <p className="text-gray-600 dark:text-gray-300">
                  {track.artist}
                </p>
              </div>
            </div>
          </div>
        ))}
      </Carousel>
    </div>
  );
};

export default RecentTracks;
