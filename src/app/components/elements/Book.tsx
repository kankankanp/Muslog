"use client";

import React, { useRef, useState, useEffect } from "react";
import HTMLFlipBook from "react-pageflip";
import { PostType } from "./BlogCard";

type PageProps = {
  number?: number;
  children: React.ReactNode;
};

// eslint-disable-next-line react/display-name
const PageCover = React.forwardRef<
  HTMLDivElement,
  { children: React.ReactNode }
>(({ children }, ref) => (
  <div
    ref={ref}
    data-density="hard"
    className="flex items-center justify-center bg-gradient-to-br from-gray-700 to-gray-900"
  >
    <h2 className="text-3xl font-bold text-white">{children}</h2>
  </div>
));

// eslint-disable-next-line react/display-name
const Page = React.forwardRef<HTMLDivElement, PageProps>(
  ({ number, children }, ref) => {
    return (
      <div ref={ref}>
        <div className="absolute inset-0 bg-black/30" />
        <div className="relative z-10 p-6 flex flex-col justify-between h-full">
          <h2 className="text-xl font-bold text-white">Page {number}</h2>
          <div className="flex-grow flex items-center justify-center">
            {children}
          </div>
          <div className="text-sm text-white text-right">
            {number ? number + 1 : ""}
          </div>
        </div>
      </div>
    );
  }
);

export default function Book({ posts }: { posts: PostType[] }) {
  const flipBook = useRef<any>(null);
  const [page, setPage] = useState(0);
  const [totalPage, setTotalPage] = useState(0);

  const nextButtonClick = () => {
    flipBook.current?.pageFlip()?.flipNext();
  };

  const prevButtonClick = () => {
    flipBook.current?.pageFlip()?.flipPrev();
  };

  const onPage = (e: { data: number }) => {
    setPage(e.data);
  };

  const onInit = () => {
    if (flipBook.current) {
      const pageCount = flipBook.current.pageFlip()?.getPageCount();
      if (pageCount) setTotalPage(pageCount);
    }
  };

  return (
    <div className="flex flex-col items-center gap-6 py-8">
      <HTMLFlipBook
        width={550}
        height={733}
        size="stretch"
        minWidth={315}
        maxWidth={1000}
        minHeight={400}
        maxHeight={1533}
        maxShadowOpacity={0.5}
        showCover={true}
        mobileScrollSupport={true}
        onFlip={onPage}
        onInit={onInit}
        className="shadow-2xl"
        ref={flipBook}
        startPage={0}
        drawShadow={false}
        flippingTime={100}
        usePortrait={false}
        startZIndex={0}
        autoSize={false}
        clickEventForward={false}
        useMouseEvents={false}
        swipeDistance={0}
        showPageCorners={false}
        disableFlipByClick={false}
        style={{}}
      >
        <PageCover>BOOK TITLE</PageCover>

        {posts.map((post, index) => (
          <Page number={index + 1} key={post.id}>
            <div className="bg-white/80 dark:bg-gray-800/80 p-6 rounded-lg shadow-lg max-w-full">
              <h3 className="text-xl font-bold text-gray-900 dark:text-gray-100">
                {post.title}
              </h3>
              <p className="mt-4 text-base text-gray-700 dark:text-gray-300">
                {post.description}
              </p>
            </div>
          </Page>
        ))}

        <PageCover>THE END</PageCover>
      </HTMLFlipBook>

      <div className="flex items-center gap-4">
        <button
          type="button"
          onClick={prevButtonClick}
          className={`px-4 py-2 bg-gray-800 text-white rounded ${
            page === 0 ? "opacity-15" : "opacity-100"
          }`}
        >
          Previous page
        </button>
        <span className="text-lg">
          [{page} / {totalPage}]
        </span>
        <button
          type="button"
          onClick={nextButtonClick}
          className={`px-4 py-2 bg-gray-800 text-white rounded ${
            page ===  totalPage - 2 ? "opacity-15" : "opacity-100"
          }`}        >
          Next page
        </button>
      </div>
    </div>
  );
}
