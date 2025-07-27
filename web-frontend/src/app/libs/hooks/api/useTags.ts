import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { TagsService } from "@/app/libs/api/generated";
import { Tag } from "@/app/libs/api/generated/models/Tag";

export const useGetAllTags = () => {
  return useQuery({
    queryKey: ["tags"],
    queryFn: TagsService.getTags,
  });
};

export const useCreateTag = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: TagsService.postTags,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["tags"] });
    },
  });
};

export const useUpdateTag = () => {
  const queryClient = useQueryClient();
  return useMutation<{
    message?: string;
    tag?: Tag;
  }, Error, { id: number; name: string }>({
    mutationFn: ({ id, name }) => TagsService.putTags(id, { name }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["tags"] });
    },
  });
};

export const useDeleteTag = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: TagsService.deleteTags,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["tags"] });
    },
  });
};

export const useAddTagsToPost = () => {
  const queryClient = useQueryClient();
  return useMutation<void, Error, { postID: number; requestBody: { tag_names: string[] } }>({
    mutationFn: ({ postID, requestBody }) => TagsService.postTagsPosts(postID, requestBody),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["posts"] });
      queryClient.invalidateQueries({ queryKey: ["tags"] });
    },
  });
};

export const useRemoveTagsFromPost = () => {
  const queryClient = useQueryClient();
  return useMutation<void, Error, { postID: number; requestBody: { tag_names: string[] } }>({
    mutationFn: ({ postID, requestBody }) => TagsService.deleteTagsPosts(postID, requestBody),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["posts"] });
      queryClient.invalidateQueries({ queryKey: ["tags"] });
    },
  });
};

export const useGetTagsByPostID = (postID: number) => {
  return useQuery({
    queryKey: ["postTags", postID],
    queryFn: () => TagsService.getTagsPosts(postID),
    enabled: !!postID,
  });
};
