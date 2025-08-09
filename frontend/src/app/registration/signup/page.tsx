import type { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'Muslog - 新規登録',
  description: 'Muslogに新規登録して、音楽ブログを始めましょう。',
};

import SignupForm from "@/components/elements/forms/SignupForm";

export default function Page() {
  return <SignupForm />;
}
