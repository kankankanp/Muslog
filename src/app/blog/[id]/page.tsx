import Header from "@/app/components/layouts/header/page";
import Footer from "@/app/components/layouts/footer/page";
import BlogCard from "@/app/components/elements/blog-card/page";
import { Anybody } from "next/font/google";

const showBlogDetails = async (id: number) => {
  const res = await fetch(`http://localhost:3000/api/blog/${id}`);
  const data = await res.json();
  return data.post;
};

const ShowBlogDetails = async ({ params }: { params: { id: number } }) => {
  const post = await showBlogDetails(params.id);
  const postarray: any = [post];

  return (
    <>
      <Header />
      <main>
        <BlogCard isDetailPage={true} posts={postarray} />
      </main>
      <Footer />
    </>
  );
};

export default ShowBlogDetails;
