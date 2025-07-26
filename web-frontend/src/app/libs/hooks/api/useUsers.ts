import { useQuery } from "@tanstack/react-query";
import { UsersService, User, Post } from "@/app/libs/api/generated";

export const useGetAllUsers = () => {
  const { data, isPending, error } = useQuery<User[], Error>({
    queryKey: ["users"],
    queryFn: async () => {
      const response = await UsersService.getUsers();
      return response.users || [];
    },
  });
  return { data, isPending, error };
};

export const useGetUserById = (id: string) => {
  const { data, isPending, error } = useQuery<User, Error>({
    queryKey: ["user", id],
    queryFn: async () => {
      const response = await UsersService.getUsers1(id);
      return response.user!;
    },
    enabled: !!id,
  });
  return { data, isPending, error };
};

export const useGetUserPosts = (id: string) => {
  const { data, isPending, error } = useQuery<Post[], Error>({
    queryKey: ["userPosts", id],
    queryFn: async () => {
      const response = await UsersService.getUsersPosts(id);
      return response.posts || [];
    },
    enabled: !!id,
  });
  return { data, isPending, error };
};