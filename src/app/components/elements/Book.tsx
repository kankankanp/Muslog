"use client";

import React, { useRef, useState, useEffect, ForwardedRef } from "react";
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
  <div className="page page-cover" ref={ref} data-density="hard">
    <div className="page-content">
      <h2>{children}</h2>
    </div>
  </div>
));

// eslint-disable-next-line react/display-name
const Page = React.forwardRef<HTMLDivElement, PageProps>(
  ({ number, children }, ref) => (
    <div className="page" ref={ref}>
      <div className="page-content">
        <h2 className="page-header">Page header - {number}</h2>
        <div className="page-image"></div>
        <div className="page-text">{children}</div>
        <div className="page-footer">{(number ?? 0) + 1}</div>
      </div>
    </div>
  )
);

export default function DemoBook({ posts }: { posts: PostType[] }) {
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

  useEffect(() => {
    if (flipBook.current) {
      const pageCount = flipBook.current.pageFlip()?.getPageCount();
      if (pageCount) setTotalPage(pageCount);
    }
  }, []);

  return (
    <div>
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
        className="demo-book"
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
        {/* 表紙 */}
        <PageCover>BOOK TITLE</PageCover>

        {/* 投稿ごとにページを作成 */}
        {posts.map((post, index) => (
          <Page number={index + 1} key={post.id}>
            img
            <div className="p-4">
              {/* BlogCard から1投稿だけ表示 */}
              <div className="p-4 sm:p-6 bg-white dark:bg-gray-800 shadow-md rounded-lg">
                <h3 className="text-lg sm:text-xl font-semibold mt-2 text-gray-900 dark:text-gray-100">
                  {post.title}
                </h3>
                <p className="text-base sm:text-lg mt-2 text-gray-700 dark:text-gray-300">
                  {post.description}
                </p>
              </div>
            </div>
          </Page>
        ))}

        {/* 裏表紙 */}
        <PageCover>THE END</PageCover>
      </HTMLFlipBook>

      <div className="container mt-4 flex justify-center gap-4">
        <button type="button" onClick={prevButtonClick}>
          Previous page
        </button>
        [<span>{page}</span> of <span>{totalPage}</span>]
        <button type="button" onClick={nextButtonClick}>
          Next page
        </button>
      </div>
    </div>
  );
}
