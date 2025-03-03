import Header from "./components/layouts/Header";
import Footer from "./components/layouts/Footer";

export default function Home() {
  return (
    <>
      <Header />
      <main className="dark:bg-gray-900 bg-gray-100 min-h-screen flex items-center justify-center">
        <div className="bg-white p-10 rounded-lg shadow-lg text-center dark:bg-gray-800">
          <h3 className="text-2xl font-bold dark:text-white text-gray-900">
            Welcome to My App!
          </h3>
        </div>
      </main>
      <Footer />
    </>
  );
}
