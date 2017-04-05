# run_tests.sh lints, tests, and benchmarks the landgrab package.

# log the current step.
log() {
  echo $1:
  echo -----------------------------------------
}

#echo usage.
usage() {
  echo 'sh run_tests.sh ?--bench'
  echo
  echo Run linter and tasts on landgrab package. Passing the bench flag will
  echo also run benchmarks.
}

if [ $# -gt 1 ]; then
  usage
  exit 1
fi

if [ $# -eq 1 ]; then
  if [ $1 = '--bench' ]; then
    BENCH='--bench=.'
  else
    usage
    exit 1
  fi
fi

log Linting
cd app
dartanalyzer --options analysis_options.yaml .
cd ..
echo

log Testing
go test $BENCH ./... -v
echo
