import Header from "./components/layouts/Header";
import Footer from "./components/layouts/Footer";
import "@/scss/layout.scss";

export default function Home() {
  return (
    <>
      <Header />
      <main className="main">
        <div className="home dark:bg-gray-400">
          <h3 className="home__title dark:text-white">Welcome!</h3>
        </div>
      </main>
      <Footer />
    </>
  );
}
