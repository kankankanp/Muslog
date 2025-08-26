import React from 'react';
import Link from 'next/link';
import { components } from '@/app/libs/api/generated/openapi-types';

type Community = components['schemas']['Community'];

interface CommunityCardProps {
  community: Community;
}

const CommunityCard: React.FC<CommunityCardProps> = ({ community }) => {
  return (
    <Link href={`/community/${community.id}`}>
      <div className="p-4 border rounded-lg shadow-md hover:shadow-lg transition-shadow duration-200 cursor-pointer">
        <h3 className="text-xl font-semibold text-gray-800 dark:text-white">{community.name}</h3>
        <p className="text-gray-600 dark:text-gray-300 mt-2">{community.description}</p>
        <p className="text-sm text-gray-500 dark:text-gray-400 mt-2">
          Created by {community.creatorId} on {new Date(community.createdAt).toLocaleDateString()}
        </p>
      </div>
    </Link>
  );
};

export default CommunityCard;
