import Axios from 'axios'
import {Problem, ProblemType} from '@/models/Problem'


const axios = Axios.create({baseURL: 'https://localhost:8090'});

export const createProblem = () => {
    return axios.post<Problem>(
        '/problem'
    )
}