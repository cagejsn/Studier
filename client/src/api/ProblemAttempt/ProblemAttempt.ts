import Axios from "axios";
import { ProblemAttempt, ProblemAttemptResponse} from '@/models/UserInteraction'

const axios = Axios.create({ baseURL: "https://localhost:8091" });

export const saveProblemAttempt = () => {
  return axios.post<ProblemAttemptResponse>("/problem-attempt");
};

export const getProblemAttemptById = (id: string) => {
  return axios.get<ProblemAttemptResponse>(`/problem-attempt/${id}`);
};

export const updateProblemAttempt = (problemAttempt: ProblemAttempt) => {
  return axios.put<ProblemAttemptResponse>("/problem-attempt", problemAttempt);
};