import prisma from "@/app/lib/db/prisma";
import { NextResponse } from "next/server";

//ブログ全記事取得API
export const GET = async (req: Request) => {
  try {
    const posts = await prisma.post.findMany();
    return NextResponse.json({ message: "Success", posts }, { status: 200 });
  } catch (error) {
    return NextResponse.json({ message: "Error", error }, { status: 500 });
  }
};

//ブログ投稿用API
export const POST = async (req: Request) => {
  try {
    const { title, description } = await req.json();
    const post = await prisma.post.create({ data: { title, description } });
    return NextResponse.json({ message: "Success", post }, { status: 201 });
  } catch (error) {
    return NextResponse.json({ message: "Error", error }, { status: 500 });
  }
};
