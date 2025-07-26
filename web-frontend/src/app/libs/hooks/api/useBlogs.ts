import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { BlogsService, Post } from "@/app/libs/api/generated";

export const useGetAllBlogs = () => {
  const { data, isPending, error } = useQuery<Post[], Error>({
    queryKey: ["blogs"],
    queryFn: async () => {
      const response = await BlogsService.getBlogs();
      return response.posts || [];
    },
  });
  return { data, isPending, error };
};

export const useGetBlogById = (id: number) => {
  const { data, isPending, error } = useQuery<Post, Error>({
    queryKey: ["blog", id],
    queryFn: async () => {
      const response = await BlogsService.getBlogs1(id);
      return response.post!;
    },
    enabled: !!id,
  });
  return { data, isPending, error };
};

export const useCreateBlog = () => {
  const queryClient = useQueryClient();
  const { mutate, data, isPending, error } = useMutation<
    Post,
    Error,
    {
      title: string;
      description: string;
      userId: string;
      tracks?: {
        spotifyId: string;
        name: string;
        artistName: string;
        albumImageUrl: string;
      }[];
    }
  >({
    mutationFn: async (newBlog) => {
      const response = await BlogsService.postBlogs(newBlog);
      return response.post!;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["blogs"] });
    },
  });
  return { mutate, data, isPending, error };
};

export const useUpdateBlog = () => {
  const queryClient = useQueryClient();
  const { mutate, isPending, error } = useMutation<
    Post,
    Error,
    { id: number; title: string; description: string }
  >({
    mutationFn: async (updatedBlog) => {
      const response = await BlogsService.putBlogs(updatedBlog.id, {
        title: updatedBlog.title,
        description: updatedBlog.description,
      });
      return response.post!;
    },
    onSuccess: (variables) => {
      queryClient.invalidateQueries({ queryKey: ["blog", variables.id] });
      queryClient.invalidateQueries({ queryKey: ["blogs"] });
    },
  });
  return { mutate, isPending, error };
};

export const useDeleteBlog = () => {
  const queryClient = useQueryClient();
  const { mutate, isPending, error } = useMutation<void, Error, number>({
    mutationFn: async (id) => {
      await BlogsService.deleteBlogs(id);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["blogs"] });
    },
  });
  return { mutate, isPending, error };
};

export const useGetBlogsByPage = (page: number) => {
  const { data, isPending, error } = useQuery<
    { posts: Post[]; totalCount: number },
    Error
  >({
    queryKey: ["blogs", "page", page],
    queryFn: async () => {
      const response = await BlogsService.getBlogsPage(page);
      return {
        posts: response.posts || [],
        totalCount: response.totalCount || 0,
      };
    },
    enabled: !!page,
  });
  return { data, isPending, error };
};
