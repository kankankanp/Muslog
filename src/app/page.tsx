import Header from "./components/layouts/Header";
import Footer from "./components/layouts/Footer";
import "@/scss/layout.scss";

export default function Home() {
  return (
    <>
      <Header />
      <main className="main">
        <div className="home">
          <h3 className="home__title">Welcome!</h3>
          <p className="home__text">
            This blog was created with Next.js, Typescrpt, and scss! test test2
            test 3
          </p>
        </div>
      </main>
      <Footer />
    </>
  );
}
