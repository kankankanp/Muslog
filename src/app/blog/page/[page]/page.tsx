import Pagination from "@/app/components/elements/pagination/page";
import Header from "@/app/components/layouts/header/page";
import AddBlogModal from "@/app/components/elements/add-blog-modal/page";
import BlogCard from "@/app/components/elements/blog-card/page";
import Footer from "@/app/components/layouts/footer/page";
import { fetchBlogsByPage } from "../../utils";

export default async function Index({ params }: { params: { page: number } }) {
  const pageIndex = params.page;
  const { posts, totalCount } = await fetchBlogsByPage(pageIndex);

  return (
    <>
      <Header />
      <main>
        <AddBlogModal />
        <BlogCard isDetailPage={false} posts={posts} />
      </main>
      <Pagination totalCount={totalCount} />
      <Footer />
    </>
  );
}
