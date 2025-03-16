import { auth } from "@/app/lib/auth/auth";
import prisma from "@/app/lib/db/prisma";
import { NextResponse } from "next/server";

//ブログ詳細記事取得API
export const GET = async (req: Request) => {
  try {
    const id = parseInt(req.url.split("/blog/")[1]);

    const post = await prisma.post.findUnique({
      where: { id },
      include: {
        tracks: true,
      },
    });

    if (!post) {
      return NextResponse.json({ message: "Not Found" }, { status: 404 });
    }

    return NextResponse.json({ message: "Success", post }, { status: 200 });
  } catch (error) {
    console.error("Error fetching post:", error);
    return NextResponse.json(
      { message: "Error", error: String(error) },
      { status: 500 }
    );
  }
};

//ブログの記事編集API
export const PUT = async (req: Request) => {
  try {
    const id: number = parseInt(req.url.split("/blog/")[1]);

    const { title, description } = await req.json();

    const post = await prisma.post.update({
      data: { title, description },
      where: { id },
      include: {
        tracks: true,
      },
    });
    return NextResponse.json({ message: "Success", post }, { status: 200 });
  } catch (error) {
    return NextResponse.json({ message: "Error", error }, { status: 500 });
  }
};

//ブログの記事削除API
export const DELETE = async (req: Request) => {
  try {
    const id = parseInt(req.url.split("/blog/")[1]);

    if (isNaN(id)) {
      console.error("Invalid ID:", req.url);
      return NextResponse.json({ message: "Invalid ID" }, { status: 400 });
    }

    const post = await prisma.post.delete({
      where: { id },
      include: { tracks: true },
    });

    return NextResponse.json({ message: "Success", post }, { status: 200 });
  } catch (error: any) {
    console.error("Error deleting post:", error);

    if (error.code === "P2025") {
      return NextResponse.json({ message: "Post not found" }, { status: 404 });
    }

    return NextResponse.json(
      { message: "Error", error: error.message },
      { status: 500 }
    );
  }
};
