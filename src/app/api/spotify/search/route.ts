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
      { message: "Missing search term" },
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
      return NextResponse.json(
        { message: "Failed to search tracks" },
        { status: 500 }
      );
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

    return NextResponse.json(
      { message: "Success", tracks: formattedTracks },
      { status: 200 }
    );
  } catch (error) {
    return NextResponse.json({ message: "Error", error }, { status: 500 });
  }
}
