create table problem_attempt
(
    id CHAR(36) primary key,
    version int default 1 not null,
    problem_id CHAR(36) not null,
    person CHAR(36) REFERENCES person(id),
    user_input text DEFAULT '""' not null,
    submission_state text default 'PENDING_SUBMISSION'
);
