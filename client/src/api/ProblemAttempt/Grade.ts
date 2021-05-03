import Axios from "axios";
import { GradeResponsev1, HumanGradeRequestv1 } from '@/models/ProblemAttempt'

const axios = Axios.create({ baseURL: "https://localhost:8091" });

export const createGrade = (problemAttemptId: string) => {
    return axios.post<GradeResponsev1>("/grade", { problemAttemptId });
};

export const getGradeById = (gradeId: string) => {
    return axios.get<GradeResponsev1>(`/grade/${gradeId}`);
};

export const updateGrade = (humanGrade: HumanGradeRequestv1) => {
    return axios.put<GradeResponsev1>("/grade", humanGrade);
};

export const getGradeByProblemAttemptId = (problemAttemptId: string) => {
    return axios.get<GradeResponsev1>("/grade", { params: { problemAttemptId } })
}

