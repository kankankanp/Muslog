import { useQuery } from "@tanstack/react-query";
import { blogsApi } from "./api/apiClient";

// export const useUserBlogs = (userId?: string | number) => {
//   return useQuery({
//     queryKey: ["user-blogs", userId],
//     queryFn: () =>
//       blogsApi
//         .usersIdPostsGet(String(userId))
//         .then((res: { data: any }) => res.data),
//     enabled: !!userId,
//   });
// };

export const useAllBlogs = () => {
  return useQuery({
    queryKey: ["blogs"],
    queryFn: () => blogsApi.blogsGet().then((res: { data: any }) => res.data),
  });
};
