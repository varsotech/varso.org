import React, { useEffect, useMemo, useState } from "react";
import { PersistQueryClientProvider as RQPersistQueryClientProvider } from "@tanstack/react-query-persist-client";
import { createSyncStoragePersister } from "@tanstack/query-sync-storage-persister";
import {
  MutationFunction,
  QueryClient,
  QueryClientProvider,
  QueryFunction,
  QueryKey,
  useMutation as rqUseMutation,
  useQuery as rqUseQuery,
  UseQueryOptions as RQUseQueryOptions,
  UseMutationOptions,
  useQueryClient,
} from "@tanstack/react-query";
import { AxiosResponse } from "axios";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      gcTime: Infinity,
      staleTime: Infinity, 
    },
  },
});

type PersistQueryClientProviderProps = {
  children: React.ReactNode;
};

export function PersistQueryClientProvider({
  children,
}: PersistQueryClientProviderProps): JSX.Element {
  const [isClient, setIsClient] = useState(false);

  useEffect(() => {
    setIsClient(true);
  }, []);

  if (isClient) {
    const persistOptions = {
      persister: createSyncStoragePersister({
        storage: localStorage,
      }),
    };

    return (
      <RQPersistQueryClientProvider
        client={queryClient}
        persistOptions={persistOptions}
      >
        {children}
      </RQPersistQueryClientProvider>
    );
  }

  return <div></div>;
}

// Use wrapper function

export type UseQueryOptions<
  TResponse,
  TQueryKey extends QueryKey = QueryKey
> = Omit<
  RQUseQueryOptions<TResponse, unknown, TResponse, TQueryKey>,
  "initialData"
> & { initialData?: () => undefined };

// useQuery is a thin react-query wrapper function
export function useQuery<TResponse, TQueryKey extends QueryKey = QueryKey>(
  queryFn: QueryFunction<AxiosResponse<TResponse>, TQueryKey>,
  options: UseQueryOptions<AxiosResponse<TResponse>, TQueryKey>
) {
  return rqUseQuery<
    AxiosResponse<TResponse>,
    unknown,
    AxiosResponse<TResponse>,
    TQueryKey
  >({
    queryFn: async (ctx) => {
      const data = await queryFn(ctx);
      if (data == null) {
        throw new Error("failed fetching data");
      }

      if (data.status !== 200) {
        throw new Error("failed fetching data, got status " + data.status);
      }

      return data;
    },
    ...options,
  });
}

// useMutation is a thin react-query wrapper function
export function useMutation<TVariables, TData>(
  mutationFn: MutationFunction<TData, TVariables>,
  options?: Omit<
    UseMutationOptions<TData, unknown, TVariables, unknown>,
    "mutationFn"
  >
) {
  const queryClient = useQueryClient();

  return rqUseMutation({
    ...options,
    mutationFn: mutationFn,
    onSettled: async (data, error, variables, context) => {
      await queryClient.invalidateQueries({
        queryKey: options?.mutationKey,
      });
      options?.onSettled &&
        (await options?.onSettled(data, error, variables, context));
    },
  });
}
