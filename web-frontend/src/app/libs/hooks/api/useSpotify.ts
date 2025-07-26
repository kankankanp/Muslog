import { useQuery } from "@tanstack/react-query";
import { SpotifyService, Track } from "@/app/libs/api/generated";

export const useSearchSpotifyTracks = (query: string) => {
  const { data, isPending, error } = useQuery<Track[], Error>({
    queryKey: ["spotifyTracks", query],
    queryFn: async () => {
      const response = await SpotifyService.getSpotifySearch(query);
      return response.tracks || [];
    },
    enabled: !!query,
  });
  return { data, isPending, error };
};