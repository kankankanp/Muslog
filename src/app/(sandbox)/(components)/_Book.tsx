"use client";

import { PostType } from "@/app/components/elements/BlogCard";
import React, { useRef, useState, useEffect } from "react";
import HTMLFlipBook from "react-pageflip";

type PageProps = {
  number?: number;
  children: React.ReactNode;
};

const BASE_WIDTH = 500;
const BASE_HEIGHT = 733;
const ASPECT_RATIO = BASE_HEIGHT / BASE_WIDTH;
const MOBILE_BREAKPOINT = 768;
const MOBILE_WIDTH = 340;

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
        <div />
        <div className="relative z-10 p-6 flex flex-col justify-between h-full">
          <h2 className="text-xl font-bold">Page {number}</h2>
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
  const [bookWidth, setBookWidth] = useState(BASE_WIDTH);
  const [bookHeight, setBookHeight] = useState(BASE_HEIGHT);

  useEffect(() => {
    const handleResize = () => {
      const isMobile = window.innerWidth <= MOBILE_BREAKPOINT;
      const newWidth = isMobile ? MOBILE_WIDTH : BASE_WIDTH;
      setBookWidth(newWidth);
      setBookHeight(Math.round(newWidth * ASPECT_RATIO));
    };

    handleResize(); // 初回チェック
    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, []);

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
      const pageCount = flipBook.current.pageFlip()?.getPageCount() - 1;
      if (pageCount) setTotalPage(pageCount);
    }
  };

  return (
    <div className="flex flex-col items-center gap-6 py-8">
      <HTMLFlipBook
        width={bookWidth}
        height={bookHeight}
        size="stretch"
        minWidth={bookWidth}
        maxWidth={1000}
        minHeight={bookHeight}
        maxHeight={1533}
        maxShadowOpacity={0.5}
        showCover={true}
        mobileScrollSupport={true}
        onFlip={onPage}
        onInit={onInit}
        className="shadow-2xl"
        ref={flipBook}
        startPage={0}
        drawShadow={true}
        flippingTime={500}
        usePortrait={true}
        startZIndex={0}
        autoSize={true}
        clickEventForward={false}
        useMouseEvents={false}
        swipeDistance={0}
        showPageCorners={true}
        disableFlipByClick={false}
        style={{}}
      >
        <PageCover>BOOK TITLE</PageCover>

        {posts.map((post, index) => (
          <Page number={index + 1} key={post.id}>
            <div>
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
            page === totalPage ? "opacity-15" : "opacity-100"
          }`}
        >
          Next page
        </button>
      </div>
    </div>
  );
}
