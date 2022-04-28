基准测试覆盖率
```shell
go test -bench=. -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
```

基准测试的cpu
> 不能多个包一起测试

```shell
go test -bench=. -cpuprofile=cpu.pprof -benchtime=5s -count=3  ./str
```

生成 cpu 、 内存、block

- cpu 使用分析：`-cpuprofile=cpu.pprof`
- 内存使用分析：`-benchmem -memprofile=mem.pprof`
- block分析：`-blockprofile=block.pprof`
```shell
go test -bench=. -benchmem -memprofile=mem.pprof -cpuprofile=cpu.pprof -blockprofile=block.pprof -benchtime=5s -count=3  ./str
```

```shell
go tool pprof -http=:6060 cpu.pprof
go tool pprof -http=:6060 mem.pprof
go tool pprof -http=:6060 block.pprof
```
常用命令
```shell
go test -bench=. -benchmem -benchtime=5s -count=3 .
```
