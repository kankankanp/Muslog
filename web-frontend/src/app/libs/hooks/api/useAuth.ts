import { useMutation } from "@tanstack/react-query";
import { AuthService } from "@/app/libs/api/generated";
import { User } from "@/app/libs/api/generated/models/User";
import { UserLogin } from "@/app/libs/api/generated/models/UserLogin";

export const useLogin = () => {
  const { mutate, isPending, error } = useMutation<{
    id: string;
    name: string;
    accessToken: string;
  }, Error, UserLogin>({
    mutationFn: async (credentials) => {
      const loginResponse = await AuthService.postLogin(credentials);
      const refreshResponse = await AuthService.postRefresh();
      return {
        id: loginResponse.user!.id,
        name: loginResponse.user!.name,
        accessToken: refreshResponse.accessToken!,
      };
    },
  });
  return { mutate, isPending, error };
};

export const useLogout = () => {
  const { mutate, isPending, error } = useMutation<void, Error>({
    mutationFn: async () => {
      await AuthService.postLogout();
    },
  });
  return { mutate, isPending, error };
};

export const useRefreshToken = () => {
  const { mutate, isPending, error } = useMutation<void, Error>({
    mutationFn: async () => {
      await AuthService.postRefresh();
    },
  });
  return { mutate, isPending, error };
};

export const useSignup = () => {
  const { mutate, data, isPending, error } = useMutation<User, Error, { name: string; email: string; password: string }>({
    mutationFn: async (credentials) => {
      const response = await AuthService.postRegister(credentials);
      return response.user!;
    },
  });
  return { mutate, data, isPending, error };
};
