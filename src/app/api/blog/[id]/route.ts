import prisma from "@/app/lib/prisma";
import { NextResponse } from "next/server";

//ブログ詳細記事取得API
export const GET = async (req: Request) => {
  try {
    const id: number = parseInt(req.url.split("/blog/")[1]);
    const post = await prisma.post.findFirst({ where: { id } });
    return NextResponse.json({ message: "Success", post }, { status: 200 });
  } catch (error) {
    return NextResponse.json({ message: "Error", error }, { status: 500 });
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
    });
    return NextResponse.json({ message: "Success", post }, { status: 200 });
  } catch (error) {
    return NextResponse.json({ message: "Error", error }, { status: 500 });
  }
};

//ブログの記事削除API
export const DELETE = async (req: Request) => {
  try {
    const id: number = parseInt(req.url.split("/blog/")[1]);

    const post = await prisma.post.delete({
      where: { id },
    });
    return NextResponse.json({ message: "Success", post }, { status: 200 });
  } catch (error) {
    return NextResponse.json({ message: "Error", error }, { status: 500 });
  }
};
