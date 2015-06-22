package main

import (
    "net"
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func CheckAllConnections() error {

    logger.Info("checking connections")

    for _, conn := range config.Connections {
        logger.Debug("checking connection and login for %s:%d", conn.Hostname, conn.Port)
        if err := CheckTCPConnection(conn.Hostname, conn.Port); err != nil {
            return err
        }
        if err := CheckMysqlLogin(conn); err != nil {
            return err
        }
    }



    return nil
}

func CheckTCPConnection(ip string, port int) error {

    addr := fmt.Sprintf("%s:%d", ip, port)

    logger.Debug("resolving %s:%d", ip, port)
    tcpAddr, resErr := net.ResolveTCPAddr("tcp4", addr)
    if resErr != nil {
        return fmt.Errorf("failed to resolve connection %s", addr)
    }

    logger.Debug("dialing %s", addr)
    conn, dialErr := net.DialTCP("tcp", nil, tcpAddr)
    if dialErr != nil {
        return fmt.Errorf("failed to connect to %s", addr)
    }

    logger.Info("OK: %s", addr)

    defer conn.Close()

    return nil
}

func CheckMysqlLogin(host ConnectionConfig) error {

    dsnBase := fmt.Sprintf("%s:%s@tcp(%s:%d)", host.Username, host.Password, host.Hostname, host.Port)

    for c, db := range host.Databases {
        logger.Debug("checking login for DB%d: %s", c, db)
        dsnFull := fmt.Sprintf("%s/%s", dsnBase, db)
        db, err := sql.Open("mysql", dsnFull)
        if err != nil { return err }

        defer db.Close()

        _, qErr := db.Query("SHOW TABLES")
        if qErr != nil { return qErr }

    }

    return nil
}