import { Problem } from '@/models/Problem';


interface ApplicationState {
    openProblem?: Problem
}



export const initialState: ApplicationState = {
    openProblem: undefined
}