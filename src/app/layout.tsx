import type { Metadata } from "next";
import { Lato } from "next/font/google";
import { Noto_Sans_JP } from "next/font/google";
import "@/scss/global.scss"
import "../../node_modules/destyle.css";
import { config } from "@fortawesome/fontawesome-svg-core";
import "@fortawesome/fontawesome-svg-core/styles.css";
config.autoAddCss = false;

const lato = Lato({
  subsets: ["latin"],
  weight: "100",
  variable: "--font-lato",
  display: "swap",
});

const notoJp = Noto_Sans_JP({
  subsets: ["latin"],
  weight: "100",
  variable: "--font-notojp",
  display: "swap",
});

const metadata: Metadata = {
  title: "Simple-blog",
  description: "My first blog using Next.js,TypeScript,Sass",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body className={`${lato.variable} ${notoJp.variable}`}>{children}</body>
    </html>
  );
}
