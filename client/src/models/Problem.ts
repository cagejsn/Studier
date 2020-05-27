

export interface Problem {
    statement: string;
    type: ProblemType;
    answers: string;
}


export enum ProblemType {
    TEXT_ENTRY,
    MULTIPLE_CHOICE,
    TRUE_OR_FALSE
}