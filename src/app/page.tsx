import Header from "./components/layouts/header/page";
import Footer from "./components/layouts/footer/page";
import "@/scss/layout.scss";

export default function Home() {
  return (
    <>
      <Header />
      <main className="main">
        <div className="home">
          <h3 className="home__title">Welcome!</h3>
          <p className="home__text">
            This blog was created with Next.js, Typescrpt, and SCSS! test test2
          </p>
        </div>
      </main>
      <Footer />
    </>
  );
}
