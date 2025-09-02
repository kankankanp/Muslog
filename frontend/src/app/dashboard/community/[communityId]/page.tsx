import ChatClient from "./chat-client";

export async function generateStaticParams() {
  // TODO: 実際のバックエンドやデータソースからコミュニティIDを取得するロジックに置き換えてください。
  // これはNext.jsの静的エクスポート要件を満たすためのプレースホルダーです。
  const communityIds = ['community1', 'community2', 'community3', '123dd969-9ec7-45c5-a379-d1928e6a9943', 'c6da7786-a3d2-4f8e-835f-99b96eb9e0ec']; 

  return communityIds.map((id) => ({
    communityId: id,
  }));
}

export default async function CommunityPage({ params }: { params: { communityId: string } }) {
  return <ChatClient params={params} />;
}