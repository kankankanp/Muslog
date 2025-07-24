import { NextResponse } from "next/server";
import prisma from "@/app/lib/db/prisma";

//ブログ詳細記事取得API
export const GET = async (
  req: Request,
  { params }: { params: { id: string } }
) => {
  try {
    const id = parseInt(params.id);

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
    return NextResponse.json({ message: "Error", error }, { status: 500 });
  }
};

//ブログの記事編集API
export const PUT = async (
  req: Request,
  { params }: { params: { id: string } }
) => {
  try {
    const id: number = parseInt(params.id);

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
export const DELETE = async (
  req: Request,
  { params }: { params: { id: string } }
) => {
  try {
    const id = parseInt(params.id);
    console.log("Attempting to delete post with ID:", id);

    // まず投稿が存在するか確認
    const existingPost = await prisma.post.findUnique({
      where: { id },
      include: { tracks: true },
    });

    if (!existingPost) {
      return NextResponse.json({ message: "Post not found" }, { status: 404 });
    }

    if (existingPost.tracks.length > 0) {
      await prisma.track.deleteMany({
        where: { postId: id },
      });
    }

    const post = await prisma.post.delete({
      where: { id },
    });

    return NextResponse.json({ message: "Success", post }, { status: 200 });
  } catch (error) {
    return NextResponse.json({ message: "Error", error }, { status: 500 });
  }
};
