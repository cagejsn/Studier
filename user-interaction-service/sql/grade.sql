create table grade
(
    id CHAR(36) primary key,
    version int default 1 not null,
    problem_attempt_id CHAR(36) REFERENCES problem_attempt not null,
    grade_type text DEFAULT 'COMPUTER_GRADED' not null,
    -- person CHAR(36) REFERENCES person(id),
    outcome text default '""' not null,
    submission_state text default 'PENDING_SUBMISSION'
);
