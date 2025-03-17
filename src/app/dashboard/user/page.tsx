import ProfileCard from "@/app/components/elements/ProfileCard";
import { auth } from "@/app/lib/auth/auth";
import prisma from "@/app/lib/db/prisma";

export default async function Page() {
  "use server";

  const session = await auth();
  if (!session?.user?.email) {
    throw new Error("ログインが必要です");
  }

  const user = await prisma.user.findUnique({
    where: { email: session.user.email },
    select: { name: true, email: true },
  });

  if (!user) {
    throw new Error("ユーザーが見つかりません");
  }
  return (
    <div>
      <ProfileCard name={user.name} email={user.email} />
    </div>
  );
}
