import { PostType } from "@/types";
import "@/scss/blog-card.scss";
import Link from "next/link";

// type BlogCardProps = {
//   isDetailPage?: boolean;
//   posts: PostType[];
// };

const BlogCard = async ({ isDetailPage, posts }: any) => {
  const safePosts = Array.isArray(posts) ? posts : [];
  const truncateText = (text: string, length: number) => {
    if (text.length <= length) {
      return text;
    }
    return text.substring(0, length) + "...";
  };

  return (
    <div className={!isDetailPage ? "blog" : "blog__detail"}>
      {safePosts.map((post: PostType) => {
        const date = new Date(post.date);
        const year = date.getFullYear();
        const month = date.getMonth() + 1;
        const day = date.getDate();
        const daysOfWeek = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
        const dayOfWeek = daysOfWeek[date.getDay()];

        return (
          <div key={post.id} className="blog__item">
            <h3 className="blog__date">
              {`${year}/${month}/${day}(${dayOfWeek})`}
            </h3>
            <h3 className="blog__title">{post.title}</h3>
            <p className="blog__text">
              {isDetailPage
                ? post.description
                : truncateText(post.description, 40)}
            </p>
            {!isDetailPage && (
              <Link href={`/blog/${post.id}`} className="blog__btn">
                <span></span>show more
              </Link>
            )}
          </div>
        );
      })}
    </div>
  );
};

export default BlogCard;
