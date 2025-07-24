import { NextResponse } from "next/server";
import { auth } from "@/app/lib/auth/auth";
import prisma from "@/app/lib/db/prisma";

// ページごとのブログ記事取得API
export const GET = async (
  req: Request,
  { params }: { params: { page: string } }
) => {
  try {
    const session = await auth();

    if (!session?.user?.id) {
      return NextResponse.json({ message: "Unauthorized" }, { status: 401 });
    }

    const PER_PAGE = 4;
    const page = parseInt(params.page, 10);
    const skip = (page - 1) * PER_PAGE;

    const [posts, totalCount] = await Promise.all([
      prisma.post.findMany({
        where: { userId: session.user.id },
        skip,
        take: PER_PAGE,
        orderBy: { createdAt: "desc" },
      }),
      prisma.post.count({
        where: { userId: session.user.id },
      }),
    ]);

    return NextResponse.json(
      { message: "Success", posts, totalCount },
      { status: 200 }
    );
  } catch (error) {
    return NextResponse.json({ message: "Error", error }, { status: 500 });
  }
};
