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
  answers: string;
}

export interface ProblemType {
  statementType: StatementType;
  answerType: AnswerType;
}

export enum StatementType {}

export enum AnswerType {}
