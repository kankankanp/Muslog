import Header from "../components/layouts/header/page";
import Footer from "../components/layouts/footer/page";
import BlogCard from "../components/elements/blog-card/page";
import AddButton from "../components/elements/add-button/page";

const Blog = () => {
  return (
    <>
      <Header />
      <main>
        <AddButton />
        <BlogCard />
      </main>
      <Footer />
    </>
  );
};

export default Blog;
