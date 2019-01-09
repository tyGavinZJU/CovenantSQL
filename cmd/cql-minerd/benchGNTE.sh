#!/bin/bash

#make -C ../../ clean && \
#make -C ../../ use_all_cores && \
go test -bench=^BenchmarkMinerGNTE1$ -benchtime=10s -run ^$ |tee gnte.log
go test -bench=^BenchmarkMinerGNTE2$ -benchtime=10s -run ^$ |tee -a gnte.log

go test -cpu=1 -bench=^BenchmarkMinerGNTE1$ -benchtime=10s -run ^$ |tee -a gnte.log
go test -cpu=1 -bench=^BenchmarkMinerGNTE2$ -benchtime=10s -run ^$ |tee -a gnte.log
