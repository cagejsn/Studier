// v1
//request
// TODO: include the user's id, that should come from some kind of token
export interface ProblemAttempt {
    problemId: string,
    userInput: Object,
    submissionState: string, // PENDING_SUBMISSION (DRAFT), SUBMITTED_PENDING_GRADING, SUBMITTED_GRADED,
}

//response
export interface ProblemAttemptResponse {
    problemAttemptId: string,
    problemId: string,
    userInput: Object,
    problemSolution: Object,
    submissionState: string, // PENDING_SUBMISSION (DRAFT), SUBMITTED_PENDING_GRADING, SUBMITTED_GRADED,
    submissionOutcome: string, // TODO: boolean? enum? if human graded, then this column will be more important
}

// export interface ProblemSet {
//     id: string, // id for the problem set
//     problemIds: string[] // ids of the constituent problems
// }

// export interface ProblemSetAttempt {
//     problemSetId: string, // id for problemSet
//     problemAttempts: ProblemAttempt[],
// }

// export interface ProblemSetAttemptResponse {
//     problemSetId: string, // id for problemSet
//     problemAttempts: ProblemAttempt[],
// }

//v2
