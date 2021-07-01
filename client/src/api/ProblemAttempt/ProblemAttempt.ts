import Axios from "axios";
import {
  ProblemAttemptv1,
  ProblemAttemptResponsev1
} from "@/models/ProblemAttempt";

const axios = Axios.create({ baseURL: "https://localhost:8091" });

export const createProblemAttempt = (problemId: string) => {
  return axios.post<ProblemAttemptResponsev1>("/problem-attempt", {
    problemId
  });
};

export const getProblemAttemptById = (id: string) => {
  return axios.get<ProblemAttemptResponsev1>(`/problem-attempt/${id}`);
};

export const updateProblemAttempt = (problemAttempt: ProblemAttemptv1) => {
  return axios.put<ProblemAttemptResponsev1>(
    "/problem-attempt",
    problemAttempt
  );
};

// export const bulkSaveProblemAttempt = (problemAttempts: [problemId: string]) => {
//   return axios.post<BulkProblemAttemptResponse>("/bulk-problem-attempt", problemAttempts)
// }
