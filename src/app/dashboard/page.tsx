import { redirect } from "next/navigation";
import ProfileCard from "../components/elements/cards/ProfileCard";
import { Book } from "../components/elements/others/Book";
import { auth } from "../lib/auth/auth";
import prisma from "../lib/db/prisma";
import { fetchAllBlogs } from "../lib/utils/blog";

export default async function Page() {
  const session = await auth();
  if (!session?.user?.email) {
    redirect("/registration/login");
  }

  const user = await prisma.user.findUnique({
    where: { email: session.user.email },
    select: { name: true, email: true },
  });

  if (!user) {
    throw new Error("ユーザーが見つかりません");
  }

  const posts = await fetchAllBlogs();

  return (
    <div className="dark:bg-gray-900 bg-gray-100">
      <div className="py-8">
        <ProfileCard name={user.name} email={user.email} />
        <Book posts={posts} />
      </div>
    </div>
  );
}
