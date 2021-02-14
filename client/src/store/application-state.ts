import { Problem } from '@/models/Problem';
import { ProblemAttemptv1 } from '@/models/ProblemAttempt';

interface ApplicationState {
    openProblem?: Problem

    currentAttempt?: ProblemAttemptv1
}



export const initialState: ApplicationState = {
    openProblem: undefined,
    currentAttempt: undefined,
}