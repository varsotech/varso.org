import { GetOrganizationsResponse } from "@luminancetech/varso/src/app/apprequests";
import axios from "./axios";
import { AxiosResponse } from "axios";
import { useQuery, UseQueryOptions } from "./reactquery";

export function useOrganization(
  options?: UseQueryOptions<AxiosResponse<GetOrganizationsResponse>>
) {
  return useQuery<GetOrganizationsResponse>(
    async () => await axios.get("/api/v1/app/organization"),
    {
      queryKey: ["organization"],
      ...options,
    }
  );
}

// export function useCreateTask() {
//   return useMutation<CreateTaskRequest, AxiosResponse<CreateTaskResponse>>(
//     async (req) => await axios.post("/api/v1/sidekick/tasks", req),
//     { mutationKey: ["tasks"] }
//   )
// }

// export function useUpdateTask() {
//   return useMutation<UpdateTaskRequest, AxiosResponse<UpdateTaskResponse>>(
//     async (req) => await axios.post("/api/v1/sidekick/tasks/update", req),
//     { mutationKey: ["tasks"] }
//   )
// }

// export function useDeleteTask() {
//   return useMutation<DeleteTaskRequest, AxiosResponse<DeleteTaskResponse>>(
//     async (req) => await axios.post("/api/v1/sidekick/tasks/delete", req),
//     {
//       mutationKey: ["tasks"],
//     }
//   )
// }
