import { AxiosResponse } from "axios";
import { UseQueryOptions, useMutation, useQuery } from "./reactquery";
import {
  DeleteNewsItemRequest,
  BlurToggleRequest,
  GetNewsResponse,
} from "../proto/src/app/apprequests";
import axios from "./axios";

export function useNews(
  feedType: string,
  options?: UseQueryOptions<AxiosResponse<GetNewsResponse>>
) {
  return useQuery<GetNewsResponse>(
    async () => await axios.get(`/api/v1/app/news/${feedType}`),
    {
      queryKey: ["news"],
      ...options,
    }
  );
}

export function useToggleBlur(newsItemId: string) {
  return useMutation<unknown, AxiosResponse<any>>(
    async () =>
      await axios.post(
        "/api/v1/app/news/item/blur/toggle",
        BlurToggleRequest.create({
          rssItemId: newsItemId,
        })
      ),
    { mutationKey: ["news"] }
  );
}

export function useDelete(newsItemId: string) {
  return useMutation<unknown, AxiosResponse<any>>(
    async () =>
      await axios.post(
        "/api/v1/app/news/item/delete",
        DeleteNewsItemRequest.create({
          rssItemId: newsItemId,
        })
      ),
    { mutationKey: ["news"] }
  );
}
