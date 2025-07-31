import { faPen, faHeart } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Image from "next/image";
import Link from "next/link";
import { CommonButton } from "../buttons/CommonButton";
import { Post } from "@/app/libs/api/generated/orval/model/post";

type BlogCardProps = {
  isDetailPage?: boolean;
  posts: Post[];
  isLiked?: boolean;
  onLikeClick?: () => void;
};

const BlogCard = ({
  isDetailPage,
  posts,
  isLiked,
  onLikeClick,
}: BlogCardProps) => {
  const safePosts = Array.isArray(posts) ? posts : [];
  const truncateText = (text: string, length: number) => {
    if (text.length <= length) {
      return text;
    }
    return text.substring(0, length) + "...";
  };

  return (
    <div
      className={`py-8 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 grid ${
        !isDetailPage
          ? "grid-cols-1 sm:grid-cols-2 gap-6 w-full"
          : "grid-cols-1 w-3/5"
      }`}
    >
      {safePosts.map((post: Post) => {
        // const date = new Date(post.date);
        // const year = date.getFullYear();
        // const month = date.getMonth() + 1;
        // const day = date.getDate();
        // const daysOfWeek = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
        // const dayOfWeek = daysOfWeek[date.getDay()];

        return (
          <div
            key={post.id}
            className="p-4 sm:p-6 bg-white dark:bg-gray-800 shadow-md rounded-lg"
          >
            {/* <h3 className="text-xl sm:text-2xl font-bold text-gray-900 dark:text-gray-100">{`${year}/${month}/${day}(${dayOfWeek})`}</h3> */}
            <h3 className="text-lg sm:text-xl font-semibold mt-2 text-gray-900 dark:text-gray-100">
              {post.title}
            </h3>
            <p className="text-base sm:text-lg mt-2 text-gray-700 dark:text-gray-300">
              {isDetailPage
                ? post.description
                : truncateText(post.description, 40)}
            </p>

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

            {post.tracks?.length > 0 && (
              <div className="mt-4 space-y-3">
                <ul className="space-y-2">
                  {post.tracks?.map((track) => (
                    <li
                      key={track.spotifyId}
                      className="flex items-center gap-4 p-2 border rounded-md bg-gray-50 dark:bg-gray-700"
                    >
                      <Image
                        src={track.albumImageUrl || "/default-image.jpg"}
                        alt={track.name || ""}
                        width={48}
                        height={48}
                        className="w-12 h-12 rounded object-cover"
                        style={{ width: "auto", height: "auto" }}
                      />
                      <div>
                        <p className="text-sm font-medium text-gray-900 dark:text-gray-100">
                          {track.name}
                        </p>
                        <p className="text-sm text-gray-600 dark:text-gray-300">
                          {track.artistName}
                        </p>
                      </div>
                    </li>
                  ))}
                </ul>
              </div>
            )}

            <div className="flex justify-between items-center mt-4">
              <div className="flex items-center">
                <FontAwesomeIcon
                  icon={faHeart}
                  className={`cursor-pointer mr-1 ${
                    isLiked ? "text-red-500" : "text-gray-400"
                  }`}
                  onClick={onLikeClick}
                />
                <span className="text-gray-700 dark:text-gray-300">
                  {post.likesCount}
                </span>
              </div>
              {!isDetailPage && (
                <Link
                  href={`/dashboard/blog/edit/${post.id}`}
                  className="text-black dark:text-gray-300 text-base sm:text-lg transition-transform transform hover:scale-125"
                >
                  <FontAwesomeIcon icon={faPen} />
                </Link>
              )}
              {isDetailPage ? (
                <CommonButton href={`/dashboard/blog/page/1`}>
                  Back
                </CommonButton>
              ) : (
                <CommonButton href={`/dashboard/blog/${post.id}`}>
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
