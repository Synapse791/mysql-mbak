package main

import (
    "net"
    "fmt"
)

func CheckAllConnections() error {

    logger.Info("checking connections")

    for c, conn := range config.Connections {
        logger.Debug("checking connection %d: %s:%d", c, conn.Hostname, conn.Port)
        if err := CheckTCPConnection(conn.Hostname, conn.Port); err != nil {
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