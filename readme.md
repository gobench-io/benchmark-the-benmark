# run benchmark the benchmark 

> Run then input benchmark client `drill`,`gobench`,`hey`,`jmeter`,`k6` and waiting for the loadtest is completed.

```bash
./run.sh  
```

Tail logs docker container to get current state:

```bash
docker logs -f benchmark-client
```