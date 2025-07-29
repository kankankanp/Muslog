import "@/scss/global.css";
import { config } from "@fortawesome/fontawesome-svg-core";
import "@fortawesome/fontawesome-svg-core/styles.css";
import "destyle.css";
import type { Metadata } from "next";
import { Noto_Sans_JP } from "next/font/google";
import { Providers } from "./providers";
import { NextAuthProvider } from "./providers/NextAuthProvider";
import Header from "./components/layouts/Header";
import Footer from "./components/layouts/Footer";
config.autoAddCss = false;

const notoJp = Noto_Sans_JP({
  subsets: ["latin"],
  weight: ["100", "300", "400", "500", "700", "900"],
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
    <html lang="ja" suppressHydrationWarning>
      <body className={`${notoJp.variable} antialiased`}>
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
