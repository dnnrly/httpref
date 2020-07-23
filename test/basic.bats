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

@test "Check that filtering from alias status works" {
    run ${BIN} statuses 1xx
    [ $status -eq 0 ]
}

@test "Check that filtering from status works" {
    run ${BIN} status 1xx
    [ $status -eq 0 ]
}

@test "Check that filtering from methods works" {
    run ${BIN} methods GET
    [ $status -eq 0 ]
}

@test "Check that filtering from methods alias works" {
    run ${BIN} method GET
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
    run ${BIN} 'Accept*'
    [ $status -eq 0 ]
    [ "$(echo $output | grep -c '/docs/Web/HTTP/Headers/Accept$')" -eq 0 ]
}

@test "Finds unique entry on exact match on headers" {
    run ${BIN} headers Accept
    [ $status -eq 0 ]
    [ "$(echo $output | grep -c '/docs/Web/HTTP/Headers/Accept$')" -eq 1 ]
}

@test "Can change the width of the output" {
    run ${BIN} -w 70 100
    [ $status -eq 0 ]
    for l in "${lines[@]}"
    do
        [ $(echo ${l} | wc -m) -le 70 ]
    done
}

@test "You can get the information for a single port" {
    run ${BIN} 80
    [ $status -eq 0 ]
    [ "$(echo $output | grep -c 'Hypertext Transfer Protocol (HTTP)$')" -eq 1 ]
}

@test "You can get filter by ports only" {
    run ${BIN} ports 4
    [ $status -eq 0 ]
    [ "$(echo $output | grep -c '^4')" -eq 50 ]
}
