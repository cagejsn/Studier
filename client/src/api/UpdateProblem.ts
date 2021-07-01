import Axios from "axios";
import { Problem, ProblemType } from "@/models/Problem";

const axios = Axios.create({ baseURL: "https://localhost:8090" });

export const updateProblem = (problem: Problem) => {
  return axios.put<Problem>("/problem", problem);
};
