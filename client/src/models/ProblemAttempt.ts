
export interface ProblemAttemptv1 {
    problemId: string,
    userInput: any,
    submissionState: ProblemAttemptSubmissionState,
}

// v1
//request
export interface ProblemAttemptRequestv1 extends ProblemAttemptv1 {
    userId?: string // TODO move this to a jwt in a query string
}


export interface ProblemAttemptResponsev1 {
    problemAttemptId: string,
    version: number,
    problemId: string,
    userInput: any,
    userId?: string
    submissionState: ProblemAttemptSubmissionState,
}

export enum ProblemAttemptSubmissionState {
    PENDING_SUBMISSION = "PENDING_SUBMISSION",
    SUBMITTED_PENDING_GRADE = "SUBMITTED_PENDING_GRADE",
    SUBMITTED_GRADED = "SUBMITTED_GRADED"
}

// export interface ProblemAttemptPresentationObject extends
//     ProblemAttemptResponsev1,
//     ProblemSolution_Response_v1,
//     GradeResponsev1 { }

// export interface ProblemSolution_Response_v1 {
//     problemSolution: Object, // comes from problem-service or saved somewhere?

// }

export interface HumanGradeRequestv1 {
    problemAttemptId: string,
    submittedByUserId: string,
    mark: string,
    reasoning: string,
}

export interface GradeResponsev1 {
    gradeId: string,
    problemAttemptId: string,
    version: number, // default to 1
    gradeType: GradeType, // Human Graded, Auto Graded,
    outcome?: HumanGradedOutcomev1 | ComputerGradedOutcomev1, // some object that depending on the 
    submissionState: GradeSubmissionState
}

export enum GradeType {
    HUMAN_GRADED = "HUMAN_GRADED",
    COMPUTER_GRADED = "COMPUTER_GRADED"
}

export interface HumanGradedOutcomev1 {
    gradedByUserId: string,
    reasoning: string,
    mark: string,
}

export interface ComputerGradedOutcomev1 {
    reasoning: string,
    mark: string,
}

export enum GradeSubmissionState {
    PENDING_SUBMISSION = "PENDING_SUBMISSION",
    GRADED_SUBMITTED = "GRADED_SUBMITTED"
}

// export interface BulkProblemAttemptRequestv1 {
    

// }

// export interface BulkProblemAttemptResponsev1 {


// }




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
