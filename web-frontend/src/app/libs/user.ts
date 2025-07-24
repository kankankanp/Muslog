import { useQuery } from "@tanstack/react-query";
import { Configuration, UsersApi } from "./api/generated";

const config = new Configuration({
  basePath: "http://localhost:8080", // 必要に応じて環境変数にしてください
});

const usersApi = new UsersApi(config);

/**
 * 全ユーザー取得
 */
export const useUsers = () => {
  return useQuery({
    queryKey: ["users"],
    queryFn: () => usersApi.usersGet().then((res) => res.data),
  });
};

// /**
//  * ユーザーIDから詳細を取得
//  */
// export const useUserById = (id: number) => {
//   return useQuery({
//     queryKey: ["users", id],
//     queryFn: () => usersApi.usersIdGet(id).then((res) => res.data),
//     enabled: !!id,
//   });
// };

// /**
//  * 特定ユーザーの投稿取得
//  */
// export const useUserPosts = (id: number) => {
//   return useQuery({
//     queryKey: ["users", id, "posts"],
//     queryFn: () => usersApi.usersIdPostsGet(id).then((res) => res.data),
//     enabled: !!id,
//   });
// };
