BIN=./httpref

@test "Fails with no parameters" {
    run ${BIN}
    echo $output
    [ $status -ne 0 ]
}

