import "@/scss/global.css";
import { config } from "@fortawesome/fontawesome-svg-core";
import "@fortawesome/fontawesome-svg-core/styles.css";
import "destyle.css";
import type { Metadata } from "next";
import { Noto_Sans_JP } from "next/font/google";
import { Providers } from "./providers";
config.autoAddCss = false;

const notoJp = Noto_Sans_JP({
  subsets: ["latin"],
  weight: ["100", "300", "400", "500", "700", "900"],
  variable: "--font-notojp",
  display: "swap",
});

export const metadata: Metadata = {
  title: "Muslog - あなたの音楽の物語をシェアしよう",
  description:
    "好きな曲と共に、あなたの想いを綴る。新しい音楽との出会いの場へ。",
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
        {/* <NextAuthProvider>
          <Providers>
            <Header /> */}
        <Providers>{children}</Providers>
        {/* <Footer />
          </Providers>
        </NextAuthProvider> */}
      </body>
    </html>
  );
}
