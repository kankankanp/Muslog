"use client";

import React from "react";
import CommunityCard from "@/components/community/CommunityCard";
import CreateCommunityForm from "@/components/community/CreateCommunityForm";
import Spinner from "@/components/layouts/Spinner";
import { useGetCommunities } from "@/libs/api/generated/orval/communities/communities";

const CommunityPage: React.FC = () => {
  const { data, isLoading, isError, error, refetch } = useGetCommunities();

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
      <div className="container mx-auto p-4">
        <div className="mb-8">
          <CreateCommunityForm onCommunityCreated={refetch} />
        </div>

        <h2 className="text-2xl font-semibold text-gray-900 dark:text-white mb-4">
          Available Communities
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
    </>
  );
};

export default CommunityPage;
