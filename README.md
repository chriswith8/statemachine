```mermaid
graph LR
    id0[broken]
    id1[off]
    id2[on]

    id1 --> |hit_someone| id0
    id1 --> |turn_on| id2
    id2 --> |hit_someone| id0
    id2 --> |turn_off| id1

    style id1 fill:#00AA00
