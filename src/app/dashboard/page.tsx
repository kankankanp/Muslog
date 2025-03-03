import ProfileCard from "../components/elements/ProfileCard";
import Footer from "../components/layouts/Footer";
import Header from "../components/layouts/Header";
import { auth } from "../lib/auth/auth";
import prisma from "../lib/db/prisma";
import { redirect } from "next/navigation";


export default async function Dashboard() {
  "use server";

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

  return (
    <div className="dark:bg-gray-900 bg-gray-100">
      <Header />
      <div className="py-8">
      <ProfileCard name={user.name} email={user.email} />
      </div>
      <Footer /> 
    </div>
  );
}
