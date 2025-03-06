// app/api/spotify/search/route.ts
import { getAccessToken } from "@/app/lib/utils/spotify";
import { NextRequest, NextResponse } from "next/server";

export async function GET(req: NextRequest) {
  const query = req.nextUrl.searchParams.get("q");

  if (!query) {
    return NextResponse.json(
      { error: "検索ワードがありません" },
      { status: 400 }
    );
  }

  try {
    const token = await getAccessToken();
    const res = await fetch(
      `https://api.spotify.com/v1/search?q=${encodeURIComponent(
        query
      )}&type=track&limit=5`,
      {
        headers: { Authorization: `Bearer ${token}` },
      }
    );

    console.log(res);

    if (!res.ok) {
      const errorText = await res.text();
      console.error("Spotify検索エラー:", errorText);
      return NextResponse.json({ error: "Spotify検索に失敗" }, { status: 500 });
    }

    const data = await res.json();
    return NextResponse.json(data.tracks.items);
  } catch (err) {
    console.error(err);
    return NextResponse.json({ error: "サーバーエラー" }, { status: 500 });
  }
}
