import Header from "./components/layouts/header/page";
import Footer from "./components/layouts/footer/page";
import "@/scss/layout.scss";

export default function Home() {
  return (
    <>
      <Header />
      <main className="main">
        <h1>Hello</h1>
      </main>
      <Footer />
    </>
  );
}
