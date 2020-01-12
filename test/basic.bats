BIN=./httpref

@test "Fails with no parameters" {
    run ${BIN}
    echo $output
    [ $status -ne 0 ]
}

@test "Check that filtering from root works" {
    run ${BIN} 1xx
    [ $status -eq 0 ]
}

@test "Check that filtering from status works" {
    run ${BIN} status 1xx
    [ $status -eq 0 ]
}

@test "Check that filtering from headers works" {
    run ${BIN} headers Accept-Ranges
    [ $status -eq 0 ]
}

@test "Check that filtering from headers alias works" {
    run ${BIN} header Accept-Ranges
    [ $status -eq 0 ]
}

