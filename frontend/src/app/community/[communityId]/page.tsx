import ChatClient from "./chat-client";

export async function generateStaticParams() {
  // TODO: 実際のバックエンドやデータソースからコミュニティIDを取得するロジックに置き換えてください。
  // これはNext.jsの静的エクスポート要件を満たすためのプレースホルダーです。
  const communityIds = ['community1', 'community2', 'community3', '123dd969-9ec7-45c5-a379-d1928e6a9943']; 

  return communityIds.map((id) => ({
    communityId: id,
  }));
}

// Next.js App Router のサーバーコンポーネントの標準的な定義
export default async function CommunityPage({ params }: { params: { communityId: string } }) {
  console.log("Server Component params:", params);
  return <ChatClient params={params} />;
}
