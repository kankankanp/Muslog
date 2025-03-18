import { NextRequest, NextResponse } from "next/server";
import { getAccessToken } from "@/app/lib/utils/spotify";

type SpotifyTrackResponse = {
  id: string;
  name: string;
  artists: { name: string }[];
  album: {
    images: { url: string }[];
  };
};

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
    const LIMIT = 10;

    const res = await fetch(
      `https://api.spotify.com/v1/search?q=${encodeURIComponent(
        query
      )}&type=track&limit=${LIMIT}`,
      {
        headers: { Authorization: `Bearer ${token}` },
      }
    );

    if (!res.ok) {
      const errorText = await res.text();
      console.error("Spotify検索エラー:", errorText);
      return NextResponse.json({ error: "Spotify検索に失敗" }, { status: 500 });
    }

    const data = await res.json();

    const formattedTracks = data.tracks.items.map(
      (track: SpotifyTrackResponse) => ({
        spotifyId: track.id,
        name: track.name,
        artistName: track.artists
          .map((artist: { name: string }) => artist.name)
          .join(", "),
        albumImageUrl:
          track.album.images.length > 0
            ? track.album.images[0].url
            : "/default-image.jpg",
      })
    );

    return NextResponse.json(formattedTracks);
  } catch (err) {
    console.error("サーバーエラー:", err);
    return NextResponse.json({ error: "サーバーエラー" }, { status: 500 });
  }
}
