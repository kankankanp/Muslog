import type { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'Muslog - ログイン',
  description: 'Muslogにログインして、音楽ブログを始めましょう。',
};

import LoginForm from "@/components/elements/forms/LoginForm";

export default function Page() {
  return <LoginForm />;
}
