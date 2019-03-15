# suco

sun colony

    suco-d     - desktop client application
    suco-front - front-door api service (run as N-replica pod)
    suco-db    - fake persistent storage service for suco-front
    suco-zone  - single-zone service (created on-demand by suco-front as 1-replica pod for every manned zone)
