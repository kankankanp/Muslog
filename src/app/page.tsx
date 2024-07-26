import Header from "./components/layouts/header/page";
import Footer from "./components/layouts/footer/page";
import styles from "@/scss/layout.module.scss";

export default function Home() {
  return (
    <>
      <Header />
      <main className={styles.main}>
        <h1>Hello</h1>
      </main>
      <Footer />
    </>
  );
}
