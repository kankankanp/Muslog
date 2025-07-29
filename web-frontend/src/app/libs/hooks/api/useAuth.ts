import { useMutation } from "@tanstack/react-query";
import { useDispatch } from "react-redux";
import {
  AuthResponse,
  AuthService,
  LoginRequest,
  RegisterRequest,
  OpenAPI,
} from "@/app/libs/api/generated";
import { login } from "@/app/libs/store/authSlice";

export const useLogin = () => {
  const dispatch = useDispatch();
  const { mutate, isPending, error } = useMutation<
    AuthResponse,
    Error,
    LoginRequest
  >({
    mutationFn: async (credentials) => {
      const loginResponse = await AuthService.postAuthLogin(credentials);
      return loginResponse.user!;
    },
    onSuccess: (data) => {
      OpenAPI.TOKEN = data.accessToken;
      dispatch(login(data));
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
  const { mutate, isPending, error } = useMutation<AuthResponse, Error, void>({
    mutationFn: async () => {
      const response = await AuthService.postRefresh();
      return response;
    },
    onSuccess: (data) => {
      OpenAPI.TOKEN = data.accessToken;
    },
  });
  return { mutate, isPending, error };
};

export const useSignup = () => {
  const { mutate, isPending, error } = useMutation<
    AuthResponse,
    Error,
    RegisterRequest
  >({
    mutationFn: async (credentials) => {
      const response = await AuthService.postRegister(credentials);
      return response.user!;
    },
  });
  return { mutate, isPending, error };
};
