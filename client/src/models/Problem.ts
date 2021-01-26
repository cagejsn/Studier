// rule:
// each version supports some number of statement_type and answer_type
// version 1 will support text statement_type and three answer_types:
// true/false textEntry multipleChoice

export interface Problem {
  id: string;
  state: string;
  version: number;
  statement: string;
  type: ProblemType;
  answers: TextEntryAnswer | MultipleChoiceAnswerCandidates | TrueFalseAnswer;
  solution: any;
}

export interface ProblemType {
  statementType: StatementType;
  answersType: AnswersType;
}

export enum StatementType {
  TEXT,
  IMAGE,
  AUDIO
}

export enum AnswersType {
  MULTIPLE_CHOICE = "MULTIPLE_CHOICE",
  TEXT_ENTRY = "TEXT_ENTRY",
  TRUE_FALSE = "TRUE_FALSE"
}


export interface MultipleChoiceAnswerCandidates {
  [letter: string]: string
}

type TrueFalseAnswer = boolean


type TextEntryAnswer = string
