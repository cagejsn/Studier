import Axios from "axios";
import { ProblemSet } from "@/models/ProblemSet";

const axios = Axios.create({ baseURL: "https://localhost:8090" });

export const createProblemSet = () => {
    return axios.post<ProblemSet>('/problem-set');
};

export const getProblemSetById = (id: string) => {
    return axios.get<ProblemSet>(`/problem-set/${id}`);
};

export const addProblemsToProblemSet = (problemIds: string[], problemSetId: string) => {
    return axios.patch<ProblemSet>(`/problem-set/${problemSetId}`, { includedProblems: problemIds });
};

// export const deleteProblemsFromProblemSet = (problemIds: string[], problemSetId: string) => {
//     return axios.delete<ProblemSet>(`/problem-set/${problemSetId}`, { params: { problemIds: problemIds } })
// };

// export const deleteProblemSet = (problemSetId: string) => {
//     return axios.delete<ProblemSet>(`/problem-set/${problemSetId}`)
// }

//TODO: break problemSet with Id and problemSet without Id into different objects
// (created in db vs. not created in db problemset)
export const updateProblemSet = (problemSet: ProblemSet) => {
    return axios.put<ProblemSet>(`/problem-set`, problemSet)
}