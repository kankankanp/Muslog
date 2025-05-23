import "@/scss/global.css";
import "../../node_modules/destyle.css";
import { config } from "@fortawesome/fontawesome-svg-core";
import "@fortawesome/fontawesome-svg-core/styles.css";
import type { Metadata } from "next";
import { Noto_Sans_JP } from "next/font/google";
import Footer from "./components/layouts/Footer";
import Header from "./components/layouts/Header";
import { NextAuthProvider } from "./lib/auth/NextAuthProvider";
import { Providers } from "./lib/store/Providers";

config.autoAddCss = false;

const notoJp = Noto_Sans_JP({
  subsets: ["latin"],
  weight: ["100", "300", "400", "500", "700", "900"], // 必要なウェイトを指定
  variable: "--font-notojp",
  display: "swap",
});

export const metadata: Metadata = {
  title: "Muslog",
  description: "Music for your everyday blog.",
  icons: {
    icon: "/favicon.ico",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body className={notoJp.variable}>
        <NextAuthProvider>
          <Providers>
            <Header />
            {children}
            <Footer />
          </Providers>
        </NextAuthProvider>
      </body>
    </html>
  );
}
