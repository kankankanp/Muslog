"use client";

import { Search } from "lucide-react";
import React, { useState } from "react";
import CommunityCard from "@/components/community/CommunityCard";
import CreateCommunityModal from "@/components/elements/modals/CreateCommunityModal";
import Spinner from "@/components/layouts/Spinner";
import { useGetCommunities } from "@/libs/api/generated/orval/communities/communities";

const CommunityPage: React.FC = () => {
  const { data, isLoading, isError, error, refetch } = useGetCommunities();
  const [isModalOpen, setIsModalOpen] = useState(false);

  if (isLoading) {
    return <Spinner />;
  }

  if (isError) {
    const errorMessage =
      error && typeof error === "object" && "message" in error
        ? (error as { message: string }).message
        : "Unknown error";
    return (
      <div className="text-center text-red-600">
        Error: {errorMessage || "Failed to load communities"}
      </div>
    );
  }

  const communities = data?.communities || [];

  return (
    <>
      <h1 className="text-3xl font-bold border-gray-100 border-b-2 bg-white px-6 py-6">
        コミュニティ
      </h1>
      <div className="relative mt-4 max-w-lg mx-auto">
        <input
          type="text"
          placeholder="検索"
          className="w-full pl-4 pr-10 py-2 rounded-full text-gray-200 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-purple-500 border-gray-200 border-2"
        />
        <Search className="absolute right-3 top-1/2 -translate-y-1/2 h-5 w-5 text-gray-400" />
      </div>
      <div className="container mx-auto p-4">
        <div className="w-3/5 mx-auto mb-8">
          <button
            onClick={() => {
              console.log(isModalOpen);
              setIsModalOpen(true);
            }}
            className="py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 dark:bg-indigo-500 dark:hover:bg-indigo-600"
          >
            コミュニティを作成する
          </button>
        </div>

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
      <CreateCommunityModal
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(false)}
        onCommunityCreated={refetch}
      />
    </>
  );
};

export default CommunityPage;
