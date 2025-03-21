"use client";

import { useCallback, useMemo } from "react";
import type { Post } from "./BlogCard";
import "@/scss/book.scss";

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
              }`}
            >
              {isCover ? (
                "表紙"
              ) : post ? (
                <div>
                  <h2>{post.title}</h2>
                  <p>{post.description}</p>
                </div>
              ) : (
                "None"
              )}
            </span>
            <span
              className={`book-inner__dummy dummy ${
                isCover ? "back-cover" : ""
              }`}
            >
              {isCover ? (
                "表紙"
              ) : post ? (
                <div>
                  <h2>{post.title}</h2>
                  <p>{post.description}</p>
                </div>
              ) : (
                "None"
              )}
            </span>
          </label>
        );
      })}
    </div>
  );
};
