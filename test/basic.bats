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

@test "Finds unique entry on exact match on root" {
    run ${BIN} Accept
    [ $status -eq 0 ]
    [ "$(echo $output | grep -c '/docs/Web/HTTP/Headers/Accept$')" -eq 1 ]
}

@test "Wildcard matching works on root" {
    run ${BIN} Accept\*
    [ $status -eq 0 ]
    [ "$(echo $output | grep -c '/docs/Web/HTTP/Headers/Accept$')" -gt 1 ]
}

@test "Finds unique entry on exact match on headers" {
    run ${BIN} headers Accept
    [ $status -eq 0 ]
    [ "$(echo $output | grep -c '/docs/Web/HTTP/Headers/Accept$')" -eq 1 ]
}

