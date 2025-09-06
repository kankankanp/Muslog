"use client";

import { Search } from "lucide-react";
import { useRouter, useSearchParams } from "next/navigation";
import React, { useState, useEffect } from "react";
import CommunityCard from "@/components/community/CommunityCard";
import CreateCommunityModal from "@/components/elements/modals/CreateCommunityModal";
import Spinner from "@/components/layouts/Spinner";
import { useGetCommunitiesSearch } from "@/libs/api/generated/orval/communities/communities";

const CommunityPage: React.FC = () => {
  const router = useRouter();
  const searchParams = useSearchParams();
  const initialSearchQuery = searchParams.get("q") || "";
  const initialPage = parseInt(searchParams.get("page") || "1", 10);

  const [searchQuery, setSearchQuery] = useState(initialSearchQuery);
  const [currentInput, setCurrentInput] = useState(initialSearchQuery);
  const [currentPage, setCurrentPage] = useState(initialPage);

  useEffect(() => {
    const params = new URLSearchParams(searchParams.toString());
    if (searchQuery) {
      params.set("q", searchQuery);
    } else {
      params.delete("q");
    }
    params.set("page", currentPage.toString());
    router.push(`?${params.toString()}`);
  }, [searchQuery, currentPage, router, searchParams]);

  const {
    data: searchResult,
    isLoading,
    isError,
    error,
    refetch,
  } = useGetCommunitiesSearch({
    q: searchQuery,
    page: currentPage,
    perPage: 10,
  });

  const [isModalOpen, setIsModalOpen] = useState(false);

  if (isLoading) {
    return <Spinner />;
  }

  if (isError) {
    return (
      <div className="text-center text-red-600">
        Error: {typeof error !== "undefined" && error ? String(error) : "Failed to load communities"}
      </div>
    );
  }

  const communities = searchResult?.communities || [];
  const totalCount = searchResult?.totalCount || 0;
  const perPage = searchResult?.perPage || 10;
  const totalPages = Math.ceil(totalCount / perPage);

  const handleSearch = () => {
    setSearchQuery(currentInput);
    setCurrentPage(1); // Reset to first page on new search
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      handleSearch();
    }
  };

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
  };

  return (
    <>
      <div className="border-gray-100 border-b-2 bg-white px-8 py-6 flex justify-between max-md:flex-col max-md:py-2 max-md:gap-2">
        <h1 className="text-3xl font-bold">コミュニティ</h1>
        <button
          onClick={() => {
            console.log(isModalOpen);
            setIsModalOpen(true);
          }}
          className="py-1 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 dark:bg-indigo-500 dark:hover:bg-indigo-600 text-center"
        >
          コミュニティを作成する
        </button>
      </div>
      <div className="relative mt-4 max-w-lg mx-auto w-4/5">
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
      <div className="container mx-auto p-4">
        <div className="w-3/5 mx-auto mb-8"></div>

        <h2 className="text-2xl font-semibold text-gray-900 dark:text-white mb-4">
          参加可能なコミュニティ
        </h2>
        {communities.length === 0 ? (
          <p className="text-gray-600 dark:text-gray-300">
            No communities found. Be the first to create one!
          </p>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {communities.map((community) => (
              <CommunityCard key={community.id} community={community} />
            ))}
          </div>
        )}
      </div>
      {totalPages > 1 && (
        <div className="flex justify-center mt-8 pb-8">
          {Array.from({ length: totalPages }, (_, i) => i + 1).map((page) => (
            <button
              key={page}
              onClick={() => handlePageChange(page)}
              className={`mx-1 px-3 py-1 rounded ${
                currentPage === page
                  ? "bg-blue-500 text-white"
                  : "bg-gray-200 text-gray-700"
              }`}
            >
              {page}
            </button>
          ))}
        </div>
      )}
      <CreateCommunityModal
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(false)}
        onCommunityCreated={refetch}
      />
    </>
  );
};

export default CommunityPage;
