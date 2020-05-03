
# esecuzione

in primo terminale


```
go build 
./jobqueues -n 2048
```

in secondo terminale

```
for i in {1..4096}; do curl localhost:8000/work -d name=$USER -d delay=$((1+$i % 9))s; done
```
