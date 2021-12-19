defport := "3000"
sname := "cy-test-server"

@server port=defport:
    mkdir -p ./bin
    go build -o ./bin/{{sname}} ./server
    ./bin/{{sname}} -port {{port}} &

@serverstop:
    killall {{sname}}

@cypressall: npmi server
    npm run cy:run
    just serverstop

@cypress: npmi
    npm run cy:open

@npmi:
    npm install --no-fund --no-audit

@dev port=defport:
    reflex -r '^server/' -s -- sh -c "go run ./server/main.go -port {{port}}"
