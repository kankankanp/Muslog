'use client';

import { useQueryClient } from '@tanstack/react-query';
import { Search } from 'lucide-react';
import { useRouter, useSearchParams } from 'next/navigation';
import { useState, useEffect } from 'react';
import { useSelector } from 'react-redux';
import BlogCard from '@/components/elements/cards/BlogCard';
import Spinner from '@/components/layouts/Spinner';
import {
  usePostPostsPostIDLike,
  useDeletePostsPostIDUnlike,
} from '@/libs/api/generated/orval/likes/likes';
import { useGetPostsSearch } from '@/libs/api/generated/orval/posts/posts';
import { RootState } from '@/libs/store/store';

export default function Page() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const initialSearchQuery = searchParams.get('q') || '';
  const initialPage = parseInt(searchParams.get('page') || '1', 10);

  const [searchQuery, setSearchQuery] = useState(initialSearchQuery);
  const [currentInput, setCurrentInput] = useState(initialSearchQuery);
  const [currentPage, setCurrentPage] = useState(initialPage);

  useEffect(() => {
    const params = new URLSearchParams(searchParams.toString());
    if (searchQuery) {
      params.set('q', searchQuery);
    } else {
      params.delete('q');
    }
    params.set('page', currentPage.toString());
    router.push(`?${params.toString()}`);
  }, [searchQuery, currentPage, router, searchParams]);

  const {
    data: searchResult,
    isPending,
    error,
  } = useGetPostsSearch({
    q: searchQuery,
    page: currentPage,
    perPage: 10,
  });

  const posts = searchResult?.posts || [];
  const totalCount = searchResult?.totalCount || 0;
  const perPage = searchResult?.perPage || 10;
  const totalPages = Math.ceil(totalCount / perPage);
  const user = useSelector((state: RootState) => state.auth.user);
  const queryClient = useQueryClient();

  const { mutate: likePost } = usePostPostsPostIDLike({
    mutation: {
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ['/posts'] });
      },
      onError: (error) => {
        console.error('Failed to like post:', error);
        alert('いいねに失敗しました');
      },
    },
  });

  const { mutate: unlikePost } = useDeletePostsPostIDUnlike({
    mutation: {
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ['/posts'] });
      },
      onError: (error) => {
        console.error('Failed to unlike post:', error);
        alert('いいね解除に失敗しました');
      },
    },
  });

  const handleLikeToggle = (postId: number, isCurrentlyLiked: boolean) => {
    if (!user) {
      alert('ログインしてください');
      return;
    }

    if (isCurrentlyLiked) {
      unlikePost({ postID: postId });
    } else {
      likePost({ postID: postId });
    }
  };

  if (isPending) return <Spinner />;
  const handleSearch = () => {
    setSearchQuery(currentInput);
    setCurrentPage(1); // Reset to first page on new search
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') {
      handleSearch();
    }
  };

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
  };

  if (isPending) return <Spinner />;
  if (error) return <div>Error loading posts: {error}</div>;

  return (
    <div className="dark:bg-gray-900 min-h-screen">
      <h1 className="border-gray-100 border-b-2 bg-white px-8 py-6 flex items-center justify-between">
        <div className="text-gray-800 text-3xl font-bold">ホーム</div>
      </h1>
      <div className="relative mt-4 max-w-lg mx-auto max-md:w-4/5">
        <input
          type="text"
          placeholder="検索"
          className="w-full pl-4 pr-10 py-2 rounded-full text-gray-800 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-purple-500 border-gray-200 border-2"
          value={currentInput}
          onChange={(e) => setCurrentInput(e.target.value)}
          onKeyDown={handleKeyDown}
        />
        <Search
          className="absolute right-3 top-1/2 -translate-y-1/2 h-5 w-5 text-gray-400 cursor-pointer"
          onClick={handleSearch}
        />
      </div>
      {searchQuery && (
        <div className="text-center mt-4 text-gray-600 dark:text-gray-300">
          {totalCount}件の検索結果
          {totalCount === 0 && (
            <p className="mt-2">記事が見つかりませんでした。</p>
          )}
        </div>
      )}
      <BlogCard posts={posts || []} onLikeToggle={handleLikeToggle} />
      {/* Pagination Component will go here */}
      {totalPages > 1 && (
        <div className="flex justify-center mt-8 pb-8">
          {Array.from({ length: totalPages }, (_, i) => i + 1).map((page) => (
            <button
              key={page}
              onClick={() => handlePageChange(page)}
              className={`mx-1 px-3 py-1 rounded ${
                currentPage === page
                  ? 'bg-blue-500 text-white'
                  : 'bg-gray-200 text-gray-700'
              }`}
            >
              {page}
            </button>
          ))}
        </div>
      )}
    </div>
  );
}
