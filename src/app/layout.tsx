import type { Metadata } from "next";
import { Noto_Sans_JP } from "next/font/google";
import "@/scss/global.scss";
import "../../node_modules/destyle.css";
import { config } from "@fortawesome/fontawesome-svg-core";
import "@fortawesome/fontawesome-svg-core/styles.css";
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
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body className={notoJp.variable}>
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
