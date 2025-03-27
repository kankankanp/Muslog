import { Track } from "@prisma/client";
import { NextResponse } from "next/server";
import { auth } from "@/app/lib/auth/auth";
import prisma from "@/app/lib/db/prisma";

//ブログ全記事取得API
export const GET = async (req: Request) => {
  try {
    const session = await auth();

    if (!session?.user?.id) {
      return NextResponse.json({ message: "Unauthorized" }, { status: 401 });
    }

    const posts = await prisma.post.findMany({
      where: { userId: session.user.id },
      orderBy: { createdAt: "desc" },
      include: {
        tracks: true,
      },
    });

    return NextResponse.json({ message: "Success", posts }, { status: 200 });
  } catch (error) {
    return NextResponse.json({ message: "Error", error }, { status: 500 });
  }
};

//ブログ投稿用API
export const POST = async (req: Request) => {
  const session = await auth();

  if (!session?.user?.id) {
    return NextResponse.json({ message: "Unauthorized" }, { status: 401 });
  }

  try {
    const { title, description, tracks } = await req.json();

    const post = await prisma.post.create({
      data: {
        title,
        description,
        userId: session.user.id,
        tracks:
          tracks && tracks.length > 0
            ? {
                create: tracks.map((track: Track) => ({
                  spotifyId: track.spotifyId,
                  name: track.name,
                  artistName: track.artistName,
                  albumImageUrl: track.albumImageUrl,
                })),
              }
            : undefined,
      },
      include: {
        tracks: true,
      },
    });

    return NextResponse.json({ message: "Success", post }, { status: 201 });
  } catch (error) {
    return NextResponse.json({ message: "Error", error }, { status: 500 });
  }
};
