package main

import (
    "context"
    _ "github.com/mattn/go-sqlite3"
    log "github.com/sirupsen/logrus"
    "github.com/umi0410/how-to-backend-in-go/testcode/ent"
    "github.com/umi0410/how-to-backend-in-go/testcode/ent/migrate"
)

func main(){
    client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()
    if err := client.Schema.Create(
        context.TODO(),
        migrate.WithDropIndex(true),
        migrate.WithDropColumn(true),
    ); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

    userRepository := NewUserRepository(client.User)
    log.Info(userRepository)

}
