import Header from "../components/layouts/header/page";
import Footer from "../components/layouts/footer/page";
import BlogCard from "../components/elements/blog-card/page";
import AddButton from "../components/elements/add-button/page";

const fetchAllBlogs = async () => {
  const res = await fetch("http://localhost:3000/api/blog", {
    cache: "no-store",
  });

  const data = await res.json();

  return data.posts;
};

const Blog = async () => {
  const posts = await fetchAllBlogs();

  return (
    <>
      <Header />
      <main>
        <AddButton />
        <BlogCard isDetailPage={false} posts={posts} />
      </main>
      <Footer />
    </>
  );
};

export default Blog;
