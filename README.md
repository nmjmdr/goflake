# goflake
Flake implementation in go based on Factual\Skuld'd flake clojure implementation

Current limitations:

The current implementation does not store last timestamp in a persistent store. This means that:

1. In case the process hosting flake is stopped
2. The system time of the node hosting the flake process is moved back

Then the current flake implementation will generate id that has a timestamp that has been used as part of an earlier id.
