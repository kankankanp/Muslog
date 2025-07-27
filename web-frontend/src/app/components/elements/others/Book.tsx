"use client";

import Image from "next/image";
import { useCallback, useMemo } from "react";
import "@/scss/book.scss";
import { Post } from "../cards/BlogCard";

type BookProps = {
  posts: Post[];
};

export const Book = ({ posts }: BookProps) => {
  const pageCount = 20;
  const pages = useMemo(
    () => Array.from({ length: pageCount }, (_, i) => 99 - i),
    [pageCount]
  );

  const playAudio = useCallback(() => {
    const audio = new Audio("/flip_sound.mp3");
    audio.play();
  }, []);

  const postPages = useMemo(
    () => pages.map((_, idx) => posts[idx - 1] || null),
    [posts, pages]
  );

  return (
    <div className="book">
      {pages.map((zIndex, idx) => {
        const isCover = idx === 0 || idx === pages.length - 1;
        const post = postPages[idx];

        return (
          <label className="book-inner" key={idx}>
            <input
              className="book-inner__flip"
              type="checkbox"
              onChange={playAudio}
            />
            <span
              className={`book-inner__page z-[${zIndex}] ${
                isCover ? "front-cover" : ""
              } p-4`}
            >
              {isCover ? (
                <p className="book-inner__cover-text">Your Diary</p>
              ) : post ? (
                <div className="flex flex-col h-full">
                  <div className="flex justify-between items-start mb-4">
                    <div className="flex-1 pr-4 text-left">
                      <div className="text-sm text-gray-600 dark:text-gray-400 pb-1 border-b border-gray-400 dark:border-gray-600">
                        Date:{" "}
                        {new Date(post.createdAt).toLocaleDateString("ja-JP")}
                      </div>
                      <div className="text-lg font-semibold text-gray-800 dark:text-gray-100 mt-2 pb-1 border-b border-gray-400 dark:border-gray-600">
                        Subject: {post.title}
                      </div>
                      <div className="h-4 border-b border-gray-400 dark:border-gray-600"></div>
                      <div className="h-4 border-b border-gray-400 dark:border-gray-600"></div>
                    </div>
                    {post.tracks?.length > 0 && (
                      <div className="flex-shrink-0 w-1/3">
                        <ul className="space-y-1">
                          {post.tracks.map((track) => (
                            <li
                              key={track.spotifyId}
                              className="flex items-center gap-2 p-1 border rounded-md bg-gray-50 dark:bg-gray-700"
                            >
                              <Image
                                src={track.albumImageUrl || "/default-image.jpg"}
                                alt={track.name || ""}
                                width={32}
                                height={32}
                                className="w-8 h-8 rounded object-cover"
                                style={{ width: "auto", height: "auto" }}
                              />
                              <div>
                                <p className="text-xs font-medium text-gray-900 dark:text-gray-100">
                                  {track.name}
                                </p>
                                <p className="text-xs text-gray-600 dark:text-gray-300">
                                  {track.artistName}
                                </p>
                              </div>
                            </li>
                          ))}
                        </ul>
                      </div>
                    )}
                  </div>
                  <div className="flex-grow overflow-hidden relative">
                    <div className="absolute inset-0 z-0 pointer-events-none">
                      {Array.from({ length: 15 }).map((_, i) => (
                        <div
                          key={i}
                          className="h-6 border-b border-gray-400 dark:border-gray-600"
                        ></div>
                      ))}
                    </div>
                    <p className="relative z-10 text-sm text-gray-700 dark:text-gray-300 leading-6 whitespace-pre-wrap h-full overflow-y-auto text-left">
                      {post.description}
                    </p>
                  </div>
                </div>
              ) : (
                <div className="flex flex-col h-full">
                  <div className="flex justify-between items-start mb-4">
                    <div className="flex-1 pr-4 text-left">
                      <div className="text-sm text-gray-600 dark:text-gray-400 pb-1 border-b border-gray-400 dark:border-gray-600">
                        Date:
                      </div>
                      <div className="text-lg font-semibold text-gray-800 dark:text-gray-100 mt-2 pb-1 border-b border-gray-400 dark:border-gray-600">
                        Subject:
                      </div>
                      <div className="h-4 border-b border-gray-400 dark:border-gray-600"></div>
                      <div className="h-4 border-b border-gray-400 dark:border-gray-600"></div>
                    </div>
                  </div>
                  <div className="flex-grow overflow-hidden relative">
                    <div className="absolute inset-0 z-0 pointer-events-none">
                      {Array.from({ length: 15 }).map((_, i) => (
                        <div
                          key={i}
                          className="h-6 border-b border-gray-400 dark:border-gray-600"
                        ></div>
                      ))}
                    </div>
                    <p className="relative z-10 text-sm text-gray-700 dark:text-gray-300 leading-6 whitespace-pre-wrap h-full overflow-y-auto text-left"></p>
                  </div>
                </div>
              )}
            </span>
            <span
              className={`book-inner__dummy dummy ${
                isCover ? "back-cover" : ""
              } p-4`}
            >
              {isCover ? (
                <p className="book-inner__cover-text">Your Diary</p>
              ) : post ? (
                <div className="flex flex-col h-full">
                  <div className="flex justify-between items-start mb-4">
                    <div className="flex-1 pr-4 text-left">
                      <div className="text-sm text-gray-600 dark:text-gray-400 pb-1 border-b border-gray-400 dark:border-gray-600">
                        Date:{" "}
                        {new Date(post.createdAt).toLocaleDateString("ja-JP")}
                      </div>
                      <div className="text-lg font-semibold text-gray-800 dark:text-gray-100 mt-2 pb-1 border-b border-gray-400 dark:border-gray-600">
                        Subject: {post.title}
                      </div>
                      <div className="h-4 border-b border-gray-400 dark:border-gray-600"></div>
                      <div className="h-4 border-b border-gray-400 dark:border-gray-600"></div>
                    </div>
                    {post.tracks?.length > 0 && (
                      <div className="flex-shrink-0 w-1/3">
                        <ul className="space-y-1">
                          {post.tracks.map((track) => (
                            <li
                              key={track.spotifyId}
                              className="flex items-center gap-2 p-1 border rounded-md bg-gray-50 dark:bg-gray-700"
                            >
                              <Image
                                src={track.albumImageUrl || "/default-image.jpg"}
                                alt={track.name || ""}
                                width={32}
                                height={32}
                                className="w-8 h-8 rounded object-cover"
                              />
                              <div>
                                <p className="text-xs font-medium text-gray-900 dark:text-gray-100">
                                  {track.name}
                                </p>
                                <p className="text-xs text-gray-600 dark:text-gray-300">
                                  {track.artistName}
                                </p>
                              </div>
                            </li>
                          ))}
                        </ul>
                      </div>
                    )}
                  </div>
                  <div className="flex-grow overflow-hidden relative">
                    <div className="absolute inset-0 z-0 pointer-events-none">
                      {Array.from({ length: 15 }).map((_, i) => (
                        <div
                          key={i}
                          className="h-6 border-b border-gray-400 dark:border-gray-600"
                        ></div>
                      ))}
                    </div>
                    <p className="relative z-10 text-sm text-gray-700 dark:text-gray-300 leading-6 whitespace-pre-wrap h-full overflow-y-auto text-left">
                      {post.description}
                    </p>
                  </div>
                </div>
              ) : (
                <div className="flex flex-col h-full">
                  <div className="flex justify-between items-start mb-4">
                    <div className="flex-1 pr-4 text-left">
                      <div className="text-sm text-gray-600 dark:text-gray-400 pb-1 border-b border-gray-400 dark:border-gray-600">
                        Date:
                      </div>
                      <div className="text-lg font-semibold text-gray-800 dark:text-gray-100 mt-2 pb-1 border-b border-gray-400 dark:border-gray-600">
                        Subject:
                      </div>
                      <div className="h-4 border-b border-gray-400 dark:border-gray-600"></div>
                      <div className="h-4 border-b border-gray-400 dark:border-gray-600"></div>
                    </div>
                  </div>
                  <div className="flex-grow overflow-hidden relative">
                    <div className="absolute inset-0 z-0 pointer-events-none">
                      {Array.from({ length: 15 }).map((_, i) => (
                        <div
                          key={i}
                          className="h-6 border-b border-gray-400 dark:border-gray-600"
                        ></div>
                      ))}
                    </div>
                    <p className="relative z-10 text-sm text-gray-700 dark:text-gray-300 leading-6 whitespace-pre-wrap h-full overflow-y-auto text-left"></p>
                  </div>
                </div>
              )}
            </span>
          </label>
        );
      })}
    </div>
  );
};
