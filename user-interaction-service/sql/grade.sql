create table grade
(
    id CHAR(36) primary key,
    version int default 1 not null,
    problem_attempt_id CHAR(36) REFERENCES problem_attempt not null,
    person CHAR(36) REFERENCES person(id),
    user_input text DEFAULT '""' not null,
    grade_submission_state text default 'PENDING_SUBMISSION'
);
