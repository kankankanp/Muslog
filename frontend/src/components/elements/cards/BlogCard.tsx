import { faHeart } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Image from "next/image";
import { CommonButton } from "../buttons/CommonButton";
import { Post } from "@/libs/api/generated/orval/model/post";

type BlogCardProps = {
  isDetailPage?: boolean;
  posts: Post[];
  isLiked?: boolean;
  onLikeClick?: () => void;
  onLikeToggle?: (postId: number, isCurrentlyLiked: boolean) => void;
};

const BlogCard = ({
  isDetailPage,
  posts,
  isLiked,
  onLikeClick,
  onLikeToggle,
}: BlogCardProps) => {
  const safePosts = Array.isArray(posts) ? posts : [];

  return (
    <div
      className={`py-8 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 grid ${
        !isDetailPage
          ? "grid-cols-1 sm:grid-cols-2 gap-6 w-full"
          : "grid-cols-1 w-3/5"
      }`}
    >
      {safePosts.map((post: Post) => {
        const headerImageSrc =
          post.headerImageUrl && post.headerImageUrl.trim() !== ""
            ? post.headerImageUrl
            : "/default-image.jpg";
        return (
          <div
            key={post.id}
            className="p-4 sm:p-6 bg-white dark:bg-gray-800 shadow-md rounded-lg"
          >
            <Image
              src={headerImageSrc}
              alt="Header Image"
              width={600}
              height={300}
              className="w-full h-48 object-cover rounded-md mb-4"
            />
            <h3 className="text-lg sm:text-xl font-semibold mt-2 text-gray-900 dark:text-gray-100">
              {post.title}
            </h3>

            {post.tags && post.tags.length > 0 && (
              <div className="mt-2 flex flex-wrap gap-2">
                {post.tags.map((tag) => (
                  <span
                    key={tag.id}
                    className="bg-blue-100 text-blue-800 text-xs font-medium px-2.5 py-0.5 rounded-full dark:bg-blue-900 dark:text-blue-300"
                  >
                    {tag.name}
                  </span>
                ))}
              </div>
            )}

            <div className="flex justify-between items-center mt-4">
              <div className="flex items-center">
                <FontAwesomeIcon
                  icon={faHeart}
                  className={`cursor-pointer mr-1 ${
                    isDetailPage
                      ? isLiked
                        ? "text-red-500"
                        : "text-gray-400"
                      : post.isLiked
                        ? "text-red-500"
                        : "text-gray-400"
                  }`}
                  onClick={
                    isDetailPage
                      ? onLikeClick
                      : () => onLikeToggle?.(post.id, post.isLiked ?? false)
                  }
                />
                <span className="text-gray-700 dark:text-gray-300">
                  {post.likesCount}
                </span>
              </div>
              {isDetailPage ? (
                <CommonButton href={`/dashboard/blog/page/1`}>
                  Back
                </CommonButton>
              ) : (
                <CommonButton href={`/dashboard/post/${post.id}`}>
                  Show more
                </CommonButton>
              )}
            </div>
          </div>
        );
      })}
    </div>
  );
};

export default BlogCard;
