import ProfileCard from "../components/elements/ProfileCard";
import Header from "../components/layouts/Header";
import { auth } from "../lib/auth/auth";
import prisma from "../lib/db/prisma";

export default async function Dashboard() {
  "use server";

  const session = await auth();
  if (!session?.user?.email) {
    throw new Error("ログインが必要です");
  }

  const user = await prisma.user.findUnique({
    where: { email: session.user.email },
    select: { name: true, email: true },
  });

  console.log(user);

  if (!user) {
    throw new Error("ユーザーが見つかりません");
  }

  return (
    <div>
      <Header />
      <ProfileCard name={user.name} email={user.email} />
    </div>
  );
}
