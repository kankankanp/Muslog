const SPOTIFY_CLIENT_ID = process.env.SPOTIFY_CLIENT_ID!;
const SPOTIFY_CLIENT_SECRET = process.env.SPOTIFY_CLIENT_SECRET!;

let cachedToken: string | null = null;
let tokenExpiresAt = 0;

export const getAccessToken = async () => {
  const currentTime = Date.now();

  if (cachedToken && currentTime < tokenExpiresAt) {
    return cachedToken;
  }

  const auth = Buffer.from(
    `${SPOTIFY_CLIENT_ID}:${SPOTIFY_CLIENT_SECRET}`
  ).toString("base64");

  const res = await fetch("https://accounts.spotify.com/api/token", {
    method: "POST",
    headers: {
      Authorization: `Basic ${auth}`,
      "Content-Type": "application/x-www-form-urlencoded",
    },
    body: "grant_type=client_credentials",
  });

  if (!res.ok) {
    const errorText = await res.text();
    console.error("Spotify認証エラー:", errorText);
    throw new Error(`Spotify認証に失敗しました (${res.status})`);
  }

  const data = await res.json();
  cachedToken = data.access_token;
  tokenExpiresAt = currentTime + data.expires_in * 1000 - 60000;

  return cachedToken;
};
