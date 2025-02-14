import { PostType } from "@/app/type/PostType";
import "@/scss/blog-card.scss";
import Link from "next/link";
import { faPen } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

type BlogCardProps = {
  isDetailPage?: boolean;
  posts: PostType[];
};

/* メモ：非同期コンポーネントでないかつchildrenを使用したい場合は
        BlogCard: FC<BlogCardProps>を使用すべき */
const BlogCard = ({ isDetailPage, posts }: BlogCardProps) => {
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
            <div className="blog__btn-area">
              {!isDetailPage && (
                <Link href={`/dashboard/blog/edit/${post.id}`} className="blog__edit-btn">
                  <span></span>
                  <FontAwesomeIcon icon={faPen} />
                </Link>
              )}
              {!isDetailPage && (
                <Link href={`/dashboard/blog/${post.id}`} className="blog__detail-btn">
                  <span></span>show more
                </Link>
              )}
            </div>
          </div>
        );
      })}
    </div>
  );
};

export default BlogCard;
